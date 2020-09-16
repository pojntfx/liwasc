package components

import "github.com/maxence-charriere/go-app/v7/pkg/app"

type AppComponent struct {
	app.Compo
}

func (c *AppComponent) Render() app.UI {
	return app.Div().Body(
		&FilterComponent{Subnets: []string{"10.0.0.0/9", "192.168.0.0/27"}},
	)
}
