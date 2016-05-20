package ui

import (
	"html/template"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/GeertJohan/go.rice"
	"github.com/juju/errgo"
)

const (
	templateDir  = "../static/templates"
	rootTemplate = "root.tpl"
)

type uiServer struct {
	rootTemplate *template.Template
}

func NewUIServer() (*uiServer, error) {
	tplBox, err := rice.FindBox(templateDir)
	if err != nil {
		return nil, errgo.Mask(err)
	}
	rootTplR, err := tplBox.Open(rootTemplate)
	if err != nil {
		return nil, errgo.Mask(err)
	}
	defer rootTplR.Close()
	rootTpl, err := tplFromReader("root", rootTplR)
	if err != nil {
		return nil, errgo.Mask(err)
	}
	return &uiServer{
		rootTemplate: rootTpl,
	}, nil
}

func (s *uiServer) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	s.rootTemplate.Execute(resp, nil)
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
