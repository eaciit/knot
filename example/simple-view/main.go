package main

import (
	"github.com/eaciit/knot/knot.v1"
	"os"
)

var (
	appViewsPath = (func(dir string, _ error) string { return dir + "/views/" }(os.Getwd()))
)

func main() {
	app := knot.NewApp("web")
	app.ViewsPath = appViewsPath
	app.Register(&MainController{})
	app.LayoutTemplate = "_template.html"
	app.DefaultOutputType = knot.OutputTemplate

	knot.RegisterApp(app)
	knot.StartContainer(&knot.AppContainerConfig{
		Address: "localhost:1234",
	})
}

type MainController struct {
}

func (w *MainController) Index(r *knot.WebContext) interface{} {
	r.Config.ViewName = "index.html"
	return nil
}
