package components

import (
	"github.com/maxence-charriere/go-app/v7/pkg/app"
)

type LoginActionsComponent struct {
	app.Compo

	Logout func()
}

func (c *LoginActionsComponent) Render() app.UI {
	return app.Button().Text("Logout").OnClick(func(ctx app.Context, e app.Event) {
		go c.Logout()
	})
}
