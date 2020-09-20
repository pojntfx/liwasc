package components

import (
	"fmt"

	"github.com/maxence-charriere/go-app/v7/pkg/app"
	"github.com/pojntfx/liwasc-frontend-web/pkg/models"
)

type NodeComponent struct {
	app.Compo

	Node                 models.Node
	servicesAndPortsOpen bool
	detailsOpen          bool
}

func (c *NodeComponent) Render() app.UI {
	return app.Div().Body(
		app.Dl().Class("pf-c-description-list pf-m-2-col pf-u-mb-md").Body(
			&DefinitionComponent{
				Title:   "IP Address",
				Icon:    "fas fa-globe",
				Content: app.Text(c.Node.IPAddress),
			},
			&DefinitionComponent{
				Title: "Vendor",
				Icon:  "fas fa-store-alt",
				Content: app.A().
					Href(fmt.Sprintf("https://duckduckgo.com/?q=%v", c.Node.Vendor)).
					Target("_blank").
					Body(
						app.I().Class("fas fa-external-link-alt pf-u-mr-xs"),
						app.Text(c.Node.Vendor),
					),
			},
		),
		&ExpandableSectionComponent{
			Open:     c.servicesAndPortsOpen,
			OnToggle: c.handleServicesAndPortsOpen,
			Title:    "Services and Ports",
			Content:  app.Text("Example services and ports content"),
		},
		&ExpandableSectionComponent{
			Open:     c.detailsOpen,
			OnToggle: c.handleToggleDetailsOpen,
			Title:    "Details",
			Content:  app.Text(c.Node.Vendor),
		},
	)
}

func (c *NodeComponent) OnMount(ctx app.Context) {
	c.servicesAndPortsOpen = true
	c.detailsOpen = true

	c.Update()
}

func (c *NodeComponent) handleServicesAndPortsOpen(ctx app.Context, e app.Event) {
	c.servicesAndPortsOpen = !c.servicesAndPortsOpen

	c.Update()
}

func (c *NodeComponent) handleToggleDetailsOpen(ctx app.Context, e app.Event) {
	c.detailsOpen = !c.detailsOpen

	c.Update()
}
