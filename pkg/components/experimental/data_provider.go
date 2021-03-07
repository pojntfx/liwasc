package experimental

import (
	"context"
	"sync"

	"github.com/maxence-charriere/go-app/v7/pkg/app"
	proto "github.com/pojntfx/liwasc-frontend-web/pkg/proto/generated"
	"google.golang.org/protobuf/types/known/emptypb"
)

type ScannerMetadata struct {
	Subnets []string
	Device  string
}

type Network struct {
	ScannerMetadata ScannerMetadata
}

type DataProviderChildrenProps struct {
	Network Network
}

type DataProviderComponent struct {
	app.Compo

	AuthenticatedContext context.Context
	MetadataService      proto.MetadataServiceClient
	Children             func(DataProviderChildrenProps) app.UI

	network     Network
	networkLock sync.Mutex
}

func (c *DataProviderComponent) Render() app.UI {
	return c.Children(DataProviderChildrenProps{
		Network: c.network,
	})
}

func (c *DataProviderComponent) OnMount(context app.Context) {
	// Initialize network struct
	c.networkLock.Lock()

	c.network = Network{
		ScannerMetadata: ScannerMetadata{
			Subnets: []string{},
			Device:  "",
		},
	}

	c.Update()
	c.networkLock.Unlock()

	// Get scanner metadata
	go func() {
		// Fetch from service
		scannerMetadata, err := c.MetadataService.GetMetadataForScanner(c.AuthenticatedContext, &emptypb.Empty{})
		if err != nil {
			panic(err)
		}

		// Write to struct
		c.networkLock.Lock()

		c.network.ScannerMetadata.Device = scannerMetadata.GetDevice()
		c.network.ScannerMetadata.Subnets = scannerMetadata.GetSubnets()

		c.Update()
		c.networkLock.Unlock()
	}()
}
