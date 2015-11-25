package knot

import (
	"fmt"
	"github.com/eaciit/toolkit"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

type Server struct {
	Address string

	mxrouter *mux.Router
	log      *toolkit.LogEngine
	status   chan string
}

func (s *Server) Log() *toolkit.LogEngine {
	if s.log == nil {
		s.log, _ = toolkit.NewLog(true, false, "", "", "")
	}
	return s.log
}

type FnContent func(svr *Server, r *http.Request) interface{}

func (s *Server) router() *mux.Router {
	if s.mxrouter == nil {
		s.mxrouter = mux.NewRouter()
	}
	return s.mxrouter
}

func (s *Server) Route(path string, fnc FnContent, cfg *RouteConfig) {
	s.router().HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		if fnc != nil {
			v := fnc(s, r)
			/*
				vof := reflect.ValueOf(v)
				kof := vof.Kind()
				s.Log().Info(fmt.Sprintf("Kind: %s", kof.String()))
				bs := toolkit.GetEncodeByte(v)
				w.Write(bs)
			*/
			fmt.Fprint(w, v)
		} else {
			w.Write([]byte(""))
		}
	})
}

func (s *Server) GetHandler(path string) http.Handler {
	mr := s.router().GetRoute(path)
	if mr == nil {
		return nil
	}
	return mr.GetHandler()
}

func (s *Server) Start(addr string) error {
	s.status = make(chan string)
	s.Log().Info("Start listening on server " + addr)
	go func() {
		http.ListenAndServe(addr, s.router())
	}()
	return nil
}

func (s *Server) Stop() {
	s.Log().Info("Stopping server")
	s.status <- "Stop"
	s.Log().Info("Server has been stopped successfully")
}

func (s *Server) Wait() {
	running := true
	for running {
		select {
		case status := <-s.status:
			if status == "Stop" {
				running = false
			}

		default:
			time.Sleep(1 * time.Millisecond)
		}
	}
}
