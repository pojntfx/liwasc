package components

import (
	"github.com/maxence-charriere/go-app/v8/pkg/app"
)

type ProgressButton struct {
	app.Compo

	Loading   bool
	Icon      string
	Text      string
	Secondary bool
	Classes   string

	OnClick func(ctx app.Context, e app.Event)
}

func (c *ProgressButton) Render() app.UI {
	return app.If(
		c.Text == "",
		app.Button().
			Class(func() string {
				classes := "pf-c-button pf-m-plain"

				if c.Loading {
					classes += " pf-m-progress pf-m-in-progress"
				}

				return classes
			}()).
			OnClick(func(ctx app.Context, e app.Event) {
				c.OnClick(ctx, e)
			}).
			Body(
				app.If(c.Loading,
					app.Span().
						Class("pf-c-button__progress").
						Body(
							app.Span().
								Class("pf-c-spinner pf-m-md").
								Aria("role", "progressbar").
								Aria("valuetext", "Loading...").
								Body(
									app.Span().Class("pf-c-spinner__clipper"),
									app.Span().Class("pf-c-spinner__lead-ball"),
									app.Span().Class("pf-c-spinner__tail-ball"),
								),
						)).Else(
					app.I().
						Class(c.Icon).
						Aria("hidden", true),
				),
			),
	).Else(
		app.Button().
			Class(func() string {
				classes := "pf-c-button pf-m-primary"

				if c.Secondary {
					classes = "pf-c-button pf-m-secondary"
				}

				if c.Loading {
					classes += " pf-m-progress pf-m-in-progress"
				}

				if c.Classes != "" {
					classes += " " + c.Classes
				}

				return classes
			}()).
			OnClick(func(ctx app.Context, e app.Event) {
				c.OnClick(ctx, e)
			}).
			Body(
				app.If(c.Loading,
					app.Span().
						Class("pf-c-button__progress").
						Body(
							app.Span().
								Class("pf-c-spinner pf-m-md").
								Aria("role", "progressbar").
								Aria("valuetext", "Loading...").
								Body(
									app.Span().Class("pf-c-spinner__clipper"),
									app.Span().Class("pf-c-spinner__lead-ball"),
									app.Span().Class("pf-c-spinner__tail-ball"),
								),
						)).Else(
					app.Span().
						Class("pf-c-button__icon pf-m-start").
						Body(
							app.I().
								Class(c.Icon).
								Aria("hidden", true),
						),
				),
				app.Text(c.Text),
			),
	)
}
