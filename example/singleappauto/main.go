package main

import (
	"github.com/eaciit/knot/knot.v1"
)

type Hello struct {
}

// http://servername/hello/morning
func (h *Hello) Morning(r *knot.WebContext) interface{} {
	return "Good morning"
}

// http://servername/hello/evening
func (h *Hello) Evening(r *knot.WebContext) interface{} {
	return "Good evening"
}

func (h *Hello) Night() interface{} {
	return "Good Night"
}

func main() {
	knot.DefaultOutputType = knot.OutputHtml
	ks := new(knot.Server)
	ks.Address = "localhost:13000"
	ks.Register(new(Hello), "")
	ks.Route("/stop", func(w *knot.WebContext) interface{} {
		defer w.Server.Stop()
		return ""
	})
	ks.Listen()
}
