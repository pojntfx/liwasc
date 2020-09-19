package components

import (
	"fmt"

	"github.com/maxence-charriere/go-app/v7/pkg/app"
)

type ExpandableSectionComponent struct {
	app.Compo
	Open     bool
	OnToggle func(ctx app.Context, e app.Event)
	Title    string
	Content  app.UI

	ref app.HTMLDiv
}

func (c *ExpandableSectionComponent) Render() app.UI {
	c.Sync()

	return app.Div().Class(fmt.Sprintf("pf-c-expandable-section pf-u-mb-md %v", func() string {
		if c.Open {
			return "pf-m-expanded"
		}

		return ""
	}())).Body(
		app.Button().Class("pf-c-expandable-section__toggle").Body(
			app.Span().Class("pf-c-expandable-section__toggle-icon").Body(
				app.I().Class("fas fa-angle-right"),
			),
			app.Span().Class("pf-c-expandable-section__toggle-text").Body(
				app.Text(c.Title),
			),
		).OnClick(c.OnToggle),
		c.ref,
	)
}

func (c *ExpandableSectionComponent) OnMount(ctx app.Context) {
	c.Sync()
}

func (c *ExpandableSectionComponent) Sync() {
	if c.ref == nil {
		c.ref = app.Div().Class("pf-c-expandable-section__content").Hidden(!c.Open).Body(
			c.Content,
		)
	} else {
		c.ref.JSValue().Set("hidden", !c.Open)
	}
}
