package components

import (
	"fmt"

	"github.com/maxence-charriere/go-app/v7/pkg/app"
	"github.com/pojntfx/liwasc-frontend-web/pkg/models"
)

type ServiceComponent struct {
	app.Compo

	Service models.Service
}

func (c *ServiceComponent) Render() app.UI {
	return app.Div().Body(
		app.Div().Class("pf-u-mb-md").Text(c.Service.Description),
		app.Dl().Class("pf-c-description-list pf-m-2-col pf-u-mb-md").Body(
			&DefinitionComponent{
				Title:   "Assignee",
				Icon:    "fas fa-user-circle",
				Content: &SearchLinkComponent{Topic: c.Service.Assignee},
			},
			&DefinitionComponent{
				Title: "Contact",
				Icon:  "fas fa-paper-plane",
				Content: app.A().
					Href(fmt.Sprintf("mailto:%v", c.Service.Contact)).
					Text(c.Service.Contact),
			},
		),
		app.Hr().Class("pf-c-divider"),
	)
}
