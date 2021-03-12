package components

import (
	"strconv"

	"github.com/maxence-charriere/go-app/v7/pkg/app"
)

type DataActionsComponent struct {
	app.Compo

	nodeScanTimeout    int64
	portScanTimeout    int64
	nodeScanMACAddress string

	nodeWakeTimeout    int64
	nodeWakeMACAddress string

	Nodes []Node

	TriggerNetworkScan func(nodeScanTimeout int64, portScanTimeout int64, macAddress string)
	StartNodeWake      func(nodeWakeTimeout int64, macAddress string)
}

const (
	// Names and IDs
	nodeScanTimeoutName    = "nodeScanTimeout"
	portScanTimeoutName    = "portScanTimeout"
	nodeScanMACAddressName = "nodeScanMACAddressTimeout"

	nodeWakeTimeoutName    = "nodeWakeTimeout"
	nodeWakeMACAddressName = "nodeWakeMACAddressTimeout"

	// Default values
	defaultNodeScanTimeout = 500
	defaultPortScanTimeout = 50
	allMACAddresses        = "ff:ff:ff:ff"

	defaultNodeWakeTimeout    = 1000
	defaultNodeWakeMACAddress = ""
)

func (c *DataActionsComponent) Render() app.UI {
	return app.Div().Body(
		app.Form().Body(
			// Node Scan Timeout Input
			app.
				Label().
				For(nodeScanTimeoutName).
				Text("Node Scan Timeout (in ms): "),
			&Controlled{
				Component: app.
					Input().
					Name(nodeScanTimeoutName).
					ID(nodeScanTimeoutName).
					Type("number").
					Required(true).
					Min(1).
					Step(1).
					Placeholder(strconv.Itoa(defaultNodeScanTimeout)).
					OnInput(func(ctx app.Context, e app.Event) {
						v, err := strconv.Atoi(ctx.JSSrc.Get("value").String())
						if err != nil || v == 0 {
							c.Update()

							return
						}

						c.nodeScanTimeout = int64(v)

						c.Update()
					}),
				Value: c.nodeScanTimeout,
			},
			app.Br(),
			// Port Scan Timeout Input
			app.
				Label().
				For(portScanTimeoutName).
				Text("Port Scan Timeout (in ms): "),
			&Controlled{
				Component: app.
					Input().
					Name(portScanTimeoutName).
					ID(portScanTimeoutName).
					Type("number").
					Required(true).
					Min(1).
					Step(1).
					Placeholder(strconv.Itoa(defaultPortScanTimeout)).
					OnInput(func(ctx app.Context, e app.Event) {
						v, err := strconv.Atoi(ctx.JSSrc.Get("value").String())
						if err != nil || v == 0 {
							c.Update()

							return
						}

						c.portScanTimeout = int64(v)

						c.Update()
					}),
				Value: c.portScanTimeout,
			},
			app.Br(),
			// Node Scan MAC Address Input
			app.
				Label().
				For(nodeScanMACAddressName).
				Text("Node Scan MAC Address: "),
			&Controlled{
				Component: app.
					Select().
					Name(nodeScanMACAddressName).
					ID(nodeScanMACAddressName).
					Required(true).
					OnInput(func(ctx app.Context, e app.Event) {
						c.nodeScanMACAddress = ctx.JSSrc.Get("value").String()

						c.Update()
					}).Body(
					append(
						[]app.UI{
							app.
								Option().
								Value(allMACAddresses).
								Text("All Addresses"),
						},
						app.Range(c.Nodes).Slice(func(i int) app.UI {
							return app.
								Option().
								Value(c.Nodes[i].MACAddress).
								Text(c.Nodes[i].MACAddress)
						}))...,
				),
				Value: c.nodeScanMACAddress,
			},
			app.Br(),
			// Network Scan Input Trigger
			app.
				Input().
				Type("submit").
				Value("Trigger network scan"),
		).OnSubmit(func(ctx app.Context, e app.Event) {
			e.PreventDefault()

			macAddress := c.nodeScanMACAddress
			if macAddress == allMACAddresses {
				macAddress = ""
			}

			go c.TriggerNetworkScan(c.nodeScanTimeout, c.portScanTimeout, macAddress)
		}),
		app.Form().Body(
			// Node Wake Timeout Input
			app.
				Label().
				For(nodeWakeTimeoutName).
				Text("Node Wake Timeout (in ms): "),
			&Controlled{
				Component: app.
					Input().
					Name(nodeWakeTimeoutName).
					ID(nodeWakeTimeoutName).
					Type("number").
					Required(true).
					Min(1).
					Step(1).
					Placeholder(strconv.Itoa(defaultNodeWakeTimeout)).
					OnInput(func(ctx app.Context, e app.Event) {
						v, err := strconv.Atoi(ctx.JSSrc.Get("value").String())
						if err != nil || v == 0 {
							c.Update()

							return
						}

						c.nodeWakeTimeout = int64(v)

						c.Update()
					}),
				Value: c.nodeWakeTimeout,
			},
			app.Br(),
			// Node Wake MAC Address Input
			app.
				Label().
				For(nodeWakeMACAddressName).
				Text("Node Wake MAC Address: "),
			&Controlled{
				Component: app.
					Select().
					Name(nodeWakeMACAddressName).
					ID(nodeWakeMACAddressName).
					Required(true).
					OnInput(func(ctx app.Context, e app.Event) {
						c.nodeWakeMACAddress = ctx.JSSrc.Get("value").String()

						c.Update()
					}).Body(
					app.Range(c.Nodes).Slice(func(i int) app.UI {
						return app.
							Option().
							Value(c.Nodes[i].MACAddress).
							Text(c.Nodes[i].MACAddress)
					}),
				),
				Value: c.nodeWakeMACAddress,
			},
			app.Br(),
			// Node Wake Input Trigger
			app.
				Input().
				Type("submit").
				Value("Trigger node wake"),
		).OnSubmit(func(ctx app.Context, e app.Event) {
			e.PreventDefault()

			go c.StartNodeWake(c.nodeWakeTimeout, c.nodeWakeMACAddress)
		}),
	)
}

func (c *DataActionsComponent) OnMount(context app.Context) {
	// Initialize node scan form
	c.nodeScanTimeout = defaultNodeScanTimeout
	c.portScanTimeout = defaultPortScanTimeout
	c.nodeScanMACAddress = allMACAddresses

	// Initialize node wake form
	c.nodeWakeTimeout = defaultNodeWakeTimeout
	c.nodeWakeMACAddress = defaultNodeWakeMACAddress
}
