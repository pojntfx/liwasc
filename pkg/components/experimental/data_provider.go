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

type Node struct {
	createdAt time.Time
	priority  int64

	MACAddress       string
	IPAddress        string
	PoweredOn        bool
	LastPortScanDate time.Time
}

type Network struct {
	ScannerMetadata  ScannerMetadata
	LastNodeScanDate time.Time
	Nodes            []Node
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
			Nodes:            []Node{},
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
				panic(err)
			}

			// Parse the scan's date
			nodeScanCreatedAt, err := time.Parse(time.RFC3339, nodeScan.GetCreatedAt())
			if err != nil {
				panic(err)
			}

			// Only continue evaluation if this scan is newer
			if nodeScanCreatedAt.After(c.network.LastNodeScanDate) {
				// Set the new latest node scan date
				c.dispatch(func() {
					c.network.LastNodeScanDate = nodeScanCreatedAt
				})

				// Subscribe to nodes
				go func(nsm *proto.NodeScanMessage) {
					// Get stream from service
					nodeStream, err := c.NodeAndPortScanService.SubscribeToNodes(c.AuthenticatedContext, nsm)
					if err != nil {
						panic(err)
					}

					// Process stream
					for {
						// Receive node from stream
						node, err := nodeStream.Recv()
						if err != nil {
							if err == io.EOF {
								break
							}

							panic(err)
						}

						// Parse the node's date
						nodeCreatedAt, err := time.Parse(time.RFC3339, node.GetCreatedAt())
						if err != nil {
							panic(err)
						}

						c.dispatch(func() {
							// Only continue if this node is newer and has a higher priority
							lastKnownNodeIndex := -1
							for i, knownNode := range c.network.Nodes {
								if knownNode.MACAddress == node.GetMACAddress() {
									if nodeCreatedAt.After(knownNode.createdAt) && knownNode.priority > node.GetPriority() {
										// Ignore the node
										return
									}

									lastKnownNodeIndex = i
								}
							}

							// If an old node exists, remove it
							if lastKnownNodeIndex != -1 {
								c.network.Nodes = append(c.network.Nodes[:lastKnownNodeIndex], c.network.Nodes[lastKnownNodeIndex+1:]...)
							}

							// Add the new node
							c.network.Nodes = append(c.network.Nodes, Node{
								createdAt: nodeCreatedAt,
								priority:  node.GetPriority(),

								MACAddress:       node.GetMACAddress(),
								IPAddress:        node.GetIPAddress(),
								PoweredOn:        node.GetPoweredOn(),
								LastPortScanDate: time.Unix(0, 0),
							})
						})

						// Subscribe to port scans
						go func(nm *proto.NodeMessage) {
							// Get stream from service
							portScanStream, err := c.NodeAndPortScanService.SubscribeToPortScans(c.AuthenticatedContext, nm)
							if err != nil {
								panic(err)
							}

							// Process stream
							for {
								// Receive scan from stream
								portScan, err := portScanStream.Recv()
								if err != nil {
									if err == io.EOF {
										break
									}

									panic(err)
								}

								// Parse the scan's date
								portScanCreatedAt, err := time.Parse(time.RFC3339, portScan.GetCreatedAt())
								if err != nil {
									panic(err)
								}

								// Check if this port scan is the newest one
								portScanIsNewest := false
								c.dispatch(func() {
									for i, currentNode := range c.network.Nodes {
										if currentNode.MACAddress == node.GetMACAddress() && portScanCreatedAt.After(currentNode.LastPortScanDate) {
											portScanIsNewest = true

											c.network.Nodes[i].LastPortScanDate = portScanCreatedAt

											break
										}
									}
								})

								// Only continue evaluation if this scan is newer
								if portScanIsNewest {
									// Subscribe to ports
									go func(ps *proto.PortScanMessage) {
										// Get stream from service
										portStream, err := c.NodeAndPortScanService.SubscribeToPorts(c.AuthenticatedContext, ps)
										if err != nil {
											panic(err)
										}

										// Process stream
										for {
											// Receive port from stream
											port, err := portStream.Recv()
											if err != nil {
												if err == io.EOF {
													break
												}

												panic(err)
											}

											log.Printf("received port %v for node %v\n", port, nm.GetMACAddress())
										}
									}(portScan)
								}
							}
						}(node)
					}
				}(nodeScan)
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
