package main

import (
	"fmt"
	"github.com/eaciit/knot/knot.v1"
	"io"
	"os"
)

type Hello struct {
}

func (h *Hello) Index(r *knot.WebContext) interface{} {
	r.Writer.Header().Set("Content-Disposition", "attachment; filename=file.zip")
	r.Writer.Header().Set("Content-Type", r.Writer.Header().Get("Content-Type"))

	f, err := os.Open("file.zip")
	if err != nil {
		fmt.Println(err.Error())
	}

	io.Copy(r.Writer, f)

	return ""
}

func main() {
	knot.DefaultOutputType = knot.OutputHtml
	ks := new(knot.Server)
	ks.Address = ":1234"
	ks.Register(new(Hello), "")
	ks.Listen()
}

// try to access http://localhost:1234/hello/index
