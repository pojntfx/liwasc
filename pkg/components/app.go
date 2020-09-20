package components

import (
	"fmt"

	"github.com/maxence-charriere/go-app/v7/pkg/app"
	"github.com/pojntfx/liwasc-frontend-web/pkg/models"
)

type AppComponent struct {
	app.Compo

	CurrentUserEmail       string
	CurrentUserDisplayName string
	DetailsOpen            bool
	SelectedNode           int
	SelectedService        int
	ServicesOpen           bool
	UserMenuOpen           bool
	Nodes                  []models.Node
}

func (c *AppComponent) Render() app.UI {
	return app.Div().Class("pf-c-page").Body(
		&NavbarComponent{
			CurrentUserEmail:       c.CurrentUserEmail,
			CurrentUserDisplayName: c.CurrentUserDisplayName,
			UserMenuOpen:           c.UserMenuOpen,
			OnUserMenuToggle:       func(ctx app.Context, e app.Event) { c.handleUserMenuToggle() },
		},
		app.Main().Class("pf-c-page__main").Body(
			app.Section().Class("pf-c-page__main-section pf-m-no-padding").Body(
				&DetailsComponent{
					Open: c.DetailsOpen,
					Title: app.If(!c.ServicesOpen,
						app.B().Text(fmt.Sprintf("Node %v", func() string {
							if c.SelectedNode == -1 {
								return ""
							}

							return c.Nodes[c.SelectedNode].MACAddress
						}()))).Else(
						app.Div().Body(
							app.Button().Class("pf-u-mr-md pf-c-button pf-m-plain").Body(
								app.I().Class("fas fa-arrow-left"),
							).OnClick(func(ctx app.Context, e app.Event) { c.handleServicesClose() }),
							app.B().Text(fmt.Sprintf("Service %v", func() string {
								if c.SelectedService == -1 {
									return ""
								}

								return c.Nodes[c.SelectedNode].Services[c.SelectedService].ServiceName
							}())),
						),
					),
					Main: app.Div().Body(
						&FilterComponent{Subnets: []string{"10.0.0.0/9", "192.168.0.0/27"}, Device: "eth0"},
						&ListingComponent{
							OnRowClick:        c.handleDetailsOpen,
							SelectedNode:      c.SelectedNode,
							Nodes:             c.Nodes,
							OnNodePowerToggle: c.handleNodePowerToggle,
						}),
					Details: app.If(!c.ServicesOpen, app.Div().Body(
						&NodeComponent{
							Node: func() models.Node {
								if c.SelectedNode == -1 {
									return models.Node{}
								}

								return c.Nodes[c.SelectedNode]
							}(),
							OnOpenService: c.handleServicesOpen}),
					).
						Else(
							&ServiceComponent{
								Service: func() models.Service {
									if c.SelectedService == -1 {
										return models.Service{}
									}

									return c.Nodes[c.SelectedNode].Services[c.SelectedService]
								}(),
							}),
					Actions: []app.UI{
						app.If(!c.ServicesOpen, &OnOffSwitchComponent{
							On: func() bool {
								if c.SelectedNode == -1 {
									return false
								}

								return c.Nodes[c.SelectedNode].PoweredOn
							}(),
							OnToggle: func(ctx app.Context, e app.Event) { c.handleNodePowerToggle(c.SelectedNode) },
						}).Else(
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
						),
						app.Button().Class("pf-u-ml-md pf-c-button pf-m-plain").Body(
							app.I().Class("fas fa-times"),
						).OnClick(func(ctx app.Context, e app.Event) { c.handleDetailsClose() }),
					},
				},
			),
		),
	)
}

func (c *AppComponent) handleDetailsOpen(i int) {
	if i == c.SelectedNode {
		if c.DetailsOpen == true {
			c.DetailsOpen = false
			c.SelectedNode = -1
		} else {
			c.DetailsOpen = true
			c.SelectedNode = i
		}
	} else {
		c.DetailsOpen = true
		c.SelectedNode = i
	}

	c.handleServicesClose()
}

func (c *AppComponent) handleDetailsClose() {
	c.DetailsOpen = false
	c.ServicesOpen = false

	c.SelectedNode = -1
	c.SelectedService = -1

	c.Update()
}

func (c *AppComponent) handleNodePowerToggle(i int) {
	c.Nodes[i].PoweredOn = !c.Nodes[i].PoweredOn

	c.Update()
}

func (c *AppComponent) handleServicesOpen(i int) {
	c.ServicesOpen = true
	c.SelectedService = i

	c.Update()
}

