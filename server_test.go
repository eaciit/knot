package knot

import (
	//"fmt"
	"github.com/eaciit/toolkit"
	"net/http"
	"testing"
)

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
	if str != "Welcome to Sebar Server" {
		t.Errorf("Invalid return")
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
