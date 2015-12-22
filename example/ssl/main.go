package main

import (
	"github.com/eaciit/knot/knot.v1"
	"os"
	"os/exec"
	"time"
)

type Hello struct {
}

func (h *Hello) Index(r *knot.WebContext) interface{} {
	return "Accessing /index using SSL enabled"
}

func main() {
	basepath, _ := os.Getwd()

	knot.DefaultOutputType = knot.OutputHtml
	ks := new(knot.Server)
	ks.Address = ":1234"
	ks.UseSSL = true

	// use golang.org/src/pkg/crypto/tls/generate_cert.go to generate sample certificate and private key
	// `go run generate_cert.go --host localhost:1234`
	// certificate and private key file location need to be full path
	ks.CertificatePath = basepath + "/cert.pem"
	ks.PrivateKeyPath = basepath + "/key.pem"

	ks.Register(new(Hello), "")

	time.AfterFunc(time.Second, func() {
		url := "https://localhost" + ks.Address + "/hello/index"
		exec.Command("open", url).Run()
	})

	ks.Listen()
}
