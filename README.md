# knot
Knot is golang application container web server. It still on experimental version

# Background
I've been working with Golang for sometime. While some stack are build on either command line or REST, I always use web based application as main UI. 
Most of the time out application will be proxied by Nginx, and to be honest, it complicated bit, because we have to change nginx config file and restart it whenever we have new application developed. These then inspired me to build Knot.

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
