package knot

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"net/http"
)

func (r *Request) Write(w http.ResponseWriter, data interface{}) error {
	var e error
	cfg := r.ResponseConfig()
	if cfg.OutputType == OutputJson {
		return r.WriteJson(w, data)
	}

	if cfg.OutputType == OutputByte {
		fmt.Fprint(w, data)
		return nil
	}

	w.Header().Set("Content-Type", "text/html")
	if cfg.ViewName != "" || cfg.LayoutTemplate != "" {
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
			if cfg.ViewName != "" {
				buf := bytes.Buffer{}
				e = r.WriteToTemplate(&buf, data, cfg.ViewName)
				if e != nil {
					r.WriteError(w, e.Error())
					return nil
				}
				e = r.WriteToTemplate(w, struct{ Content interface{} }{
					template.HTML(string(buf.Bytes()))}, cfg.LayoutTemplate)
				if e != nil {
					r.WriteError(w, e.Error())
					return nil
				}
			} else {
				e = r.WriteToTemplate(w, struct{ Content interface{} }{data}, cfg.LayoutTemplate)
			}
		} else {
			e = r.WriteToTemplate(w, data, cfg.ViewName)
		}
		if e != nil {
			r.WriteError(w, e.Error())
			return nil
		}
	} else {
		fmt.Fprint(w, data)
		return nil
	}

	return nil
}

func (r *Request) WriteToTemplate(w io.Writer, data interface{}, templateFile string) error {
	cfg := r.ResponseConfig()
	viewsPath := cfg.ViewsPath
	viewFile := viewsPath + templateFile
	t, e := template.ParseFiles(viewFile)
	if e != nil {
		return e
	}
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

func (r *Request) WriteError(w http.ResponseWriter, errorString string) {
	hr := r.HttpRequest()
	r.Server().Log().Error(fmt.Sprintf("%s %s Error: %s", hr.URL.String(), hr.RemoteAddr, errorString))
	fmt.Fprint(w, errorString)
}
