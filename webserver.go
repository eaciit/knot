package knot

import (
	"fmt"
	"github.com/eaciit/toolkit"
	"github.com/gorilla/mux"
	"net/http"
	"reflect"
	"strings"
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

type FnContent func(r *Request) interface{}

func (s *Server) router() *mux.Router {
	if s.mxrouter == nil {
		s.mxrouter = mux.NewRouter()
	}
	return s.mxrouter
}

func (s *Server) Register(c interface{}, prefix string, cfgs map[string]*RouteConfig) error {
	var t reflect.Type
	v := reflect.ValueOf(c)
	if v.Kind() != reflect.Ptr {
		return fmt.Errorf("Invalid controller object passed (%s). Controller object should be a pointer", v.Kind())
	}
	t = reflect.TypeOf(c)
	controllerName := reflect.Indirect(v).Type().Name()

	s.Log().Info(fmt.Sprintf("Registering %s", controllerName))
	path := prefix
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}
	if !strings.HasSuffix(path, "/") {
		path = path + "/"
	}

	controllerName = strings.ToLower(controllerName)
	if strings.HasSuffix(controllerName, "controller") {
		controllerName = controllerName[0 : len(controllerName)-len("controller")]
	}
	path += controllerName + "/"

	if t == nil {
	}
	methodCount := t.NumMethod()
	for mi := 0; mi < methodCount; mi++ {
		method := t.Method(mi)

		// validate if this method match FnContent
		isFnContent := false
		tm := method.Type
		if tm.NumIn() == 2 && tm.In(1).String() == "*knot.Request" {
			if tm.NumOut() == 1 && tm.Out(0).Kind() == reflect.Interface {
				isFnContent = true
			}
		}

		if isFnContent {
			var cfg *RouteConfig
			var fnc FnContent
			fnc = v.MethodByName(method.Name).Interface().(func(*Request) interface{})
			handlerPath := path + strings.ToLower(method.Name)
			s.Log().Info(fmt.Sprintf("Registering handler for %s", handlerPath))
			s.Route(handlerPath, fnc, cfg)
		}
	}

	return nil
}

func (s *Server) Route(path string, fnc FnContent, cfg *RouteConfig) {
	s.router().HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		if fnc != nil {
			kr := new(Request)
			kr.server = s
			kr.httpRequest = r
			v := fnc(kr)
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

func (s *Server) Listen() {
	s.start()
	s.listen()
}

func (s *Server) start() error {
	addr := s.Address
	s.status = make(chan string)
	s.Log().Info("Start listening on server " + addr)
	go func() {
		http.ListenAndServe(addr, s.router())
	}()
	return nil
}

func (s *Server) Stop() {
	s.Log().Info(fmt.Sprintf("Stopping server %s", s.Address))
	s.status <- "Stop"
}

func (s *Server) listen() {
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
