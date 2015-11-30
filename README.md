# knot
Knot is a web server and application container for Golang Web Based App. It still on experimental version

# Background
I've been working with Golang for sometime. While some stack are build on either command line or REST, I always use web based application as main UI. 
Most of the time our application will be proxied by Nginx, and to be honest, it is complicated effort, because we have to change nginx config file and restart it whenever we have new application developed. These then inspired me to build Knot.

# Usage
Load knot
```go
go get -u github.com/eaciit/knot
```
## FnContent
In knot, we handle function based on FnContent contract. Where contract of FnContent is as follow
```go
var FnContent func(*knot.Request)interface{}
```
knot.Request is wrapper of KnotServer object, HttpRequest and other related context of Knot. Any data returned by this function will be used as body of the response.

## Route 
Route to a function handler
```go
func main(){
  ks := new(knot.Server)
  ks.Route("hi",Hi)
  ks.Listen()
}

func Hi(r *knot.Request)interface{}{
  return "Hello World!"
}
```

Route to static folder
```go
ks.RouteStatic("static","/Users/ariefdarmawan/Temp/knot/app1/static")
```

## Register a controller
Controller is a struct with sets of FnContent. Knot have ability to scan registered controller for FnContent function and autoregister them

Below code will define controller called as Hello with 3 functions: Morning, Evening and Night. But since Morning and Evening are only function match with FnContent contract, those 2 functions will be registered as RouteHandler.
```go
// here is our controller
type Hello struct{
}

// http://servername/hello/morning
func (h *Hello) Morning(r *knot.Request) interface{}{
  return "Good morning"
}

// http://servername/hello/evening
func (h *Hello) Evening(r *knot.Request) interface{}{
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

## App Container
By applying app container, we can host many go web based application and run it within a single instance of web server.

```go
package main

import (
  // KnotApp Start
  // include all namespace for application to be listed
  _ "github.com/eaciit/knot/example/hello"
  // KnotApp End
)

....

func main() {
  knot.DefaultOutputType = knot.OutputTemplate
  appcontainer.Start(&appcontainer.Config{
    Address: "localhost:9876",
  })
}
```

now we need to create knot application that will be read by above daemon
```go
package hello

import (
  "github.com/eaciit/knot"
  "github.com/eaciit/knot/appcontainer"
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

//--- THIS FUNCTION IS IMPORTANT
func init() {
  app := appcontainer.NewApp("Hello")
  app.ViewsPath = appViewsPath
  app.Register(&WorldController{})
  app.Static("static", "/Users/ariefdarmawan/Temp")
  app.LayoutTemplate = "_template.html"
  appcontainer.RegisterApp(app)
}

type WorldController struct {
}

// Will be autoregistered as http://appserver/hello/world/index
// It will automatically read content from /views/world/index.html
func (w *WorldController) Index(r *knot.Request) interface{} {
  // r.ResponseConfig().ViewName = "someother_template.html" 
  // unmark above line to change viewname
  return struct{Message string}{"Hello from Knot"} 
}
```

## Handling Session

```go
type TestController struct{
}

func (c *TestController) Session(r *knot.Request) interface{}{
    t0 := time.Now()
    visitLength := r.Session("VisitLength", 0).(int)
    newVisitLength = visitLength + int(time.Since(t0))
    r.SetSession("VisitLength",newVisitLength)
    return fmt.Sprintf("Old length: %v, new length: %v", time.Duration(visitLenght), time.Duration(newVisitLength))
}
```

## Undocumented Feature
Below feature are available already on Knot, but not yet documented properly yet

- Template
- Json 
- Multi application

