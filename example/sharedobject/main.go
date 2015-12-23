package main

import (
	. "github.com/eaciit/knot/example/sharedobject/evening"
	. "github.com/eaciit/knot/example/sharedobject/morning"
	"github.com/eaciit/knot/knot.v1"
)

func main() {
	knot.DefaultOutputType = knot.OutputHtml
	ks := new(knot.Server)
	ks.Address = ":1234"
	ks.Register(new(Morning), "")
	ks.Register(new(Evening), "")

	ks.Listen()
}
