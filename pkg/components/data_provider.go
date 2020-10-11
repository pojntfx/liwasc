package components

import (
	"context"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/maxence-charriere/go-app/v7/pkg/app"
	"github.com/pojntfx/liwasc-frontend-web/pkg/models"
	proto "github.com/pojntfx/liwasc-frontend-web/pkg/proto/generated"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/emptypb"
)

type DataProviderChildrenProps struct {
	Nodes []*models.Node

	Connected bool
	Scanning  bool

	Subnets []string
	Device  string

	TriggerNetworkScan func(*proto.NetworkScanTriggerMessage)
}

type DataProviderComponent struct {
	app.Compo

	nodes        []*models.Node
	nodesLock    *sync.Mutex
	nodesCounter sync.WaitGroup

	IDToken string

	NetworkAndNodeScanServiceClient proto.NetworkAndNodeScanServiceClient
	NodeWakeServiceClient           proto.NodeWakeServiceClient
	MetadataServiceClient           proto.MetadataServiceClient
	Children                        func(DataProviderChildrenProps) app.UI

	connected bool
	scanning  bool

	subnets []string
	device  string
}

func (c *DataProviderComponent) Render() app.UI {
	return c.Children(DataProviderChildrenProps{
		Nodes: c.nodes,

		Connected: c.connected,
		Scanning:  c.scanning,

		Subnets: c.subnets,
		Device:  c.device,

		TriggerNetworkScan: c.triggerNetworkScan,
	})
}

func (c *DataProviderComponent) OnMount(ctx app.Context) {
	c.nodesLock = &sync.Mutex{}

	app.Dispatch(func() {
		c.connected = true
		c.scanning = false

		c.Update()
	})

	go func() {
		log.Println("getting metadata from service")

		metadata, err := c.MetadataServiceClient.GetMetadata(c.getAuthenticatedContext(), &emptypb.Empty{})
		if err != nil {
			log.Println("could not get metadata", err)

			c.invalidateConnection()

			return
		}

		log.Printf("received metadata: %v\n", metadata)

		c.subnets = metadata.GetSubnets()
		c.device = metadata.GetDevice()

		c.Update()
	}()

	go func() {
		log.Println("subscribing to periodic background network scan IDs")

		periodicBackgroundNetworkScanStream, err := c.NetworkAndNodeScanServiceClient.SubscribeToNewPeriodicNetworkScans(c.getAuthenticatedContext(), &emptypb.Empty{})
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

			if err := c.subscribeToNetworkScan(periodicNetworkScanReference); err != nil {
				log.Println("could not subscribe to network scan, retrying in 5s", err)

				c.invalidateConnection()

				time.Sleep(time.Second * 5)

				continue
			}
		}
	}()
}

func (c *DataProviderComponent) triggerNetworkScan(networkScanMessage *proto.NetworkScanTriggerMessage) {
	log.Println("triggering network scan")

	app.Dispatch(func() {
		c.scanning = true

		c.Update()
	})

	networkScanReferenceMessage, err := c.NetworkAndNodeScanServiceClient.TriggerNetworkScan(c.getAuthenticatedContext(), networkScanMessage)
	if err != nil {
		log.Println("could not trigger network scan", err)

		c.invalidateConnection()

		return
	}

	log.Printf("subscribing to network scan %v\n", networkScanReferenceMessage.GetNetworkScanID())

	if err := c.subscribeToNetworkScan(networkScanReferenceMessage); err != nil {
		log.Printf("could not subscribe to network scan %v: %v\n", networkScanReferenceMessage.GetNetworkScanID(), err)

		c.invalidateConnection()

		return
	}
}

func (c *DataProviderComponent) subscribeToNetworkScan(networkScanReference *proto.NetworkScanReferenceMessage) error {
	app.Dispatch(func() {
		c.nodes = []*models.Node{}

		c.Update()
	})

	nodeStream, err := c.NetworkAndNodeScanServiceClient.SubscribeToNewNodes(c.getAuthenticatedContext(), networkScanReference)
	if err != nil {
		return err
	}

	for {
		protoNode, err := nodeStream.Recv()
		if err != nil {
			if strings.Contains(err.Error(), "EOF") {
				log.Printf("network scan %v done, subscribing to next periodic background network scan\n", networkScanReference.GetNetworkScanID())

				go func(innerNodesCounter *sync.WaitGroup) {
					log.Printf("waiting for all node scans for network scan %v to finish (%v)\n", networkScanReference.GetNetworkScanID(), innerNodesCounter)

					innerNodesCounter.Wait()

					app.Dispatch(func() {
						c.scanning = false

						c.Update()
					})

					log.Printf("all node scans for network scan %v have finished (%v)\n", networkScanReference.GetNetworkScanID(), innerNodesCounter)
				}(&c.nodesCounter)

				break
			}

			log.Printf("could not receive node from network scan %v, retrying in 5s: %v\n", networkScanReference.GetNetworkScanID(), err)

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

		log.Printf("received node %v from network scan %v\n", node.MACAddress, networkScanReference.GetNetworkScanID())

		c.nodesLock.Lock()
		existingNodeIndex := -1
		for i, oldNode := range c.nodes {
			if oldNode.MACAddress == node.MACAddress {
				existingNodeIndex = i

				break
			}
		}

		if existingNodeIndex != -1 {
			c.nodes[existingNodeIndex] = node
		} else {
			c.nodes = append(c.nodes, node)
		}
		c.nodesLock.Unlock()

		c.Update()

		protoNodeScanReferenceMessage := &proto.NodeScanReferenceMessage{
			NodeScanID: protoNode.GetNodeScanID(),
		}

		if protoNode.GetNodeScanID() == -1 {
			log.Printf("node %v did not specify a node scan ID, skipping\n", node.MACAddress)

			continue
		}

		log.Printf("subscribing to services for node %v and node scan %v\n", node.MACAddress, protoNode.GetNodeScanID())

		c.nodesCounter.Add(1)

		serviceStream, err := c.NetworkAndNodeScanServiceClient.SubscribeToNewOpenServices(c.getAuthenticatedContext(), protoNodeScanReferenceMessage)
		if err != nil {
			log.Println("could not subscribe to node scan, retrying in 5s", err)

			c.invalidateConnection()

			time.Sleep(time.Second * 5)

			continue
		}

		go func(innerNodesCounter *sync.WaitGroup) {
			for {
				protoService, err := serviceStream.Recv()
				if err != nil {
					if strings.Contains(err.Error(), "EOF") {
						log.Printf("node scan %v done\n", protoNodeScanReferenceMessage.GetNodeScanID())

						innerNodesCounter.Done()

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

				c.nodesLock.Lock()

				node.Services = append(node.Services, service)

				c.nodesLock.Unlock()

				c.Update()
			}

		}(&c.nodesCounter)
	}

	return nil
}

func (c *DataProviderComponent) invalidateConnection() {
	app.Dispatch(func() {
		c.connected = false
		c.scanning = false

		c.Update()
	})
}

func (c *DataProviderComponent) getAuthenticatedContext() context.Context {
	return metadata.AppendToOutgoingContext(context.Background(), AUTHORIZATION_METADATA_KEY, c.IDToken)
}
