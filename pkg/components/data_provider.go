package components

import (
	"context"
	"fmt"
	"log"
	"sync"

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

	TriggerNodeScan func(*proto.NodeScanStartMessage)
}

type DataProviderComponent struct {
	app.Compo

	nodes        []*models.Node
	nodesLock    *sync.Mutex
	nodesCounter sync.WaitGroup

	IDToken string

	MetadataServiceClient        proto.MetadataServiceClient
	NodeAndPortScanServiceClient proto.NodeAndPortScanServiceClient
	Children                     func(DataProviderChildrenProps) app.UI

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

		TriggerNodeScan: c.triggerNodeScan,
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
		log.Println("getting metadata")

		metadata, err := c.MetadataServiceClient.GetMetadataForScanner(c.getAuthenticatedContext(), &emptypb.Empty{})
		if err != nil {
			log.Println("could not get metadata", err)

			c.invalidateConnection()

			return
		}

		log.Printf("received metadata: %v\n", metadata)

		app.Dispatch(func() {
			c.subnets = metadata.GetSubnets()
			c.device = metadata.GetDevice()

			c.Update()
		})
	}()
}

func (c *DataProviderComponent) triggerNodeScan(nodeScanMessage *proto.NodeScanStartMessage) {
	app.Dispatch(func() {
		c.scanning = true

		c.Update()
	})

	go func() {
		log.Println("starting node scan")

		nodeScanMessage, err := c.NodeAndPortScanServiceClient.StartNodeScan(c.getAuthenticatedContext(), nodeScanMessage)
		if err != nil {
			log.Println("could not start node scan", err)

			c.invalidateConnection()

			return
		}

		log.Printf("received node scan message: %v\n", nodeScanMessage)
	}()
}

func (c *DataProviderComponent) invalidateConnection() {
	app.Dispatch(func() {
		c.connected = false
		c.scanning = false

		c.Update()
	})
}

func (c *DataProviderComponent) getAuthenticatedContext() context.Context {
	fmt.Println(c.IDToken)

	return metadata.AppendToOutgoingContext(context.Background(), AUTHORIZATION_METADATA_KEY, c.IDToken)
}
