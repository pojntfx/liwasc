package components

import (
	"context"
	"log"
	"strings"
	"time"

	"github.com/maxence-charriere/go-app/v7/pkg/app"
	"github.com/pojntfx/liwasc-frontend-web/pkg/models"
	proto "github.com/pojntfx/liwasc-frontend-web/pkg/proto/generated"
	"google.golang.org/protobuf/types/known/emptypb"
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

	connected bool
	scanning  bool
}

func (c *DataProviderComponent) Render() app.UI {
	return c.Children(DataProviderChildrenProps{
		Nodes: c.nodes,

		Connected: c.connected,
		Scanning:  c.scanning,
	})
}

func (c *DataProviderComponent) OnMount(ctx app.Context) {
	app.Dispatch(func() {
		c.connected = true
		c.scanning = false

		c.Update()
	})

	go func() {
		log.Println("subscribing to periodic background network scan IDs")

		periodicBackgroundNetworkScanStream, err := c.NetworkAndNodeScanServiceClient.SubscribeToNewPeriodicNetworkScans(context.Background(), &emptypb.Empty{})
		if err != nil {
			log.Println("could not subscribe to periodic background network scan IDs", err)

			c.invalidateConnection()

			return
		}

		for {
			periodicNetworkScanReference, err := periodicBackgroundNetworkScanStream.Recv()
			if err != nil {
				log.Println("could not receive periodic background network scan ID", err)

				c.invalidateConnection()

				break
			}

			log.Printf("subscribing to periodic background network scan %v\n", periodicNetworkScanReference.GetNetworkScanID())

			nodeStream, err := c.NetworkAndNodeScanServiceClient.SubscribeToNewNodes(context.Background(), periodicNetworkScanReference)
			if err != nil {
				log.Println("could not subscribe to network scan, retrying in 5s", err)

				c.invalidateConnection()

				time.Sleep(time.Second * 5)

				continue
			}

			go func() {
				for {
					protoNode, err := nodeStream.Recv()
					if err != nil {
						if strings.Contains(err.Error(), "EOF") {
							log.Printf("network scan %v done, subscribing to next periodic background network scan\n", periodicNetworkScanReference.GetNetworkScanID())

							break
						}

						log.Printf("could not receive node from network scan %v, retrying in 5s: %v\n", periodicNetworkScanReference.GetNetworkScanID(), err)

						c.invalidateConnection()

						time.Sleep(time.Second * 5)

						continue
					}

					node := &models.Node{
						PoweredOn:    protoNode.LucidNode.GetPoweredOn(),
						Address:      protoNode.LucidNode.GetAddress(),
						IPAddress:    protoNode.LucidNode.GetIPAddress(),
						MACAddress:   protoNode.LucidNode.GetMACAddress(),
						Organization: protoNode.LucidNode.GetOrganization(),
						Registry:     protoNode.LucidNode.GetRegistry(),
						Services:     []*models.Service{},
						Vendor:       protoNode.LucidNode.GetVendor(),
						Visible:      protoNode.LucidNode.GetVisible(),
					}

					log.Printf("received node %v from network scan %v\n", node.MACAddress, periodicNetworkScanReference.GetNetworkScanID())

					existingIndex := -1
					for i, oldNode := range c.nodes {
						if oldNode.MACAddress == node.MACAddress {
							existingIndex = i

							break
						}
					}

					if existingIndex != -1 {
						c.nodes[existingIndex] = node
					} else {
						c.nodes = append(c.nodes, node)
					}

					c.Update()

					protoNodeScanReferenceMessage := &proto.NodeScanReferenceMessage{
						NodeScanID: protoNode.GetNodeScanID(),
					}

					if protoNode.GetNodeScanID() == -1 {
						log.Printf("node %v did not specify a node scan ID, skipping\n", node.MACAddress)

						continue
					}

					log.Printf("subscribing to services for node %v and node scan %v\n", node.MACAddress, protoNode.GetNodeScanID())

					serviceStream, err := c.NetworkAndNodeScanServiceClient.SubscribeToNewOpenServices(context.Background(), protoNodeScanReferenceMessage)
					if err != nil {
						log.Println("could not subscribe to node scan, retrying in 5s", err)

						c.invalidateConnection()

						time.Sleep(time.Second * 5)

						continue
					}

					go func() {
						for {
							protoService, err := serviceStream.Recv()
							if err != nil {
								if strings.Contains(err.Error(), "EOF") {
									log.Printf("node scan %v done\n", protoNodeScanReferenceMessage.GetNodeScanID())

									break
								}

								log.Printf("could not receive service for node scan %v, retrying in 5s: %v\n", protoNodeScanReferenceMessage.GetNodeScanID(), err)

								c.invalidateConnection()

								time.Sleep(time.Second * 5)

								continue
							}

							service := &models.Service{
								Assignee:                protoService.GetAssignee(),
								AssignmentNotes:         protoService.GetAssignmentNotes(),
								Contact:                 protoService.GetContact(),
								Description:             protoService.GetDescription(),
								ModificationDate:        protoService.GetModificationDate(),
								PortNumber:              int(protoService.GetPortNumber()),
								Reference:               protoService.GetReference(),
								RegistrationDate:        protoService.GetRegistrationDate(),
								ServiceCode:             protoService.GetServiceCode(),
								ServiceName:             protoService.GetServiceName(),
								TransportProtocol:       protoService.GetTransportProtocol(),
								UnauthorizedUseReported: protoService.GetUnauthorizedUseReported(),
							}

							log.Printf("received service %v/%v from node scan %v\n", service.PortNumber, service.TransportProtocol, protoNodeScanReferenceMessage.GetNodeScanID())

							existingIndex := -1
							for i, oldService := range node.Services {
								if oldService.PortNumber == service.PortNumber || oldService.TransportProtocol == service.TransportProtocol {
									existingIndex = i

									break
								}
							}

							if existingIndex != -1 {
								node.Services[existingIndex] = service
							} else {
								node.Services = append(node.Services, service)
							}

							c.Update()
						}
					}()
				}
			}()
		}
	}()
}

func (c *DataProviderComponent) invalidateConnection() {
	app.Dispatch(func() {
		c.connected = false
		c.scanning = false

		c.Update()
	})
}
