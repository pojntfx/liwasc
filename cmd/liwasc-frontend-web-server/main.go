package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/maxence-charriere/go-app/v7/pkg/app"
)

func main() {
	listenAddress := flag.String("listenAddress", "0.0.0.0:7000", "Listen address")

	flag.Parse()

	h := app.Handler{
		Title:  "liwasc",
		Author: "Felicitas Pojtinger",
		Styles: []string{
			"https://unpkg.com/@patternfly/patternfly@4.42.2/patternfly.css",
			"https://unpkg.com/@patternfly/patternfly@4.42.2/patternfly-addons.css",
			"/web/index.css",
		},
	}

	log.Println("Listening on", *listenAddress)

	if err := http.ListenAndServe(*listenAddress, &h); err != nil {
		log.Fatal("could not start server", err)
	}
}
