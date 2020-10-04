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
	Nodes []*models.Node

	Connected bool
	Scanning  bool
}

type DataProviderComponent struct {
	app.Compo

	nodes []*models.Node

	NetworkAndNodeScanServiceClient proto.NetworkAndNodeScanServiceClient
	NodeWakeServiceClient           proto.NodeWakeServiceClient
	Children                        func(DataProviderChildrenProps) app.UI
}

func (c *DataProviderComponent) Render() app.UI {
	return c.Children(DataProviderChildrenProps{
		Nodes: c.nodes,

		Connected: true,
		Scanning:  true,
	})
}

func (c *DataProviderComponent) OnMount(ctx app.Context) {
	go func() {
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

			newNode := &models.Node{
				PoweredOn:    protoNode.LucidNode.PoweredOn,
				MACAddress:   protoNode.LucidNode.MACAddress,
				IPAddress:    protoNode.LucidNode.IPAddress,
				Vendor:       protoNode.LucidNode.Vendor,
				Registry:     protoNode.LucidNode.Registry,
				Organization: protoNode.LucidNode.Organization,
				Address:      protoNode.LucidNode.Address,
				Visible:      protoNode.LucidNode.Visible,
				Services:     []*models.Service{},
			}

			if protoNode.NodeScanID != -1 {
				go func() {
					protoNodeScanReferenceMessage := &proto.NodeScanReferenceMessage{
						NodeScanID: protoNode.NodeScanID,
					}

					stream, err := c.NetworkAndNodeScanServiceClient.SubscribeToNewOpenServices(context.Background(), protoNodeScanReferenceMessage)
					if err != nil {
						log.Fatal(err)
					}

					for {
						protoService, err := stream.Recv()
						if err != nil {
							// All have been received
							if strings.Contains(err.Error(), "EOF") {
								return
							}

							log.Fatal(err)
						}

						service := &models.Service{
							ServiceName:             protoService.ServiceName,
							PortNumber:              int(protoService.PortNumber),
							TransportProtocol:       protoService.TransportProtocol,
							Description:             protoService.Description,
							Assignee:                protoService.Assignee,
							Contact:                 protoService.Contact,
							RegistrationDate:        protoService.RegistrationDate,
							ModificationDate:        protoService.ModificationDate,
							Reference:               protoService.Reference,
							ServiceCode:             protoService.ServiceCode,
							UnauthorizedUseReported: protoService.UnauthorizedUseReported,
							AssignmentNotes:         protoService.UnauthorizedUseReported,
						}

						newNode.Services = append(newNode.Services, service)

						app.Dispatch(func() {
							c.Update()
						})
					}
				}()
			}

			exists := false
			for _, node := range c.nodes {
				if node.MACAddress == protoNode.LucidNode.MACAddress {
					exists = true

					break
				}
			}

			if exists {
				for _, node := range c.nodes {
					if node.MACAddress == protoNode.LucidNode.MACAddress {
						node = newNode

						break
					}
				}
			} else {
				c.nodes = append(c.nodes, newNode)
			}

			// TODO: Subscribe to node scans if nodeScanID != -1

			app.Dispatch(func() {
				c.Update()
			})
		}
	}()
}
