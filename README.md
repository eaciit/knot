# knot
Knot is a web server and application container for Golang Web Based App. It is still on experimental version

# Background
I've been working with Golang for sometime. While some stack are build on either command line or REST, I always use web based application as main UI. 
Most of the time, our application will be proxied by Nginx, and to be honest, it is complicated effort, because we have to change nginx config file and restart it whenever we have new application developed. These then inspired me to build Knot.

# Usage
## Load Knot
```go
go get -u github.com/eaciit/knot
```

## FnContent
In knot, we handle function based on FnContent contract. Where contract of FnContent is as follow
```go
var FnContent func(*WebContext)interface{}
```
knot.WebContext is wrapper of KnotServer object, HttpRequest, Writer and other related context of Knot. Any data returned by this function will be used as body of the response depend on OutputType defined.

## Route 
To do routing in knot, is simply by invoke Route function
```go
func main(){
  knot.DefaultOutputType = knot.OutputHtml
  ks := new(knot.Server)
  ks.Route("hi",Hi)
  ks.Listen()
}

func Hi(r *knot.WebContext)interface{}{
  return "Hello World!"
}
```

Route to static folder
```go
ks.RouteStatic("static","/Users/ariefdarmawan/Temp/knot/app1/static")
```

## Start Knot as Single App Container
Below code will run knot on locahost:13000 and add 2 web method / and  /stop

### Manual Route
```go
package main

import (
  . "github.com/eaciit/knot/knot.v1"
)

func main() {
  ks := new(Server)
  ks.Address = "localhost:13000"
  DefaultOutputType = OutputHtml
  ks.Route("/", func(wc *WebContext) interface{} {
    return "Welcome to Knot Server"
  })
  ks.Route("/stop", func(wc *WebContext) interface{} {
    defer wc.Server.Stop()
    return "Server will be stopped. Bye"
  })
  ks.Listen()
}
```

### AutoRoute
Controller is a struct with sets of FnContent. Knot have ability to scan registered controller for FnContent function and autoregister them

Below code will define controller called as Hello with 3 functions: Morning, Evening and Night. But since Morning and Evening are only function match with FnContent contract, those 2 functions will be registered as RouteHandler.
```go
// here is our controller
type Hello struct{
}

// http://servername/hello/morning
func (h *Hello) Morning(r *knot.WebContext) interface{}{
  return "Good morning"
}

// http://servername/hello/evening
func (h *Hello) Evening(r *knot.WebContext) interface{}{
  return "Good evening"
}

func (h *Hello) Night() interface{}{
  return "Good Night"
}

func main(){
  ks := new(knot.Server)
  ks.Register(new(Hello),"")
  ks.Listen()
}
```

## Start Knot as Multi Application Container
To run knot as Multiple Application Container we need to do following:
- Register application to run in Application Container
- Start the knot server

### appcontainer
```go
package main

import (
  "github.com/eaciit/kingpin"
  "github.com/eaciit/knot/knot.v1"

  // KnotApp Start
  // This is where we need to write down all application namespace 
  // need to be run
  _ "github.com/eaciit/knot/example/hello"
  // KnotApp End
)

var (
  ks          *knot.Server
  flagAddress = kingpin.Flag("address",
    "Address to be used by Knot Server. It normally formatted as SERVERNAME:PORTNUMBER").Default("localhost:9876").String()
)

func main() {
  kingpin.Parse()

  knot.DefaultOutputType = knot.OutputTemplate

  //--- it will run all application registered in namespace
  knot.StartContainer(&knot.AppContainerConfig{
    Address: *flagAddress,
  })
}
```

### create knot application
Now we need to create knot application that will be read by above daemon
```go
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
  r.Config.ViewName = "hello.html"
  return (toolkit.M{}).Set("message", "This data is sent to knot controller method")
}
...
```

## Handling Session
```go
type TestController struct{
}

func (c *TestController) Session(r *knot.WebContext) interface{}{
    t0 := time.Now()
    visitLength := r.Session("VisitLength", 0).(int)
    newVisitLength = visitLength + int(time.Since(t0))
    r.SetSession("VisitLength",newVisitLength)
    return fmt.Sprintf("Old length: %v, new length: %v", time.Duration(visitLenght), time.Duration(newVisitLength))
}
```

## Handling Cookie
```go
type WorldController struct{
}

func (w *WorldController) Cookie(r *knot.WebContext) interface{} {
  r.Config.OutputType = knot.OutputHtml
  cvalue := ""
  cookiename := "mycookie"
  c, _ := r.Cookie(cookiename, "")
  if c == nil {
    r.SetCookie(cookiename, "Arief Darmawan ", 30*24*time.Hour)
  } else {
    cvalue = c.Value
    c.Value = "Arief Darmawan " + time.Now().String()
    c.Expires = time.Now().Add(24 * 30 * time.Hour)
  }
  return "Cookie value is " + cvalue
}
```

## Undocumented Feature
Below feature are available already on Knot, but not documented properly yet

- Template
- Json 

