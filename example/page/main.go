package main

import (
	"github.com/eaciit/knot/knot.v1"
	"net/http"
)

func main() {
	app := knot.NewApp("page")
	app.Register(&PageController{})
	knot.RegisterApp(app)

	otherRoutes := map[string]knot.FnContent{
		"/": knot.FnContent(func(r *knot.WebContext) interface{} {
			r.Config.OutputType = knot.OutputHtml
			return "index"
		}),
		"page": knot.FnContent(func(r *knot.WebContext) interface{} {
			r.Config.OutputType = knot.OutputHtml
			pageID := r.Request.Header.Get("PAGE_ID")
			return pageID
		}),
	}

	knot.StartAppWithFn(app, ":8999", otherRoutes)
}

type PageController struct {
}

func (h *PageController) Index(r *knot.WebContext) interface{} {
	r.Config.OutputType = knot.OutputHtml
	http.Redirect(r.Writer, r.Request, "/page/from_page_index", http.StatusTemporaryRedirect)
	return true
}
