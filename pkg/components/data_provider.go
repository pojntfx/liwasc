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

	TriggerNetworkScan func(*proto.NodeScanStartMessage)
}

type DataProviderComponent struct {
	app.Compo

	nodes        []*models.Node
	nodesLock    *sync.Mutex
	nodesCounter sync.WaitGroup

	IDToken string

	MetadataServiceClient proto.MetadataServiceClient
	Children              func(DataProviderChildrenProps) app.UI

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

		metadata, err := c.MetadataServiceClient.GetMetadataForScanner(c.getAuthenticatedContext(), &emptypb.Empty{})
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
}

func (c *DataProviderComponent) triggerNetworkScan(networkScanMessage *proto.NodeScanStartMessage) {
	log.Println("triggering network scan")

	app.Dispatch(func() {
		c.scanning = true

		c.Update()
	})
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
