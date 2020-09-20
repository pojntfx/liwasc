package main

import (
	"github.com/maxence-charriere/go-app/v7/pkg/app"
	"github.com/pojntfx/liwasc-frontend-web/pkg/components"
)

func main() {
	app.Route("/", &components.AppComponent{
		CurrentUserEmail:       "felix@pojtinger.com",
		CurrentUserDisplayName: "Felix Pojtinger",
		DetailsOpen:            false,
		ServicesOpen:           false,
		SelectedNode:           -1,
		SelectedService:        -1,
		UserMenuOpen:           false,
	})

	app.Run()
}
