package knot

import "net/http"

type Request struct {
	server      *Server
	httpRequest *http.Request
}

func (r *Request) Server() *Server {
	return r.server
}

func (r *Request) HttpRequest() *http.Request {
	return r.httpRequest
}
