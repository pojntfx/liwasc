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

			for {
				protoNode, err := nodeStream.Recv()
				if err != nil {
					if strings.Contains(err.Error(), "EOF") {
						log.Printf("network scan %v done, subscribing to next periodic background network scan\n", periodicNetworkScanReference.GetNetworkScanID())

						break
					}

					log.Println("could not receive node, retrying in 5s", err)

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
					Services:     []*models.Service{}, // TODO: Subscribe to service and add here
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
			}
		}

		// protoNetworkScanReferenceMessage := proto.NetworkScanReferenceMessage{
		// 	NetworkScanID: -1,
		// }

		// stream, err := c.NetworkAndNodeScanServiceClient.SubscribeToNewNodes(context.Background(), &protoNetworkScanReferenceMessage)
		// if err != nil {
		// 	log.Fatal(err)
		// }

		// for {
		// 	protoNode, err := stream.Recv()
		// 	if err != nil {
		// 		// All have been received
		// 		if strings.Contains(err.Error(), "EOF") {
		// 			return
		// 		}

		// 		log.Fatal(err)
		// 	}

		// 	newNode := &models.Node{
		// 		PoweredOn:    protoNode.LucidNode.PoweredOn,
		// 		MACAddress:   protoNode.LucidNode.MACAddress,
		// 		IPAddress:    protoNode.LucidNode.IPAddress,
		// 		Vendor:       protoNode.LucidNode.Vendor,
		// 		Registry:     protoNode.LucidNode.Registry,
		// 		Organization: protoNode.LucidNode.Organization,
		// 		Address:      protoNode.LucidNode.Address,
		// 		Visible:      protoNode.LucidNode.Visible,
		// 		Services:     []*models.Service{},
		// 	}

		// 	if protoNode.NodeScanID != -1 {
		// 		go func() {
		// 			protoNodeScanReferenceMessage := &proto.NodeScanReferenceMessage{
		// 				NodeScanID: protoNode.NodeScanID,
		// 			}

		// 			stream, err := c.NetworkAndNodeScanServiceClient.SubscribeToNewOpenServices(context.Background(), protoNodeScanReferenceMessage)
		// 			if err != nil {
		// 				log.Fatal(err)
		// 			}

		// 			for {
		// 				protoService, err := stream.Recv()
		// 				if err != nil {
		// 					// All have been received
		// 					if strings.Contains(err.Error(), "EOF") {
		// 						return
		// 					}

		// 					log.Fatal(err)
		// 				}

		// 				service := &models.Service{
		// 					ServiceName:             protoService.ServiceName,
		// 					PortNumber:              int(protoService.PortNumber),
		// 					TransportProtocol:       protoService.TransportProtocol,
		// 					Description:             protoService.Description,
		// 					Assignee:                protoService.Assignee,
		// 					Contact:                 protoService.Contact,
		// 					RegistrationDate:        protoService.RegistrationDate,
		// 					ModificationDate:        protoService.ModificationDate,
		// 					Reference:               protoService.Reference,
		// 					ServiceCode:             protoService.ServiceCode,
		// 					UnauthorizedUseReported: protoService.UnauthorizedUseReported,
		// 					AssignmentNotes:         protoService.UnauthorizedUseReported,
		// 				}

		// 				newNode.Services = append(newNode.Services, service)

		// 				app.Dispatch(func() {
		// 					c.Update()
		// 				})
		// 			}
		// 		}()
		// 	}

		// 	exists := false
		// 	for _, node := range c.nodes {
		// 		if node.MACAddress == protoNode.LucidNode.MACAddress {
		// 			exists = true

		// 			break
		// 		}
		// 	}

		// 	if exists {
		// 		for _, node := range c.nodes {
		// 			if node.MACAddress == protoNode.LucidNode.MACAddress {
		// 				node = newNode

		// 				break
		// 			}
		// 		}
		// 	} else {
		// 		c.nodes = append(c.nodes, newNode)
		// 	}

		// 	// TODO: Subscribe to node scans if nodeScanID != -1

		// 	app.Dispatch(func() {
		// 		c.Update()
		// 	})
		// }
	}()
}

func (c *DataProviderComponent) invalidateConnection() {
	app.Dispatch(func() {
		c.connected = false
		c.scanning = false

		c.Update()
	})
}
