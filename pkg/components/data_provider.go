package components

import (
	"context"
	"log"
	"strings"

	"github.com/maxence-charriere/go-app/v7/pkg/app"
	"github.com/pojntfx/liwasc-frontend-web/pkg/models"
	proto "github.com/pojntfx/liwasc-frontend-web/pkg/proto/generated"
)

type DataProviderChildrenProps struct {
	Nodes []models.Node
}

type DataProviderComponent struct {
	app.Compo

	nodes []models.Node

	NetworkAndNodeScanServiceClient proto.NetworkAndNodeScanServiceClient
	NodeWakeServiceClient           proto.NodeWakeServiceClient
	Children                        func(DataProviderChildrenProps) app.UI
}

func (c *DataProviderComponent) Render() app.UI {
	return c.Children(DataProviderChildrenProps{c.nodes})
}

func (c *DataProviderComponent) OnMount(ctx app.Context) {
	go func() {
		c.nodes = []models.Node{
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
		}

		protoNetworkScanReferenceMessage := proto.NetworkScanReferenceMessage{
			NetworkScanID: -1,
		}

		stream, err := c.NetworkAndNodeScanServiceClient.SubscribeToNewNodes(context.Background(), &protoNetworkScanReferenceMessage)
		if err != nil {
			log.Fatal(err)
		}

		for {
			protoNode, err := stream.Recv()
			if err != nil {
				// All have been received
				if strings.Contains(err.Error(), "EOF") {
					return
				}

				log.Fatal(err)
			}

			log.Println(protoNode)
			// TODO: Convert protoNode to internal node, update, and subscribe to new services

			app.Dispatch(func() {
				c.Update()
			})
		}
	}()
}
