package knot

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

func (r *Request) Write(w http.ResponseWriter, data interface{}) error {
	cfg := r.ResponseConfig()
	if cfg.OutputType == OutputTemplate {
		return r.WriteError(w, r.WriteTemplate(w, data))
	}

	if cfg.OutputType == OutputJson {
		return r.WriteError(w, r.WriteJson(w, data))
	}

	if cfg.OutputType == OutputByte || cfg.OutputType == OutputHtml {
		fmt.Fprint(w, data)
		return nil
	}

	return nil
}

func (r *Request) WriteTemplate(w http.ResponseWriter, data interface{}) error {
	var e error
	cfg := r.ResponseConfig()
	//w.Header().Set("Content-Type", "text/html")
	if cfg.ViewName != "" {
		useLayout := false
		viewsPath := cfg.ViewsPath
		fixLogicalPath(&viewsPath, true, true)
		viewFile := viewsPath
		if cfg.LayoutTemplate != "" {
			useLayout = true
			viewFile += cfg.LayoutTemplate
		} else {
			viewFile += cfg.ViewName
		}
		if useLayout {
			buf := bytes.Buffer{}
			e = r.writeToTemplate(&buf, data, cfg.ViewName)
			if e != nil {
				return e
			}
			e = r.writeToTemplate(w, struct{ Content interface{} }{
				template.HTML(string(buf.Bytes()))}, cfg.LayoutTemplate)
			if e != nil {
				return e
			}
		} else {
			e = r.writeToTemplate(w, data, cfg.ViewName)
		}
		if e != nil {
			return e
		}
	} else {
		return fmt.Errorf("No template define for %s", strings.ToLower(r.httpRequest.URL.String()))
	}
	return nil
}

func (r *Request) writeToTemplate(w io.Writer, data interface{}, templateFile string) error {
	cfg := r.ResponseConfig()
	viewsPath := cfg.ViewsPath
	viewFile := viewsPath + templateFile
	bs, e := ioutil.ReadFile(viewFile)
	if e != nil {
		return e
	}
	t, e := template.New("main").Funcs(template.FuncMap{
		"BaseUrl": func() string {
			base := "/"
			if cfg.AppName != "" {
				base += strings.ToLower(cfg.AppName)
			}
			if base != "/" {
				base += "/"
			}
			return base
		},
	}).Parse(string(bs))
	/*
		t.Funcs(template.FuncMap{
			"BaseUrl": func() string {
				base := "/"
				if cfg.AppName != "" {
					base += strings.ToLower(cfg.AppName)
				}
				if base != "/" {
					base += "/"
				}
				return base
			},
		})
	*/
	for _, includeFile := range cfg.IncludeFiles {
		if includeFile != cfg.LayoutTemplate && includeFile != templateFile {
			includeFilePath := viewsPath + includeFile
			_, e = t.New(includeFile).ParseFiles(includeFilePath)
			if e != nil {
				return e
			}
		}
	}
	e = t.Execute(w, data)
	if e != nil {
		return e
	}
	return nil
}

func (r *Request) WriteJson(w http.ResponseWriter, data interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(data)
}

func (r *Request) WriteError(w http.ResponseWriter, e error) error {
	if e != nil {
		errorString := e.Error()
		hr := r.HttpRequest()
		r.Server().Log().Error(fmt.Sprintf("%s %s Error: %s", hr.URL.String(), hr.RemoteAddr, errorString))
		fmt.Fprint(w, errorString)
	}
	return e
}
