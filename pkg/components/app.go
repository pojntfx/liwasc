package components

import (
	"fmt"

	"github.com/maxence-charriere/go-app/v7/pkg/app"
	"github.com/pojntfx/liwasc-frontend-web/pkg/models"
)

type AppComponent struct {
	app.Compo

	UserMenuOpen bool
	UserAvatar   string
	UserName     string

	Subnets         []string
	Device          string
	NodeSearchValue string

	Nodes        []models.Node
	SelectedNode int

	InspectorOpen        bool
	ServicesAndPortsOpen bool
	DetailsOpen          bool

	ServicesOpen         bool
	SelectedService      int
	InspectorSearchValue string
}

func (c *AppComponent) Render() app.UI {
	return app.Div().Class("pf-c-page").Body(
		&NavbarComponent{
			UserMenuOpen: c.UserMenuOpen,
			UserAvatar:   c.UserAvatar,
			UserName:     c.UserName,

			OnUserMenuToggle: func(ctx app.Context, e app.Event) {
				c.handleUserMenuToggle()
			},
			OnSignOutClick: func(ctx app.Context, e app.Event) {
				c.handleSignOutClick()
			},
		},
		app.Main().Class("pf-c-page__main").Body(
			app.Section().Class("pf-c-page__main-section pf-m-no-padding").Body(
				&DrawerComponent{
					Open: c.InspectorOpen,
					Title: app.If(
						c.ServicesOpen,
						app.Div().Body(
							app.Button().Class("pf-u-mr-md pf-c-button pf-m-plain").Body(
								app.I().Class("fas fa-arrow-left"),
							).OnClick(func(ctx app.Context, e app.Event) {
								c.handleInspectorBackClick()
							}),
							app.B().Text(fmt.Sprintf("Service %v", func() string {
								if c.SelectedService == -1 {
									return ""
								}

								return c.Nodes[c.SelectedNode].Services[c.SelectedService].ServiceName
							}())),
						),
					).Else(
						app.B().Text(fmt.Sprintf("Node %v", func() string {
							if c.SelectedNode == -1 {
								return ""
							}

							return c.Nodes[c.SelectedNode].MACAddress
						}())),
					),
					Main: app.Div().Body(
						&ToolbarComponent{
							Subnets:     []string{"10.0.0.0/9", "192.168.0.0/27"},
							Device:      "eth0",
							SearchValue: c.NodeSearchValue,

							OnSearchChange: func(newSearchValue string) {
								c.handleNodeSearchChange(newSearchValue)
							},
							OnTriggerClick: func(ctx app.Context, e app.Event) {
								c.handleNodeTriggerClick()
							},
						},
						&TableComponent{
							Nodes:        c.Nodes,
							SelectedNode: c.SelectedNode,

							OnRowClick: func(i int) {
								c.handleRowClick(i)
							},
							OnPowerToggle: func(i int) {
								c.handlePowerToggle(i)
							},
						}),
					Details: app.If(
						c.ServicesOpen,
						&ServiceInspectorComponent{
							Service: func() models.Service {
								if c.SelectedService == -1 {
									return models.Service{}
								}

								return c.Nodes[c.SelectedNode].Services[c.SelectedService]
							}(),
						},
					).
						Else(
							app.Div().Body(
								&NodeInspectorComponent{
									Node: func() models.Node {
										if c.SelectedNode == -1 {
											return models.Node{}
										}

										return c.Nodes[c.SelectedNode]
									}(),
									ServicesAndPortsOpen: c.ServicesAndPortsOpen,
									DetailsOpen:          c.DetailsOpen,
									SearchValue:          c.InspectorSearchValue,

									OnServicesAndPortsToggle: func(ctx app.Context, e app.Event) {
										c.handleServiceAndPortsToggle()
									},
									OnDetailsToggle: func(ctx app.Context, e app.Event) {
										c.handleDetailsToggle()
									},
									OnSearchChange: func(newSearchValue string) {
										c.handleInspectorSearchChange(newSearchValue)
									},
									OnReScanClick: func(ctx app.Context, e app.Event) {
										c.handleReScanClick()
									},
									OnServiceClick: func(i int) {
										c.handleServiceClick(i)
									},
								}),
						),
					Actions: []app.UI{
						app.If(
							c.ServicesOpen,
							app.Text(fmt.Sprintf("%v/%v",
								func() int {
									if c.SelectedService == -1 {
										return -1
									}

									return c.Nodes[c.SelectedNode].Services[c.SelectedService].PortNumber
								}(),
								func() string {
									if c.SelectedService == -1 {
										return ""
									}

									return c.Nodes[c.SelectedNode].Services[c.SelectedService].TransportProtocol
								}())),
						).Else(
							&OnOffSwitchComponent{
								On: func() bool {
									if c.SelectedNode == -1 {
										return false
									}

									return c.Nodes[c.SelectedNode].PoweredOn
								}(),

								OnToggleClick: func(ctx app.Context, e app.Event) {
									c.handlePowerToggle(c.SelectedNode)
								},
							},
						),
						app.Button().Class("pf-u-ml-md pf-c-button pf-m-plain").Body(
							app.I().Class("fas fa-times"),
						).OnClick(func(ctx app.Context, e app.Event) {
							c.handleInspectorCloseClick()
						}),
					},
				},
			),
		),
	)
}

func (c *AppComponent) handleUserMenuToggle() {
	c.UserMenuOpen = !c.UserMenuOpen

	c.Update()
}

func (c *AppComponent) handleSignOutClick() {
	c.UserMenuOpen = false

	c.Update()
}

func (c *AppComponent) handleNodeSearchChange(newSearchValue string) {
	c.NodeSearchValue = newSearchValue

	c.Update()
}

func (c *AppComponent) handleNodeTriggerClick() {

}

func (c *AppComponent) handleRowClick(i int) {
	if c.SelectedNode == i {
		c.handleInspectorCloseClick()
	} else {
		c.InspectorOpen = true
		c.ServicesOpen = false
		c.SelectedNode = i
		c.SelectedService = -1
	}

	c.Update()
}

func (c *AppComponent) handlePowerToggle(i int) {
	c.Nodes[i].PoweredOn = !c.Nodes[i].PoweredOn

	c.Update()
}

func (c *AppComponent) handleInspectorCloseClick() {
	c.InspectorOpen = false
	c.ServicesOpen = false
	c.SelectedNode = -1
	c.SelectedService = -1

	c.Update()
}

func (c *AppComponent) handleServiceAndPortsToggle() {
	c.ServicesAndPortsOpen = !c.ServicesAndPortsOpen

	c.Update()
}

func (c *AppComponent) handleDetailsToggle() {
	c.DetailsOpen = !c.DetailsOpen

	c.Update()
}

func (c *AppComponent) handleInspectorSearchChange(newSearchValue string) {
	c.InspectorSearchValue = newSearchValue

	c.Update()
}

func (c *AppComponent) handleReScanClick() {

}

func (c *AppComponent) handleServiceClick(i int) {
	c.SelectedService = i
	c.ServicesOpen = true

	c.Update()
}

func (c *AppComponent) handleInspectorBackClick() {
	c.SelectedService = -1
	c.ServicesOpen = false

	c.Update()
}
