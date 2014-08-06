package main

import (
	"flag"
	"log"

	"github.com/oremj/go-freddo/freddo"
	"github.com/zenazn/goji"
)

var config = flag.String("config", "", "Config location.")

func main() {
	flag.Parse()

	if *config == "" {
		log.Fatal("-config must be specified.")
	}

	freddo, err := freddo.NewFreddo(*config)
	if err != nil {
		log.Fatal(err)
	}

	goji.Post("/update/:appname", freddo.UpdateApp)
	goji.Serve()
}
