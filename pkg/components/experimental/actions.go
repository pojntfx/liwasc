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
)

func (c *ActionsComponent) Render() app.UI {
	return app.Div().Body(
		app.Form().Body(
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
						if err != nil {
							panic(err)
						}

						c.nodeScanTimeout = int64(v)
					}),
				Value: c.nodeScanTimeout,
			},
			app.
				Input().
				Type("submit").
				Value("Trigger network scan"),
		).OnSubmit(func(ctx app.Context, e app.Event) {
			e.PreventDefault()

			go c.TriggerNetworkScan(c.nodeScanTimeout, c.portScanTimeout, c.macAddress)

			c.clearNetworkScanForm()
		}),
	)
}

func (c *ActionsComponent) clearNetworkScanForm() {
	// Reset the values
	c.nodeScanTimeout = 0
	c.portScanTimeout = 0
	c.macAddress = ""

	c.Update()
}

func (c *ActionsComponent) OnMount(context app.Context) {
	// Initialize form state
	c.nodeScanTimeout = 0
	c.portScanTimeout = 0
	c.macAddress = ""
}
