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
		Author:          "Felix Pojtinger",
		BackgroundColor: "#151515",
		Description:     "List, wake and scan nodes in a network.",
		Icon: app.Icon{
			Default: "/web/icon.png",
		},
		Keywords: []string{
			"network",
			"network-scanner",
			"port-scanner",
			"ip-scanner",
			"arp-scanner",
			"arp",
			"iana",
			"ports2packets",
			"liwasc",
			"vendor2mac",
			"wake-on-lan",
			"wol",
			"service-name",
		},
		LoadingLabel: "List, wake and scan nodes in a network.",
		Name:         "liwasc",
		RawHeaders: []string{
			`<meta property="og:url" content="https://liwasc.felix.pojtinger.com">`,
			`<meta property="og:title" content="liwasc">`,
			`<meta property="og:description" content="List, wake and scan nodes in a network.">`,
			`<meta property="og:image" content="https://liwasc.felix.pojtinger.com/web/icon.png">`,
		},
		Styles: []string{
			"https://unpkg.com/@patternfly/patternfly@4.42.2/patternfly.css",
			"https://unpkg.com/@patternfly/patternfly@4.42.2/patternfly-addons.css",
			"/web/index.css",
		},
		ThemeColor: "#151515",
		Title:      "liwasc",
	}

	log.Println("Listening on", *listenAddress)

	if err := http.ListenAndServe(*listenAddress, &h); err != nil {
		log.Fatal("could not start server", err)
	}
}
