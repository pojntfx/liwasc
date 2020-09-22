package main

import (
	"github.com/maxence-charriere/go-app/v7/pkg/app"
	"github.com/pojntfx/liwasc-frontend-web/pkg/components"
	"github.com/pojntfx/liwasc-frontend-web/pkg/models"
)

func main() {
	app.Route("/", &components.AppComponent{
		UserMenuOpen: false,
		UserAvatar:   "https://www.gravatar.com/avatar/db856df33fa4f4bce441819f604c90d5",
		UserName:     "Felicitas Pojtinger",

		Subnets:         []string{"10.0.0.0/9", "192.168.0.0/27"},
		Device:          "eth0",
		NodeSearchValue: "",

		Nodes: []models.Node{
			{
				PoweredOn:  false,
				MACAddress: "ff:ff:ff:ff",
				IPAddress:  "10.0.0.1",
				Vendor:     "TP-Link",
				Services: []models.Service{
					{
						ServiceName:             "ssh",
						PortNumber:              22,
						TransportProtocol:       "tcp",
						Description:             "Lorem dolor sit amet",
						Assignee:                "Felicitas Pojtinger",
						Contact:                 "felicitas@pojtinger.com",
						RegistrationDate:        "2002-01-01",
						ModificationDate:        "2002-02-02",
						Reference:               "RFC1234",
						ServiceCode:             "C241",
						UnauthorizedUseReported: "Aliens have not registered their protocols running on this port!",
						AssignmentNotes:         "Might glow in the dark.",
					},
					{
						ServiceName:             "dns",
						PortNumber:              53,
						TransportProtocol:       "udp",
						Description:             "Lorem dolor sit amet",
						Assignee:                "Felicitas Pojtinger",
						Contact:                 "felicitas@pojtinger.com",
						RegistrationDate:        "2002-01-01",
						ModificationDate:        "2002-02-02",
						Reference:               "RFC1234",
						ServiceCode:             "C241",
						UnauthorizedUseReported: "Aliens have not registered their protocols running on this port!",
						AssignmentNotes:         "Might glow in the dark.",
					},
					{
						ServiceName:             "http",
						PortNumber:              80,
						TransportProtocol:       "tcp",
						Description:             "Lorem dolor sit amet",
						Assignee:                "Felicitas Pojtinger",
						Contact:                 "felicitas@pojtinger.com",
						RegistrationDate:        "2002-01-01",
						ModificationDate:        "2002-02-02",
						Reference:               "RFC1234",
						ServiceCode:             "C241",
						UnauthorizedUseReported: "Aliens have not registered their protocols running on this port!",
						AssignmentNotes:         "Might glow in the dark.",
					},
				},
				Registry:     "MA-1",
				Organization: "TP-Link",
				Address:      "One Hacker Way",
				Visible:      true,
			},
			{
				PoweredOn:  true,
				MACAddress: "00:1B:44:11:3A:B7",
				IPAddress:  "10.0.0.2",
				Vendor:     "Realtek",
				Services: []models.Service{
					{
						ServiceName:             "echo",
						PortNumber:              7,
						TransportProtocol:       "tcp",
						Description:             "Lorem dolor sit amet",
						Assignee:                "Felicitas Pojtinger",
						Contact:                 "felicitas@pojtinger.com",
						RegistrationDate:        "2002-01-01",
						ModificationDate:        "2002-02-02",
						Reference:               "RFC1234",
						ServiceCode:             "C241",
						UnauthorizedUseReported: "Aliens have not registered their protocols running on this port!",
						AssignmentNotes:         "Might glow in the dark.",
					},
					{
						ServiceName:             "echo",
						PortNumber:              7,
						TransportProtocol:       "udp",
						Description:             "Lorem dolor sit amet",
						Assignee:                "Felicitas Pojtinger",
						Contact:                 "felicitas@pojtinger.com",
						RegistrationDate:        "2002-01-01",
						ModificationDate:        "2002-02-02",
						Reference:               "RFC1234",
						ServiceCode:             "C241",
						UnauthorizedUseReported: "Aliens have not registered their protocols running on this port!",
						AssignmentNotes:         "Might glow in the dark.",
					},
					{
						ServiceName:             "http",
						PortNumber:              80,
						TransportProtocol:       "tcp",
						Description:             "Lorem dolor sit amet",
						Assignee:                "Felicitas Pojtinger",
						Contact:                 "felicitas@pojtinger.com",
						RegistrationDate:        "2002-01-01",
						ModificationDate:        "2002-02-02",
						Reference:               "RFC1234",
						ServiceCode:             "C241",
						UnauthorizedUseReported: "Aliens have not registered their protocols running on this port!",
						AssignmentNotes:         "Might glow in the dark.",
					},
				},
				Registry:     "MA-1",
				Organization: "Realtek",
				Address:      "Two Hacker Way",
				Visible:      false,
			},
		},
		SelectedNode: -1,

		InspectorOpen:        false,
		ServicesAndPortsOpen: true,
		DetailsOpen:          false,

		ServicesOpen:         false,
		SelectedService:      -1,
		InspectorSearchValue: "",
	})

	app.Run()
}
