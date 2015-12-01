package hello

import (
	"github.com/eaciit/knot/knot.v1"
	"github.com/eaciit/toolkit"
	"os"
	"time"
)

var (
	//appViewsPath = "/Users/ariefdarmawan/goapp/src/github.com/eaciit/knot/example/hello/views/"
	appViewsPath = func() string {
		d, _ := os.Getwd()
		return d
	}() + "/../example/hello/views/"
)

func init() {
	app := knot.NewApp("Hello")
	app.ViewsPath = appViewsPath
	app.Register(&WorldController{})
	app.Static("static", "/Users/ariefdarmawan/Temp")
	app.LayoutTemplate = "_template.html"
	knot.RegisterApp(app)
}

type WorldController struct {
}

func (w *WorldController) Say(r *knot.WebContext) interface{} {
	r.Config.OutputType = knot.OutputHtml
	s := "<b>Hello World</b>&nbsp;"
	name := r.Query("name")
	if name != "" {
		s = s + " " + name
	}
	s += "</br>"
	return s
}

func (w *WorldController) Index(r *knot.WebContext) interface{} {
	//r.Config.ViewName = "hello.html"
	return (toolkit.M{}).Set("message", "This is data passed to the template")
}

func (w *WorldController) Cookie(r *knot.WebContext) interface{} {
	r.Config.OutputType = knot.OutputHtml
	cvalue := ""
	cookiename := "mycookie"
	c, _ := r.Cookie(cookiename, "")
	if c == nil {
		r.SetCookie(cookiename, "Arief Darmawan", 30*24*time.Hour)
	} else {
		cvalue = c.Value
		c.Value = "Arief Darmawan" + time.Now().String()
		c.Expires = time.Now().Add(24 * 30 * time.Hour)
	}
	return "Cookie value is " + cvalue
}

func (w *WorldController) Session(r *knot.WebContext) interface{} {
	r.Config.OutputType = knot.OutputHtml
	s := r.Session("NameAndTime", "").(string)
	r.SetSession("NameAndTime", "Arief Darmawan : "+time.Now().String())
	return "Session value is " + s
}

func (w *WorldController) Json(r *knot.WebContext) interface{} {
	r.Config.OutputType = knot.OutputJson
	return struct{ ID, Title string }{"JsonID 00001", "Json Row 01 - Title"}
}
