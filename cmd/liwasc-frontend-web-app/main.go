package main

import (
	"github.com/maxence-charriere/go-app/v7/pkg/app"
	"github.com/pojntfx/liwasc-frontend-web/pkg/components"
)

func main() {
	app.Route("/", &components.AppComponent{})

	app.Run()
}
