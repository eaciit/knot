package appcontainer

import (
	"fmt"
	"github.com/eaciit/kingpin"
	"github.com/eaciit/knot"
	"reflect"
	"strings"
	// -- KnotApp Registration Start
	// -- KnotAppRegistration End
)

var (
	apps = map[string]*App{}
	ks   *knot.Server

	flagAddress = kingpin.Flag("address",
		"Address to be used by Knot Server. It normally formatted as SERVERNAME:PORTNUMBER").Default("localhost:9876").String()
)

type App struct {
	Name   string
	Enable bool

	views       []string
	controllers map[string]interface{}
	statics     map[string]string
}

func (a *App) Register(c interface{}) error {
	v := reflect.ValueOf(c)
	if v.Kind() != reflect.Ptr {
		return fmt.Errorf("Unable to register %v, type is not pointer \n", c)
	}

	name := strings.ToLower(reflect.Indirect(v).Type().Name())
	a.Controllers()[name] = c
	return nil
}

func (a *App) Statics() map[string]string {
	if a.statics == nil {
		a.statics = map[string]string{}
	}
	return a.statics
}

func (a *App) Static(prefix, path string) {
	a.Statics()[prefix] = path
}

func (a *App) Controllers() map[string]interface{} {
	if a.controllers == nil {
		a.controllers = map[string]interface{}{}
	}
	return a.controllers
}

func (a *App) View(s string) {
	if a.views == nil {
		a.views = []string{}
	}
	a.views = append(a.views, s)
}

func NewApp(name string) *App {
	app := new(App)
	app.Name = name
	app.Enable = true
	return app
}

type Config struct {
	Address string
}

func RegisterApp(app *App) {
	apps[app.Name] = app
}

func Start(c *Config) {
	if ks == nil {
		ks = new(knot.Server)
	}
	ks.Address = c.Address

	for k, app := range apps {
		appname := strings.ToLower(k)
		//-- need to handle appname translation in Regex way
		if strings.Contains(appname, " ") {
			appname = strings.Replace(appname, " ", "", 0)
		}
		//-- end of regex
		ks.Log().Info("Scan application " + appname + " for controller registration")
		for _, controller := range app.Controllers() {
			ks.RegisterWithConfig(controller, appname, &knot.ResponseConfig{
				AppName: k,
				Views:   app.views,
			})
		}

		for surl, spath := range app.Statics() {
			staticUrlPrefix := appname + "/" + surl
			ks.RouteStatic(staticUrlPrefix, spath)
		}
	}

	ks.Route("/status", Status)
	ks.Route("/stop", StopServer)
	ks.Listen()
}

func StopServer(r *knot.Request) interface{} {
	r.Server().Stop()
	return nil
}

func Status(r *knot.Request) interface{} {
	str := "Knot Server v0.8 (c) Eaciit"
	return str
}
