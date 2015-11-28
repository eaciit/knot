package hello

import (
	"github.com/eaciit/knot"
	"github.com/eaciit/knot/appcontainer"
	"github.com/eaciit/toolkit"
)

var (
	appViewsPath = "/Users/ariefdarmawan/goapp/src/github.com/eaciit/knot/example/hello/views/"
)

func init() {
	app := appcontainer.NewApp("Hello")
	app.Register(&WorldController{})
	app.Static("static", "/Users/ariefdarmawan/Temp")
	//app.View(toolkit.PathDefault(false) + "views")
	app.View(appViewsPath)
	appcontainer.RegisterApp(app)
}

type WorldController struct {
}

func (w *WorldController) Say(r *knot.Request) interface{} {
	s := "<b>Hello World</b>&nbsp;"
	name := r.Query("name")
	if name != "" {
		s = s + " " + name
	}
	s += "</br>"
	return s
}

func (w *WorldController) Index(r *knot.Request) interface{} {
	r.ResponseConfig().ViewName = appViewsPath + "template.html"
	return (toolkit.M{}).Set("message", "This is data passed to the template")
}
