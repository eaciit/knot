package main

import (
	"github.com/eaciit/knot/knot.v1"
	"os"
	"os/exec"
	"time"
)

// use golang.org/src/pkg/crypto/tls/generate_cert.go to generate sample certificate and private key
// `go run generate_cert.go --host localhost:1234`
// certificate and private key file location need to be full path

func main() {
	basepath, _ := os.Getwd()

	app := knot.NewApp("startapp using ssl")
	app.Register(&Hello{})

	app.UseSSL = true
	app.CertificatePath = basepath + "/cert.pem"
	app.PrivateKeyPath = basepath + "/key.pem"

	time.AfterFunc(time.Second, func() {
		url := "https://localhost:1234/hello/index"
		exec.Command("open", url).Run()
	})

	knot.RegisterApp(app)
	knot.StartApp(app, "localhost:1234")
}

type Hello struct {
}

func (h *Hello) Index(r *knot.WebContext) interface{} {
	r.Config.OutputType = knot.OutputHtml
	return "Accessing /index using SSL enabled from knot.StartApp"
}
