package components

import "github.com/maxence-charriere/go-app/v7/pkg/app"

type AppComponent struct {
	app.Compo
}

func (c *AppComponent) Render() app.UI {
	return app.Div().Body(
		&FilterComponent{Subnets: []string{"10.0.0.0/9", "192.168.0.0/27"}, Device: "eth0"},
		&ListingComponent{
			Nodes: []ListingNode{
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
			},
		},
	)
}
