package knot

import (
	"net/http"
)

type OutputType int

const (
	OutputTemplate OutputType = 0
	OutputHtml     OutputType = 10
	OutputJson     OutputType = 100
	OutputByte     OutputType = 1000
)

func (o OutputType) String() string {
	if o == OutputTemplate {
		return "Template"
	} else if o == OutputHtml {
		return "HTML"
	} else if o == OutputJson {
		return "JSON"
	} else if o == OutputByte {
		return "Byte"
	}
	return "N/A"
}

type ResponseConfig struct {
	AppName        string
	ViewName       string
	OutputType     OutputType
	LayoutTemplate string
	ViewsPath      string
	IncludeFiles   []string

	Headers map[string]string
	cookies map[string]*http.Cookie
}

func NewResponseConfig() *ResponseConfig {
	c := new(ResponseConfig)
	c.Headers = map[string]string{}
	c.IncludeFiles = []string{}
	c.OutputType = DefaultOutputType
	return c
}

func (r *ResponseConfig) Cookies() map[string]*http.Cookie {
	if r.cookies == nil {
		r.cookies = map[string]*http.Cookie{}
	}
	return r.cookies
}
