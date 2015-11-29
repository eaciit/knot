package knot

import (
	"encoding/json"
	"errors"
	"fmt"
	"mime/multipart"
	"net/http"
)

var (
	DefaultOutputType OutputType
)

type Request struct {
	server         *Server
	httpRequest    *http.Request
	responseConfig *ResponseConfig

	queryKeys []string
}

func (r *Request) Server() *Server {
	if r.server == nil {
		r.server = new(Server)
	}
	return r.server
}

func (r *Request) HttpRequest() *http.Request {
	if r.httpRequest == nil {
		r.httpRequest = new(http.Request)
	}
	return r.httpRequest
}

func (r *Request) ResponseConfig() *ResponseConfig {
	if r.responseConfig == nil {
		r.responseConfig = NewResponseConfig()
	}
	r.responseConfig.OutputType = DefaultOutputType
	return r.responseConfig
}

func (r *Request) QueryKeys() []string {
	if len(r.queryKeys) == 0 {
		if r.HttpRequest() == nil {
			return r.queryKeys
		}

		values := r.HttpRequest().URL.Query()
		for k, _ := range values {
			r.queryKeys = append(r.queryKeys, k)
		}
	}
	return r.queryKeys
}

func (r *Request) Query(id string) string {
	if r.httpRequest == nil {
		return ""
	}

	return r.HttpRequest().URL.Query().Get(id)
}

func (r *Request) GetPayload(result interface{}) error {
	if r.httpRequest == nil {
		return errors.New("HttpRequest object is not properly setup")
	}

	body := r.httpRequest.Body
	defer body.Close()
	decoder := json.NewDecoder(body)
	return decoder.Decode(result)
}

func (r *Request) GetPayloadMultipart(result interface{}) (map[string][]*multipart.FileHeader,
	map[string][]string, error) {
	var e error
	if r.httpRequest == nil {
		return nil, nil, errors.New("HttpRequest object is not properly setup")
	}
	e = r.httpRequest.ParseMultipartForm(1024 * 1024 * 1024 * 1024)
	if e != nil {
		return nil, nil, fmt.Errorf("Unable to parse: %s", e.Error())
	}
	m := r.httpRequest.MultipartForm
	return m.File, m.Value, nil
}

func (r *Request) setHeaders(w http.ResponseWriter, data interface{}) {
	cfg := r.ResponseConfig()
	for k, v := range cfg.Headers {
		w.Header().Set(k, v)
	}
}
