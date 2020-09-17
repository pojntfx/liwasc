package components

import (
	"fmt"

	"github.com/maxence-charriere/go-app/v7/pkg/app"
)

type AppComponent struct {
	app.Compo
	DetailsOpen  bool
	SelectedNode int
}

func (c *AppComponent) Render() app.UI {
	nodes := []ListingNode{
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

	return app.Div().Body(
		&FilterComponent{Subnets: []string{"10.0.0.0/9", "192.168.0.0/27"}, Device: "eth0"},
		&DetailsComponent{
			Open:    c.DetailsOpen,
			OnClose: func(ctx app.Context, e app.Event) { c.handleDetailsClose() },
			Title: fmt.Sprintf("Node %v", func() string {
				if c.SelectedNode == -1 {
					return ""
				}

				return nodes[c.SelectedNode].MACAddress
			}()),
			Main: &ListingComponent{
				OnRowClick:   c.handleDetailsOpen,
				SelectedNode: c.SelectedNode,
				Nodes:        nodes,
			},
			Details: app.Div().Body(app.Text(fmt.Sprintf("%v", func() ListingNode {
				if c.SelectedNode == -1 {
					return ListingNode{}
				}

				return nodes[c.SelectedNode]
			}()))),
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
