# knot
Knot is a web server and application container for Golang Web Based App. It still on experimental version

# Background
I've been working with Golang for sometime. While some stack are build on either command line or REST, I always use web based application as main UI. 
Most of the time our application will be proxied by Nginx, and to be honest, it is complicated effort, because we have to change nginx config file and restart it whenever we have new application developed. These then inspired me to build Knot.

# Usage
Load knot
```
go get -u github.com/eaciit/knot
```
## FnContent
In knot, we handle function based on FnContent contract. Where contract of FnContent is as follow
```
var FnContent func(*knot.Request)interface{}
```
knot.Request is wrapper of KnotServer object, HttpRequest and other related context of Knot. Any data returned by this function will be used as body of the response.

## Route 
Route to a function handler
```
func main(){
  ks := new(knot.Server)
  ks.Route("hi",Hi)
  ks.Listen()
}

func Hi(r *knot.Request)interaface{}{
  return "Hello World!"
}
```

Route to static folder
```
ks.RouteStatic("static","/Users/ariefdarmawan/Temp/knot/app1/static")
```

## Register a controller
Controller is a struct with sets of FnContent. Knot have ability to scan registered controller for FnContent function and autoregister them

Below code will define controller called as Hello with 3 functions: Morning, Evening and Night. But since Morning and Evening are only function match with FnContent contract, those 2 functions will be registered as RouteHandler.
```
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

## Undocumented Feature
Below feature are available already on Knot, but not yet documented properly yet

- Cookie
- Template
- Json 
- Multi application

