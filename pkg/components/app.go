package components

import (
	"fmt"

	"github.com/maxence-charriere/go-app/v7/pkg/app"
	"github.com/pojntfx/liwasc-frontend-web/pkg/models"
)

type AppComponent struct {
	app.Compo
	DetailsOpen  bool
	SelectedNode int
	Nodes        []models.Node
}

func (c *AppComponent) Render() app.UI {
	return app.Div().Body(
		&FilterComponent{Subnets: []string{"10.0.0.0/9", "192.168.0.0/27"}, Device: "eth0"},
		&DetailsComponent{
			Open: c.DetailsOpen,
			Title: fmt.Sprintf("Node %v", func() string {
				if c.SelectedNode == -1 {
					return ""
				}

				return c.Nodes[c.SelectedNode].MACAddress
			}()),
			Main: &ListingComponent{
				OnRowClick:        c.handleDetailsOpen,
				SelectedNode:      c.SelectedNode,
				Nodes:             c.Nodes,
				OnNodePowerToggle: c.handleNodePowerToggle,
			},
			Details: app.Div().Body(
				&NodeComponent{
					Node: func() models.Node {
						if c.SelectedNode == -1 {
							return models.Node{}
						}

						return c.Nodes[c.SelectedNode]
					}()},
			),
			Actions: []app.UI{
				&OnOffSwitchComponent{
					On: func() bool {
						if c.SelectedNode == -1 {
							return false
						}

						return c.Nodes[c.SelectedNode].PoweredOn
					}(),
					OnToggle: func(ctx app.Context, e app.Event) { c.handleNodePowerToggle(c.SelectedNode) },
				},
				app.Button().Class("pf-u-ml-md pf-c-button pf-m-plain").Body(
					app.I().Class("fas fa-times"),
				).OnClick(func(ctx app.Context, e app.Event) { c.handleDetailsClose() }),
			},
		},
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

	c.Update()
}

func (c *AppComponent) handleDetailsClose() {
	c.DetailsOpen = false
	c.SelectedNode = -1

	c.Update()
}

func (c *AppComponent) handleNodePowerToggle(i int) {
	c.Nodes[i].PoweredOn = !c.Nodes[i].PoweredOn

	c.Update()
}

func (c *AppComponent) OnMount(ctx app.Context) {
	c.Nodes = []models.Node{
		{
			PoweredOn:  false,
			MACAddress: "ff:ff:ff:ff",
			IPAddress:  "10.0.0.1",
			Vendor:     "TP-Link",
			ServicesAndPorts: []string{
				"22/tcp (ssh)",
				"53/udp (dns)",
				"80/tcp (http)",
			},
		},
		{
			PoweredOn:  true,
			MACAddress: "00:1B:44:11:3A:B7",
			IPAddress:  "10.0.0.2",
			Vendor:     "Realtek",
			ServicesAndPorts: []string{
				"7/tcp (echo)",
				"7/tcp (echo)",
				"80/tcp (http)",
			},
		},
	}

	c.Update()
}
