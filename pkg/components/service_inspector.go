package components

import (
	"fmt"

	"github.com/maxence-charriere/go-app/v7/pkg/app"
	"github.com/pojntfx/liwasc-frontend-web/pkg/models"
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
				Title:   "Assignee",
				Icon:    "fas fa-user-circle",
				Content: &SearchLinkComponent{Topic: c.Service.Assignee},
			},
			&DefinitionComponent{
				Title: "Contact",
				Icon:  "fas fa-paper-plane",
				Content: app.A().
					Href(fmt.Sprintf("mailto:%v", c.Service.Contact)).
					Text(c.Service.Contact),
			},
		),
		app.Dl().Class("pf-c-description-list pf-m-2-col pf-u-mb-md").Body(
			&DefinitionComponent{
				Title:   "Registration Date",
				Icon:    "fas fa-calendar-plus",
				Content: app.Text(c.Service.RegistrationDate),
			},
			&DefinitionComponent{
				Title:   "Modification Date",
				Icon:    "fas fa-pen-square",
				Content: app.Text(c.Service.ModificationDate),
			},
		),
		app.Dl().Class("pf-c-description-list pf-m-2-col pf-u-mb-md").Body(
			&DefinitionComponent{
				Title:   "Reference",
				Icon:    "fas fa-book-open",
				Content: &SearchLinkComponent{Topic: c.Service.Reference},
			},
			&DefinitionComponent{
				Title:   "Service Code",
				Icon:    "fas fa-barcode",
				Content: app.Text(c.Service.ServiceCode),
			},
		),
		app.Hr().Class("pf-c-divider pf-u-mb-md"),
		app.Dl().Class("pf-c-description-list").Body(
			&DefinitionComponent{
				Title:   "Unauthorized Use Reported",
				Icon:    "fas fa-exclamation-triangle",
				Content: app.Text(c.Service.UnauthorizedUseReported),
			},
			&DefinitionComponent{
				Title:   "Assignment Notes",
				Icon:    "fas fa-sticky-note",
				Content: app.Text(c.Service.AssignmentNotes),
			},
		),
	)
}
