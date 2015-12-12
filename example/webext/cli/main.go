package main

import (
	_ "github.com/eaciit/knot/example/webext"
	"github.com/eaciit/knot/knot.v1"
)

func main() {
	/*
		knot.StartContainer(&knot.AppContainerConfig{
			Address: "localhost:12345",
		})
	*/

	app := knot.GetApp("ext")
	if app == nil {
		return
	}
	knot.StartApp(app, "localhost:12345")

}
