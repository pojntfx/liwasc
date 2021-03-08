package experimental

import (
	"strconv"

	"github.com/maxence-charriere/go-app/v7/pkg/app"
	"github.com/pojntfx/liwasc-frontend-web/pkg/components/experimental/helpers"
)

type ActionsComponent struct {
	app.Compo

	nodeScanTimeout int64
	portScanTimeout int64
	macAddress      string

	TriggerNetworkScan func(nodeScanTimeout int64, portScanTimeout int64, macAddress string)
}

const (
	nodeScanTimeoutName = "nodeScanTimeout"
	portScanTimeoutName = "portScanTimeout"

	defaultNodeScanTimeout = 500
	defaultPortScanTimeout = 50
	defaultMACAddress      = ""
)

func (c *ActionsComponent) Render() app.UI {
	return app.Div().Body(
		app.Form().Body(
			// Node Scan Timeout Input
			app.
				Label().
				For(nodeScanTimeoutName).
				Text("Node Scan Timeout (in ms)"),
			&helpers.Controlled{
				Component: app.
					Input().
					Name(nodeScanTimeoutName).
					Type("number").
					Required(true).
					Min(1).
					Step(1).
					Placeholder("500").
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
				Text("Port Scan Timeout (in ms)"),
			&helpers.Controlled{
				Component: app.
					Input().
					Name(portScanTimeoutName).
					Type("number").
					Required(true).
					Min(1).
					Step(1).
					Placeholder("50").
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
			// Input Trigger
			app.
				Input().
				Type("submit").
				Value("Trigger network scan"),
		).OnSubmit(func(ctx app.Context, e app.Event) {
			e.PreventDefault()

			go c.TriggerNetworkScan(c.nodeScanTimeout, c.portScanTimeout, c.macAddress)
		}),
	)
}

func (c *ActionsComponent) OnMount(context app.Context) {
	// Initialize form state
	c.nodeScanTimeout = defaultNodeScanTimeout
	c.portScanTimeout = defaultPortScanTimeout
	c.macAddress = defaultMACAddress
}
