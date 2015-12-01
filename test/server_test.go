package test

import (
	//"fmt"
	. "github.com/eaciit/knot/knot.v1"
	"github.com/eaciit/toolkit"
	"net/http"
	"testing"
)

var ks *Server = new(Server)

func init() {
	ks.Address = ":13000"
	DefaultOutputType = OutputHtml
	ks.Route("/", func(wc *WebContext) interface{} {
		return "Welcome to Knot Server"
	})
	ks.Route("/stop", func(wc *WebContext) interface{} {
		defer wc.Server.Stop()
		return "Server will be stopped. Bye"
	})
	go func() {
		ks.Listen()
	}()
}

func call(url string) (*http.Response, error) {
	surl := "http://localhost:13000" + url
	r, e := toolkit.HttpCall(surl, "GET", nil, nil)
	return r, e
}

func TestServer(t *testing.T) {
	r, e := call("/")
	if e != nil {
		t.Errorf("Fail: %s", e.Error())
		return
	}

	if r.StatusCode != 200 {
		t.Errorf("Error: %s", e.Error())
		return
	}

	str := toolkit.HttpContentString(r)
	want := "Welcome to Knot Server"
	if str != want {
		t.Errorf("Invalid return. Expecting %s got %s", want, str)
	}
}

func TestClose(t *testing.T) {
	r, e := call("/stop")
	if e != nil {
		t.Errorf("Fail: %s", e.Error())
	} else {
		t.Logf("Respond: %s", toolkit.HttpContentString(r))
	}
}
