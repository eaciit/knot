package knot

import ()

type OutputType int

const (
	OutputHtml OutputType = 1
	OutputJson OutputType = 10
	OutputByte OutputType = 100
)

type ResponseConfig struct {
	AppName    string
	ViewName   string
	OutputType OutputType
	Views      []string

	Headers map[string]string
}

func NewResponseConfig() *ResponseConfig {
	c := new(ResponseConfig)
	c.Headers = map[string]string{}
	return c
}
