package webext

import (
	"github.com/eaciit/knot/knot.v1"
	"os"
)

var (
	wd = func() string {
		d, _ := os.Getwd()
		return d + "/../"
	}()
)

func init() {
	app := knot.NewApp("ext")
	app.ViewsPath = wd + "views/"
	app.Register(&AppController{})
	app.Static("static", wd+"assets")
	app.LayoutTemplate = "_layout.html"
	knot.RegisterApp(app)
}

type AppController struct {
}

func (a *AppController) Default(k *knot.WebContext) interface{} {
	k.Config.OutputType = knot.OutputTemplate
	return ""
}
