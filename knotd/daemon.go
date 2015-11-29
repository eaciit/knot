package main

import (
	"github.com/eaciit/kingpin"
	"github.com/eaciit/knot"
	"github.com/eaciit/knot/appcontainer"

	// KnotApp Start
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
	appcontainer.Start(&appcontainer.Config{
		Address: *flagAddress,
	})
}
