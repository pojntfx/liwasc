package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/kataras/compress"
	"github.com/maxence-charriere/go-app/v8/pkg/app"
	"github.com/pojntfx/liwasc/pkg/components"
)

func main() {
	// Client-side code
	{
		// Define the routes
		app.Route("/", &components.Home{})

		// Start the app
		app.RunWhenOnBrowser()
	}

	// Server-/build-side code
	{
		// Parse the flags
		build := flag.Bool("build", false, "Create static build")
		out := flag.String("out", "out/liwasc-frontend", "Out directory for static build")
		path := flag.String("path", "", "Base path for static build")
		serve := flag.Bool("serve", false, "Build and serve the frontend")
		laddr := flag.String("laddr", "localhost:15125", "Address to serve the frontend on")

		flag.Parse()

		// Define the handler
		h := &app.Handler{
			Author:          "Felicitas Pojtinger",
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
				`<meta property="og:url" content="https://pojntfx.github.io/liwasc/">`,
				`<meta property="og:title" content="liwasc">`,
				`<meta property="og:description" content="List, wake and scan nodes in a network.">`,
				`<meta property="og:image" content="https://pojntfx.github.io/liwasc/web/icon.png">`,
			},
			Styles: []string{
				`https://unpkg.com/@patternfly/patternfly@4.96.2/patternfly.css`,
				`https://unpkg.com/@patternfly/patternfly@4.96.2/patternfly-addons.css`,
				`/web/index.css`,
			},
			ThemeColor: "#151515",
			Title:      "liwasc",
		}

		// Create static build if specified
		if *build {
			// Deploy under a path
			if *path != "" {
				h.Resources = app.GitHubPages(*path)
			}

			if err := app.GenerateStaticWebsite(*out, h); err != nil {
				log.Fatalf("could not build: %v\n", err)
			}
		}

		// Serve if specified
		if *serve {
			log.Printf("liwasc frontend listening on %v\n", *laddr)

			if err := http.ListenAndServe(*laddr, compress.Handler(h)); err != nil {
				log.Fatalf("could not open liwasc frontend: %v\n", err)
			}
		}
	}
}
