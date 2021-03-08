package experimental

import "github.com/maxence-charriere/go-app/v7/pkg/app"

type ActionsComponent struct {
	app.Compo

	TriggerNetworkScan func(nodeScanTimeout int64, portScanTimeout int64, macAddress string)
}

func (c *ActionsComponent) Render() app.UI {
	return app.Div().Body(
		app.Button().Text("Trigger network scan").OnClick(func(ctx app.Context, e app.Event) {
			c.TriggerNetworkScan(500, 50, "")
		}),
	)
}
