package main

import (
	. "github.com/eaciit/knot/knot.v1"
)

func main() {
	ks := new(Server)
	ks.Address = "localhost:13000"
	DefaultOutputType = OutputHtml

	// http://localhost:13000/
	ks.Route("/", func(wc *WebContext) interface{} {
		return "Welcome to Knot Server"
	})

	// http://localhost:13000/stop
	ks.Route("/stop", func(wc *WebContext) interface{} {
		defer wc.Server.Stop()
		return "Server will be stopped. Bye"
	})
	ks.Listen()
}
