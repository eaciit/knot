package hello

import (
	"github.com/eaciit/knot"
	"github.com/eaciit/knot/appcontainer"
)

func init() {
	app := appcontainer.NewApp("Hello")
	app.Register(&WorldController{})
	app.Static("static", "/Users/ariefdarmawan/Temp")
	appcontainer.RegisterApp(app)
}

type WorldController struct {
}

func (w *WorldController) Say(r *knot.Request) interface{} {
	s := "Hello World"
	name := r.Query("name")
	if name != "" {
		s = s + " " + name
	}
	return s
}
