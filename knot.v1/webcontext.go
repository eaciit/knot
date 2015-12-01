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

type WebContext struct {
	Config  *ResponseConfig
	Server  *Server
	Request *http.Request
	Writer  http.ResponseWriter

	queryKeys []string
	cookies   map[string]*http.Cookie
}

func (r *WebContext) QueryKeys() []string {
	if len(r.queryKeys) == 0 {
		if r.Request == nil {
			return r.queryKeys
		}

		values := r.Request.URL.Query()
		for k, _ := range values {
			r.queryKeys = append(r.queryKeys, k)
		}
	}
	return r.queryKeys
}

func (r *WebContext) Query(id string) string {
	if r.Request == nil {
		return ""
	}

	return r.Request.URL.Query().Get(id)
}

func (r *WebContext) GetPayload(result interface{}) error {
	if r.Request == nil {
		return errors.New("HttpRequest object is not properly setup")
	}

	body := r.Request.Body
	defer body.Close()
	decoder := json.NewDecoder(body)
	return decoder.Decode(result)
}

func (r *WebContext) GetPayloadMultipart(result interface{}) (map[string][]*multipart.FileHeader,
	map[string][]string, error) {
	var e error
	if r.Request == nil {
		return nil, nil, errors.New("HttpRequest object is not properly setup")
	}
	e = r.Request.ParseMultipartForm(1024 * 1024 * 1024 * 1024)
	if e != nil {
		return nil, nil, fmt.Errorf("Unable to parse: %s", e.Error())
	}
	m := r.Request.MultipartForm
	return m.File, m.Value, nil
}
