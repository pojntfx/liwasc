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
		app.Form().
			Class("pf-c-form").
			Body(
				// Node Scan Timeout Input
				&FormGroupComponent{
					Label: app.
						Label().
						For(nodeScanTimeoutName).
						Class("pf-c-form__label").
						Body(
							app.
								Span().
								Class("pf-c-form__label-text").
								Text("Node Scan Timeout (in ms)"),
						),
					Input: &Controlled{
						Component: app.
							Input().
							Name(nodeScanTimeoutName).
							ID(nodeScanTimeoutName).
							Type("number").
							Required(true).
							Min(1).
							Step(1).
							Placeholder(strconv.Itoa(defaultNodeScanTimeout)).
							Class("pf-c-form-control").
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
					Required: true,
				},
				// Port Scan Timeout Input
				&FormGroupComponent{
					Label: app.
						Label().
						For(portScanTimeoutName).
						Class("pf-c-form__label").
						Body(
							app.
								Span().
								Class("pf-c-form__label-text").
								Text("Port Scan Timeout (in ms)"),
						),
					Input: &Controlled{
						Component: app.
							Input().
							Name(portScanTimeoutName).
							ID(portScanTimeoutName).
							Type("number").
							Required(true).
							Min(1).
							Step(1).
							Placeholder(strconv.Itoa(defaultPortScanTimeout)).
							Class("pf-c-form-control").
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
					Required: true,
				},
				// Node Scan MAC Address Input
				&FormGroupComponent{
					Label: app.
						Label().
						For(nodeScanMACAddressName).
						Class("pf-c-form__label").
						Body(
							app.
								Span().
								Class("pf-c-form__label-text").
								Text("Node Scan MAC Address"),
						),
					Input: &Controlled{
						Component: app.
							Select().
							Name(nodeScanMACAddressName).
							ID(nodeScanMACAddressName).
							Required(true).
							Class("pf-c-form-control").
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
					Required: true,
				},
				// Network Scan Input Trigger
				app.Div().
					Class("pf-c-form__group pf-m-action").
					Body(
						app.Div().
							Class("pf-c-form__actions").
							Body(
								app.
									Button().
									Type("submit").
									Class("pf-c-button pf-m-primary").
									Text("Trigger network scan"),
							),
					),
			).OnSubmit(func(ctx app.Context, e app.Event) {
			e.PreventDefault()

			macAddress := c.nodeScanMACAddress
			if macAddress == allMACAddresses {
				macAddress = ""
			}

			go c.TriggerNetworkScan(c.nodeScanTimeout, c.portScanTimeout, macAddress)
		}),
		app.Form().
			Class("pf-c-form").
			Body(
				// Node Wake Timeout Input
				&FormGroupComponent{
					Label: app.
						Label().
						For(nodeWakeTimeoutName).
						Class("pf-c-form__label").
						Body(
							app.
								Span().
								Class("pf-c-form__label-text").
								Text("Node Wake Timeout (in ms)"),
						),
					Input: &Controlled{
						Component: app.
							Input().
							Name(nodeWakeTimeoutName).
							ID(nodeWakeTimeoutName).
							Type("number").
							Required(true).
							Min(1).
							Step(1).
							Placeholder(strconv.Itoa(defaultNodeWakeTimeout)).
							Class("pf-c-form-control").
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
					Required: true,
				},
				// Node Wake MAC Address Input
				&FormGroupComponent{
					Label: app.
						Label().
						For(nodeWakeMACAddressName).
						Class("pf-c-form__label").
						Body(
							app.
								Span().
								Class("pf-c-form__label-text").
								Text("Node Wake MAC Address"),
						),
					Input: &Controlled{
						Component: app.
							Select().
							Name(nodeWakeMACAddressName).
							ID(nodeWakeMACAddressName).
							Required(true).
							Class("pf-c-form-control").
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
					Required: true,
				},
				// Node Wake Input Trigger
				app.Div().
					Class("pf-c-form__group pf-m-action").
					Body(
						app.Div().
							Class("pf-c-form__actions").
							Body(
								app.
									Button().
									Type("submit").
									Class("pf-c-button pf-m-primary").
									Text("Trigger node wake"),
							),
					),
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
