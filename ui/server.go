package ui

import (
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/GeertJohan/go.rice"
	"github.com/juju/errgo"
	"github.com/julienschmidt/httprouter"
)

const (
	templateDir = "../static/templates"
)

var (
	templates = map[string]string{
		"root": "root.tpl",
	}
)

type storage interface {
	Put(io.ReadSeeker) (string, error)
	Get(string) (io.ReadCloser, error)
	List() ([]string, error)
}

type uiServer struct {
	templates map[string]*template.Template
	router    *httprouter.Router
	storage   storage
}

func NewServer(st storage) (*uiServer, error) {
	tpls, err := loadTemplates(templates)
	if err != nil {
		return nil, errgo.Mask(err)
	}

	s := &uiServer{
		templates: tpls,
		router:    httprouter.New(),
		storage:   st,
	}
	s.router.POST("/upload", s.upload)
	return s, nil
}

func (s *uiServer) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	s.router.ServeHTTP(resp, req)
	//s.templates["root"].Execute(resp, nil)
}

func (s *uiServer) upload(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// TODO check secret here.
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
	tplBox, err := rice.FindBox(templateDir)
	if err != nil {
		return nil, errgo.Mask(err)
	}
	result := make(map[string]*template.Template)
	for name, file := range templateLocations {
		tplRaw, err := tplBox.Open(file)
		if err != nil {
			return nil, errgo.Mask(err)
		}
		tpl, err := tplFromReader("root", tplRaw)
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
