package knot

import ()

type OutputType int

const (
	OutputTemplate OutputType = 0
	OutputHtml     OutputType = 1
	OutputJson     OutputType = 10
	OutputByte     OutputType = 100
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
}

func NewResponseConfig() *ResponseConfig {
	c := new(ResponseConfig)
	c.Headers = map[string]string{}
	c.IncludeFiles = []string{}
	return c
}
