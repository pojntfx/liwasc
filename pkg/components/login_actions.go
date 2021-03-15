package components

import (
	"github.com/maxence-charriere/go-app/v7/pkg/app"
)

type LoginActionsComponent struct {
	app.Compo

	Logout func()
}

func (c *LoginActionsComponent) Render() app.UI {
	return app.Button().
		Text("Logout").
		Class("pf-c-button pf-m-primary pf-m-control").
		Type("button").
		OnClick(func(ctx app.Context, e app.Event) {
			go c.Logout()
		})
}
