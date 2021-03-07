package experimental

import (
	"context"
	"io"
	"log"
	"sync"
	"time"

	"github.com/maxence-charriere/go-app/v7/pkg/app"
	proto "github.com/pojntfx/liwasc-frontend-web/pkg/proto/generated"
	"google.golang.org/protobuf/types/known/emptypb"
)

type ScannerMetadata struct {
	Subnets []string
	Device  string
}

type Network struct {
	ScannerMetadata  ScannerMetadata
	LastNodeScanDate time.Time
}

type DataProviderChildrenProps struct {
	Network Network
}

type DataProviderComponent struct {
	app.Compo

	AuthenticatedContext   context.Context
	MetadataService        proto.MetadataServiceClient
	NodeAndPortScanService proto.NodeAndPortScanServiceClient
	Children               func(DataProviderChildrenProps) app.UI

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
	c.dispatch(func() {
		c.network = Network{
			ScannerMetadata: ScannerMetadata{
				Subnets: []string{},
				Device:  "",
			},
			LastNodeScanDate: time.Unix(0, 0),
		}
	})

	// Get scanner metadata
	go func() {
		// Fetch from service
		scannerMetadata, err := c.MetadataService.GetMetadataForScanner(c.AuthenticatedContext, &emptypb.Empty{})
		if err != nil {
			panic(err)
		}

		// Write to struct
		c.dispatch(func() {
			c.network.ScannerMetadata.Device = scannerMetadata.GetDevice()
			c.network.ScannerMetadata.Subnets = scannerMetadata.GetSubnets()
		})
	}()

	// Subscribe to node scans
	go func() {
		// Get stream from service
		nodeScanStream, err := c.NodeAndPortScanService.SubscribeToNodeScans(c.AuthenticatedContext, &emptypb.Empty{})
		if err != nil {
			panic(err)
		}

		// Process stream
		for {
			// Receive scan from stream
			nodeScan, err := nodeScanStream.Recv()
			if err != nil {
				if err == io.EOF {
					break
				}

				panic(err)
			}

			// Only continue evaluation if this scan is newer than the newest one
			createdAt, err := time.Parse(time.RFC3339, nodeScan.GetCreatedAt())
			if err != nil {
				panic(err)
			}

			if createdAt.After(c.network.LastNodeScanDate) {
				// Set the new latest node scan date
				c.dispatch(func() {
					c.network.LastNodeScanDate = createdAt
				})

				log.Printf("continuing to evaluate node scan %v\n", c.network.LastNodeScanDate)
			}
		}
	}()
}

func (c *DataProviderComponent) dispatch(action func()) {
	c.networkLock.Lock()

	action()

	c.Update()
	c.networkLock.Unlock()
}
