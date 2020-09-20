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
				Title:   "Vendor",
				Icon:    "fas fa-store-alt",
				Content: &SearchLinkComponent{Topic: c.Node.Vendor},
			},
		),
		&ExpandableSectionComponent{
			Open:     c.servicesAndPortsOpen,
			OnToggle: c.handleServicesAndPortsOpen,
			Title:    "Services and Ports",
			Content: app.Raw(`<ul class="pf-c-data-list">
  <li class="pf-c-data-list__item pf-m-selectable">
    <div class="pf-c-data-list__item-row">
      <div class="pf-c-data-list__item-content">
		  <div class="pf-c-data-list__cell pf-u-display-flex pf-u-justify-content-space-between">
	          <span>Service http</span>
			  <span class="pf-u-ml-md">80/tcp</span>
		  </div>
	  </div>
    </div>
  </li>

<li class="pf-c-data-list__item pf-m-selectable">
    <div class="pf-c-data-list__item-row">
      <div class="pf-c-data-list__item-content">
        <div class="pf-c-data-list__cell pf-u-display-flex pf-u-justify-content-space-between">
          <span>Service dns</span>
		  <span class="pf-u-ml-md">53/udp</span>
        </div>
      </div>
    </div>
  </li>
</ul>`),
		},
		&ExpandableSectionComponent{
			Open:     c.detailsOpen,
			OnToggle: c.handleToggleDetailsOpen,
			Title:    "Details",
			Content: app.Dl().Class("pf-c-description-list pf-m-2-col pf-u-mb-md").Body(
				&DefinitionComponent{
					Title:   "Registry",
					Icon:    "fas fa-list",
					Content: app.Text(c.Node.Registry),
				},
				&DefinitionComponent{
					Title:   "Organization",
					Icon:    "fas fa-university",
					Content: &SearchLinkComponent{Topic: c.Node.Organization},
				},
				&DefinitionComponent{
					Title:   "Registered Address",
					Icon:    "fas fa-map-marker-alt",
					Content: &SearchLinkComponent{Topic: c.Node.Address},
				},
				&DefinitionComponent{
					Title: "Visible Address",
					Icon:  "fas fa-binoculars",
					Content: app.Div().Body(
						app.I().Class(fmt.Sprintf("fas %v pf-u-mr-xs", func() string {
							if c.Node.Visible {
								return "fas fa-eye"
							}

							return "fas fa-eye-slash"
						}())),
						app.Text(func() string {
							if c.Node.Visible {
								return "Visible"
							}

							return "Hidden"
						}()),
					),
				},
			),
		},
	)
}

func (c *NodeComponent) OnMount(ctx app.Context) {
	c.servicesAndPortsOpen = true
	c.detailsOpen = false

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
