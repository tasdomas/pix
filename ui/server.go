package ui

import (
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"sort"

	"github.com/GeertJohan/go.rice"
	"github.com/juju/errgo"
	"github.com/julienschmidt/httprouter"
)

const baseTemplate = "base.tpl"

var (
	templates = map[string]string{
		"root": "root.tpl",
		"img":  "img.tpl",
	}
)

type storage interface {
	Put(io.ReadSeeker) (string, error)
	Get(string, string) (io.ReadCloser, error)
	List() ([]string, error)
}

type uiServer struct {
	templates map[string]*template.Template
	router    *httprouter.Router
	storage   storage
	secret    string
	name      string
}

func NewServer(st storage, name, secret string) (*uiServer, error) {
	tpls, err := loadTemplates(templates)
	if err != nil {
		return nil, errgo.Mask(err)
	}

	s := &uiServer{
		name:      name,
		secret:    secret,
		templates: tpls,
		router:    httprouter.New(),
		storage:   st,
	}

	staticBox, err := rice.FindBox("../static/serve")
	if err != nil {
		return nil, errgo.Mask(err)
	}

	s.router.ServeFiles("/static/*filepath", staticBox.HTTPBox())
	s.router.GET("/favicon.ico", mustServeFile(staticBox, "favicon.ico"))
	s.router.POST("/upload", s.upload)
	s.router.GET("/", s.root)
	s.router.GET("/image/:img", s.imagePage)
	s.router.GET("/image/:img/raw", s.serveImageMod(""))
	s.router.GET("/image/:img/thumb", s.serveImageMod("thb"))
	s.router.GET("/image/:img/large", s.serveImageMod("large"))
	return s, nil
}

func (s *uiServer) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	s.router.ServeHTTP(resp, req)
}

func (s *uiServer) imagePage(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	img := params.ByName("img")
	if img == "" {
		w.WriteHeader(http.StatusBadRequest)
	}
	_, err := s.storage.Get(img, "")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("could not retrieve image from storage: %s", err.Error())
	}
	tplParams := struct {
		SiteName string
		Image    string
	}{
		SiteName: s.name,
		Image:    img,
	}
	s.templates["img"].Execute(w, tplParams)
}

func (s *uiServer) serveImageMod(mod string) func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		w.Header().Set("Content-Type", "image/jpeg")
		img := params.ByName("img")
		if img == "" {
			w.WriteHeader(http.StatusBadRequest)
		}
		f, err := s.storage.Get(img, mod)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Printf("could not retrieve image from storage: %s", err.Error())
		}
		defer f.Close()
		io.Copy(w, f)
	}
}

func (s *uiServer) root(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	list, err := s.storage.List()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "%s", err.Error())
	}
	sort.Strings(list)

	params := struct {
		SiteName string
		Images   []string
	}{
		SiteName: s.name,
		Images:   list,
	}

	s.templates["root"].Execute(w, params)
}

func (s *uiServer) upload(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	secret := r.FormValue("secret")
	if secret == "" || secret != s.secret || s.secret == "" {
		w.WriteHeader(http.StatusForbidden)
		log.Printf("rejected file upload, provided key was %q", secret)
		return
	}
	err := r.ParseMultipartForm(10 * 1024 * 1024)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "%s", err.Error())
	}
	files := r.MultipartForm.File["image"]
	if len(files) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "no image upload found")
	}
	for _, f := range files {
		data, err := f.Open()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "%s", err.Error())
		}
		hash, err := s.storage.Put(data)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "%s", err.Error())
		}
		log.Printf("stored image: %s", hash)
	}
}

func loadTemplates(templateLocations map[string]string) (map[string]*template.Template, error) {
	tplBox, err := rice.FindBox("../static/templates")
	if err != nil {
		return nil, errgo.Mask(err)
	}
	tplBase, err := tplBox.Open(baseTemplate)
	if err != nil {
		return nil, errgo.Mask(err)
	}

	result := make(map[string]*template.Template)
	for name, file := range templateLocations {
		tplBase.Seek(0, 0)
		tplRaw, err := tplBox.Open(file)
		if err != nil {
			return nil, errgo.Mask(err)
		}
		tpl, err := tplFromReader("root", tplRaw, tplBase)
		if err != nil {
			tplRaw.Close()
			return nil, errgo.Mask(err)
		}
		tplRaw.Close()
		result[name] = tpl
	}
	return result, nil

}

func tplFromReader(name string, contents ...io.Reader) (*template.Template, error) {
	tpl := template.New(name)
	for _, t := range contents {
		data, err := ioutil.ReadAll(t)
		if err != nil {
			return nil, errgo.Mask(err)
		}
		_, err = tpl.Parse(string(data))
		if err != nil {
			return nil, errgo.Mask(err)
		}
	}
	return tpl, nil
}

func mustServeFile(box *rice.Box, name string) httprouter.Handle {
	f, err := box.Open(name)
	if err != nil {
		panic(err)
	}
	return func(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
		f.Seek(0, 0)
		io.Copy(w, f)
	}
}
