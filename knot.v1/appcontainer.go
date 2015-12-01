package knot

import (
	"fmt"
	"io/ioutil"
	//"os"
	"path/filepath"
	"reflect"
	"strings"
	// -- KnotApp Registration Start
	// -- KnotAppRegistration End
)

var (
	apps = map[string]*App{}
)

type App struct {
	Name           string
	Enable         bool
	LayoutTemplate string
	ViewsPath      string

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
	if path == "" {
		delete(a.Statics(), prefix)
		return
	}
	a.Statics()[prefix] = path
}

func (a *App) Controllers() map[string]interface{} {
	if a.controllers == nil {
		a.controllers = map[string]interface{}{}
	}
	return a.controllers
}

func NewApp(name string) *App {
	app := new(App)
	app.Name = name
	app.Enable = true
	return app
}

type AppContainerConfig struct {
	Address string
}

func RegisterApp(app *App) {
	apps[app.Name] = app
}

func getIncludeFiles(dirname string) []string {
	fis, e := ioutil.ReadDir(dirname)
	if e != nil {
		return []string{}
	}

	files := []string{}
	for _, fi := range fis {
		if fi.IsDir() {
			files = append(files, getIncludeFiles(filepath.Join(dirname, fi.Name()))...)
		} else if strings.HasPrefix(fi.Name(), "_") { //--- include is file started with _
			files = append(files, fi.Name())
		}
	}
	return files
}

func StartContainer(c *AppContainerConfig) *Server {
	ks := new(Server)
	ks.Address = c.Address

	for k, app := range apps {
		appname := strings.ToLower(k)
		//-- need to handle appname translation in Regex way
		if strings.Contains(appname, " ") {
			appname = strings.Replace(appname, " ", "", 0)
		}
		//-- end of regex
		includes := []string{}
		if app.ViewsPath != "" {
			includes = getIncludeFiles(app.ViewsPath)
		}
		ks.Log().Info("Scan application " + appname + " for controller registration")
		for _, controller := range app.Controllers() {
			ks.RegisterWithConfig(controller, appname, &ResponseConfig{
				AppName:        k,
				ViewsPath:      app.ViewsPath,
				LayoutTemplate: app.LayoutTemplate,
				IncludeFiles:   includes,
			})
		}

		for surl, spath := range app.Statics() {
			staticUrlPrefix := appname + "/" + surl
			ks.RouteStatic(staticUrlPrefix, spath)
		}
	}

	ks.Route("/status", statusContainer)
	ks.Route("/stop", stopContainer)
	//ks.Route("/p", ShowPage)
	ks.Listen()

	return ks
}

func stopContainer(r *WebContext) interface{} {
	defer r.Server.Stop()
	return "Knot Server (" + r.Server.Address + ") will be stopped. Bye"
}

func statusContainer(r *WebContext) interface{} {
	str := "Knot Server v1.0 (c) Eaciit"
	return str
}
