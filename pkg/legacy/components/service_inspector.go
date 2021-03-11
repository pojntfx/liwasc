package components

import (
	"fmt"

	"github.com/maxence-charriere/go-app/v7/pkg/app"
	"github.com/pojntfx/liwasc-frontend/pkg/legacy/models"
)

type ServiceInspectorComponent struct {
	app.Compo

	Service *models.Service
}

func (c *ServiceInspectorComponent) Render() app.UI {
	return app.Div().Body(
		app.Div().Class("pf-u-mb-lg").Text(c.Service.Description),
		app.Dl().Class("pf-c-description-list pf-m-2-col pf-u-mb-md").Body(
			&DefinitionComponent{
				Title: "Assignee",
				Icon:  "fas fa-user-circle",
				Content: func() app.UI {
					if c.Service.Assignee != "" {
						return &SearchLinkComponent{Topic: c.Service.Assignee}
					}

					return app.Text("Unknown assignee")
				}(),
			},
			&DefinitionComponent{
				Title: "Contact",
				Icon:  "fas fa-paper-plane",
				Content: func() app.UI {
					if c.Service.Contact != "" {
						return app.A().
							Href(fmt.Sprintf("mailto:%v", c.Service.Contact)).
							Text(c.Service.Contact)
					}

					return app.Text("Unknown contact")
				}(),
			},
		),
		app.Dl().Class("pf-c-description-list pf-m-2-col pf-u-mb-md").Body(
			&DefinitionComponent{
				Title: "Registration Date",
				Icon:  "fas fa-calendar-plus",
				Content: app.Text(func() string {
					if c.Service.RegistrationDate != "" {
						return c.Service.RegistrationDate
					}

					return "Unknown registration date"
				}()),
			},
			&DefinitionComponent{
				Title: "Modification Date",
				Icon:  "fas fa-pen-square",
				Content: app.Text(func() string {
					if c.Service.ModificationDate != "" {
						return c.Service.ModificationDate
					}

					return "Unknown modification date"
				}()),
			},
		),
		app.Dl().Class("pf-c-description-list pf-m-2-col pf-u-mb-md").Body(
			&DefinitionComponent{
				Title: "Reference",
				Icon:  "fas fa-book-open",
				Content: func() app.UI {
					if c.Service.Reference != "" {
						return &SearchLinkComponent{Topic: c.Service.Reference}
					}

					return app.Text("Unknown reference")
				}(),
			},
			&DefinitionComponent{
				Title: "Service Code",
				Icon:  "fas fa-barcode",
				Content: app.Text(func() string {
					if c.Service.ServiceCode != "" {
						return c.Service.ServiceCode
					}

					return "Unknown service code"
				}()),
			},
		),
		app.Hr().Class("pf-c-divider pf-u-mb-md"),
		app.Dl().Class("pf-c-description-list").Body(
			&DefinitionComponent{
				Title: "Unauthorized Use Reported",
				Icon:  "fas fa-exclamation-triangle",
				Content: app.Text(func() string {
					if c.Service.UnauthorizedUseReported != "" {
						return c.Service.UnauthorizedUseReported
					}

					return "Not known"
				}()),
			},
			&DefinitionComponent{
				Title: "Assignment Notes",
				Icon:  "fas fa-sticky-note",
				Content: app.Text(func() string {
					if c.Service.AssignmentNotes != "" {
						return c.Service.AssignmentNotes
					}

					return "No assignment notes"
				}()),
			},
		),
	)
}
