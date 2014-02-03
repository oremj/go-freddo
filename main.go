package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/codegangsta/martini"
	"github.com/oremj/go-freddo/freddo"
)

var addr = flag.String("addr", ":8080", "Bind address.")
var config = flag.String("config", "", "config location")

func main() {
	flag.Parse()

	if *config == "" {
		log.Fatal("-config must be specified.")
	}

	freddo, err := freddo.NewFreddo(*config)
	if err != nil {
		log.Fatal(err)
	}

	m := martini.Classic()
	m.Post(`/update/(?P<app_name>\w+?)/`, freddo.UpdateApp)
	log.Fatal(http.ListenAndServe(*addr, m))
}
