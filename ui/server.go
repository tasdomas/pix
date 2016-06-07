package ui

import (
	"html/template"
	"io"
	"io/ioutil"
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

type uiServer struct {
	templates map[string]*template.Template
	router    *httprouter.Router
}

func NewUIServer() (*uiServer, error) {
	tpls, err := loadTemplates(templates)
	if err != nil {
		return nil, errgo.Mask(err)
	}

	s := &uiServer{
		templates: tpls,
		router:    httprouter.New(),
	}
	s.POST("/upload", s.upload)
	return s, nil
}

func (s *uiServer) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	s.router.ServeHTTP(resp, req)
	//s.templates["root"].Execute(resp, nil)
}

func (s *uiServer) upload(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

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