func (c *AppComponent) handleServicesClose() {
	c.ServicesOpen = false
	c.SelectedService = -1

	c.Update()
}

func (c *AppComponent) handleUserMenuToggle() {
	c.UserMenuOpen = !c.UserMenuOpen

	c.Update()
}

func (c *AppComponent) OnMount(ctx app.Context) {
	c.Nodes = []models.Node{
		{
			PoweredOn:  false,
			MACAddress: "ff:ff:ff:ff",
			IPAddress:  "10.0.0.1",
			Vendor:     "TP-Link",
			Services: []models.Service{
				{
					ServiceName:             "ssh",
					PortNumber:              22,
					TransportProtocol:       "tcp",
					Description:             "Lorem dolor sit amet",
					Assignee:                "Felix Pojtinger",
					Contact:                 "felix@pojtinger.com",
					RegistrationDate:        "2002-01-01",
					ModificationDate:        "2002-02-02",
					Reference:               "RFC1234",
					ServiceCode:             "C241",
					UnauthorizedUseReported: "Aliens have not registered their protocols running on this port!",
					AssignmentNotes:         "Might glow in the dark.",
				},
				{
					ServiceName:             "dns",
					PortNumber:              53,
					TransportProtocol:       "udp",
					Description:             "Lorem dolor sit amet",
					Assignee:                "Felix Pojtinger",
					Contact:                 "felix@pojtinger.com",
					RegistrationDate:        "2002-01-01",
					ModificationDate:        "2002-02-02",
					Reference:               "RFC1234",
					ServiceCode:             "C241",
					UnauthorizedUseReported: "Aliens have not registered their protocols running on this port!",
					AssignmentNotes:         "Might glow in the dark.",
				},
				{
					ServiceName:             "http",
					PortNumber:              80,
					TransportProtocol:       "tcp",
					Description:             "Lorem dolor sit amet",
					Assignee:                "Felix Pojtinger",
					Contact:                 "felix@pojtinger.com",
					RegistrationDate:        "2002-01-01",
					ModificationDate:        "2002-02-02",
					Reference:               "RFC1234",
					ServiceCode:             "C241",
					UnauthorizedUseReported: "Aliens have not registered their protocols running on this port!",
					AssignmentNotes:         "Might glow in the dark.",
				},
			},
			Registry:     "MA-1",
			Organization: "TP-Link",
			Address:      "One Hacker Way",
			Visible:      true,
		},
		{
			PoweredOn:  true,
			MACAddress: "00:1B:44:11:3A:B7",
			IPAddress:  "10.0.0.2",
			Vendor:     "Realtek",
			Services: []models.Service{
				{
					ServiceName:             "echo",
					PortNumber:              7,
					TransportProtocol:       "tcp",
					Description:             "Lorem dolor sit amet",
					Assignee:                "Felix Pojtinger",
					Contact:                 "felix@pojtinger.com",
					RegistrationDate:        "2002-01-01",
					ModificationDate:        "2002-02-02",
					Reference:               "RFC1234",
					ServiceCode:             "C241",
					UnauthorizedUseReported: "Aliens have not registered their protocols running on this port!",
					AssignmentNotes:         "Might glow in the dark.",
				},
				{
					ServiceName:             "echo",
					PortNumber:              7,
					TransportProtocol:       "udp",
					Description:             "Lorem dolor sit amet",
					Assignee:                "Felix Pojtinger",
					Contact:                 "felix@pojtinger.com",
					RegistrationDate:        "2002-01-01",
					ModificationDate:        "2002-02-02",
					Reference:               "RFC1234",
					ServiceCode:             "C241",
					UnauthorizedUseReported: "Aliens have not registered their protocols running on this port!",
					AssignmentNotes:         "Might glow in the dark.",
				},
				{
					ServiceName:             "http",
					PortNumber:              80,
					TransportProtocol:       "tcp",
					Description:             "Lorem dolor sit amet",
					Assignee:                "Felix Pojtinger",
					Contact:                 "felix@pojtinger.com",
					RegistrationDate:        "2002-01-01",
					ModificationDate:        "2002-02-02",
					Reference:               "RFC1234",
					ServiceCode:             "C241",
					UnauthorizedUseReported: "Aliens have not registered their protocols running on this port!",
					AssignmentNotes:         "Might glow in the dark.",
				},
			},
			Registry:     "MA-1",
			Organization: "Realtek",
			Address:      "Two Hacker Way",
			Visible:      false,
		},
	}

	c.Update()
}
