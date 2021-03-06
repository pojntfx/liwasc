package components

import (
	"fmt"

	"github.com/maxence-charriere/go-app/v7/pkg/app"
	"github.com/pojntfx/liwasc-frontend-web/pkg/models"
)

type AppComponent struct {
	app.Compo

	userMenuOpen bool
	UserAvatar   string
	UserName     string

	Logout func()

	Subnets         []string
	Device          string
	NodeSearchValue string

	Nodes        []*models.Node
	selectedNode int

	inspectorOpen        bool
	servicesAndPortsOpen bool
	detailsOpen          bool

	servicesOpen         bool
	selectedService      int
	InspectorSearchValue string

	Connected bool
	Scanning  bool

	TriggerNodeScan func()
}

func (c *AppComponent) Render() app.UI {
	if c.Nodes == nil {
		c.userMenuOpen = false
		c.selectedNode = -1

		c.inspectorOpen = false
		c.servicesAndPortsOpen = true
		c.detailsOpen = false

		c.servicesOpen = false
		c.selectedService = -1
	}

	return app.Div().TabIndex(0).Class("pf-c-page").Body(
		&NavbarComponent{
			UserMenuOpen: c.userMenuOpen,
			UserAvatar:   c.UserAvatar,
			UserName:     c.UserName,

			Connected: c.Connected,

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
					Open: c.inspectorOpen,
					Title: app.If(
						c.servicesOpen,
						app.Div().Body(
							app.Button().Class("pf-u-mr-md pf-c-button pf-m-plain").Body(
								app.I().Class("fas fa-arrow-left"),
							).OnClick(func(ctx app.Context, e app.Event) {
								c.handleInspectorBackClick()
							}),
							app.B().Text(fmt.Sprintf("Service %v", func() string {
								if c.selectedService == -1 {
									return ""
								}

								serviceName := c.Nodes[c.selectedNode].Services[c.selectedService].ServiceName

								if serviceName == "" {
									return "Non-Registered service"
								}

								return serviceName
							}())),
						),
					).Else(
						app.B().Text(fmt.Sprintf("Node %v", func() string {
							if c.selectedNode == -1 {
								return ""
							}

							return c.Nodes[c.selectedNode].MACAddress
						}())),
					),
					Main: app.Div().Body(
						&ToolbarComponent{
							Subnets:     c.Subnets,
							Device:      c.Device,
							SearchValue: c.NodeSearchValue,

							Scanning: c.Scanning,

							OnSearchChange: func(newSearchValue string) {
								c.handleNodeSearchChange(newSearchValue)
							},
							OnTriggerClick: func(ctx app.Context, e app.Event) {
								c.TriggerNodeScan()
							},
						},
						&TableComponent{
							Nodes:        c.Nodes,
							SelectedNode: c.selectedNode,

							OnRowClick: func(i int) {
								c.handleRowClick(i)
							},
							OnPowerToggle: func(i int) {
								c.handlePowerToggle(i)
							},
						}),
					Details: app.If(
						c.servicesOpen,
						&ServiceInspectorComponent{
							Service: func() *models.Service {
								if c.selectedService == -1 {
									return &models.Service{}
								}

								return c.Nodes[c.selectedNode].Services[c.selectedService]
							}(),
						},
					).
						Else(
							app.Div().Body(
								&NodeInspectorComponent{
									Node: func() *models.Node {
										if c.selectedNode == -1 {
											return &models.Node{}
										}

										return c.Nodes[c.selectedNode]
									}(),
									ServicesAndPortsOpen: c.servicesAndPortsOpen,
									DetailsOpen:          c.detailsOpen,
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
							c.servicesOpen,
							app.Text(fmt.Sprintf("%v/%v",
								func() int {
									if c.selectedService == -1 {
										return -1
									}

									return c.Nodes[c.selectedNode].Services[c.selectedService].PortNumber
								}(),
								func() string {
									if c.selectedService == -1 {
										return ""
									}

									return c.Nodes[c.selectedNode].Services[c.selectedService].TransportProtocol
								}())),
						).Else(
							&OnOffSwitchComponent{
								On: func() bool {
									if c.selectedNode == -1 {
										return false
									}

									return c.Nodes[c.selectedNode].PoweredOn
								}(),

								OnToggleClick: func(ctx app.Context, e app.Event) {
									c.handlePowerToggle(c.selectedNode)
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
	c.userMenuOpen = !c.userMenuOpen

	c.Update()
}

func (c *AppComponent) handleSignOutClick() {
	c.userMenuOpen = false

	c.Update()

	c.Logout()
}

func (c *AppComponent) handleNodeSearchChange(newSearchValue string) {
	c.NodeSearchValue = newSearchValue

	c.Update()
}

func (c *AppComponent) handleNodeTriggerClick() {

}

func (c *AppComponent) handleRowClick(i int) {
	if c.selectedNode == i {
		c.handleInspectorCloseClick()
	} else {
		c.inspectorOpen = true
		c.servicesOpen = false
		c.selectedNode = i
		c.selectedService = -1
	}

	c.Update()
}

func (c *AppComponent) handlePowerToggle(i int) {
	c.Nodes[i].PoweredOn = !c.Nodes[i].PoweredOn

	c.Update()
}

func (c *AppComponent) handleInspectorCloseClick() {
	c.inspectorOpen = false
	c.servicesOpen = false
	c.selectedNode = -1
	c.selectedService = -1

	c.Update()
}

func (c *AppComponent) handleServiceAndPortsToggle() {
	c.servicesAndPortsOpen = !c.servicesAndPortsOpen

	c.Update()
}

func (c *AppComponent) handleDetailsToggle() {
	c.detailsOpen = !c.detailsOpen

	c.Update()
}

func (c *AppComponent) handleInspectorSearchChange(newSearchValue string) {
	c.InspectorSearchValue = newSearchValue

	c.Update()
}

func (c *AppComponent) handleReScanClick() {

}

func (c *AppComponent) handleServiceClick(i int) {
	c.selectedService = i
	c.servicesOpen = true

	c.Update()
}

func (c *AppComponent) handleInspectorBackClick() {
	c.selectedService = -1
	c.servicesOpen = false

	c.Update()
}
