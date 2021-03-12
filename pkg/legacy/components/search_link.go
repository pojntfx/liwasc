package components

import (
	"fmt"
	"net/url"

	"github.com/maxence-charriere/go-app/v7/pkg/app"
)

type SearchLinkComponent struct {
	app.Compo

	Topic string
}

func (c *SearchLinkComponent) Render() app.UI {
	return app.A().
		Href(fmt.Sprintf("https://duckduckgo.com/?q=%v", url.QueryEscape(c.Topic))).
		Target("_blank").
		Body(
			app.I().Class("fas fa-external-link-alt pf-u-mr-xs"),
			app.Text(c.Topic),
		)
}
