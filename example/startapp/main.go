package main

import (
	"github.com/eaciit/knot/knot.v1"
	"net/http"
)

type Hello struct {
}

func (h *Hello) Index(r *knot.WebContext) interface{} {
	r.Config.OutputType = knot.OutputHtml
	return `Redirected from /`
}

func main() {
	app := knot.NewApp("test")
	app.Register(&Hello{})
	knot.RegisterApp(app)

	otherRoutes := map[string]knot.FnContent{
		"/": func(r *knot.WebContext) interface{} {
			http.Redirect(r.Writer, r.Request, "/hello/index", http.StatusTemporaryRedirect)
			return true
		},
	}

	knot.StartAppWithFn(app, "localhost:8999", otherRoutes)

}
