package experimental

import (
	"context"
	"io"
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

type Port struct {
	createdAt time.Time
	priority  int64

	PortNumber        int64
	TransportProtocol string
}

type Node struct {
	createdAt time.Time
	priority  int64

	MACAddress       string
	IPAddress        string
	PoweredOn        bool
	LastPortScanDate time.Time
	Ports            []Port
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

									break
								}
							}

							// If an old node exists, remove it, but keep the ports
							ports := []Port{}
							if lastKnownNodeIndex != -1 {
								ports = c.network.Nodes[lastKnownNodeIndex].Ports
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
								Ports:            ports,
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

											// Parse the port's date
											portCreatedAt, err := time.Parse(time.RFC3339, port.GetCreatedAt())
											if err != nil {
												panic(err)
											}

											c.dispatch(func() {
												// Only continue if this port is newer and has a higher priority
												lastKnownNodeIndex := -1
												lastKnownPortIndex := -1
												for nodeIndex, knownNode := range c.network.Nodes {
													if knownNode.MACAddress == nm.GetMACAddress() {
														lastKnownNodeIndex = nodeIndex

														for portIndex, knownPort := range knownNode.Ports {
															if knownPort.PortNumber == port.PortNumber && knownPort.TransportProtocol == port.TransportProtocol {
																if portCreatedAt.After(knownPort.createdAt) && knownPort.priority > port.GetPriority() {
																	// Ignore the port
																	return
																}

																lastKnownPortIndex = portIndex

																break
															}
														}

														break
													}
												}

												// Asynchronity; the node specified in the outer loops might not exist anymore
												if lastKnownNodeIndex != -1 {
													// If an old port exists, remove it.
													if lastKnownPortIndex != -1 {
														c.network.Nodes[lastKnownNodeIndex].Ports = append(c.network.Nodes[lastKnownNodeIndex].Ports[:lastKnownPortIndex], c.network.Nodes[lastKnownNodeIndex].Ports[lastKnownPortIndex+1:]...)
													}

													// Add the new port
													c.network.Nodes[lastKnownNodeIndex].Ports = append(c.network.Nodes[lastKnownNodeIndex].Ports, Port{
														createdAt: portCreatedAt,
														priority:  port.GetPriority(),

														PortNumber:        port.GetPortNumber(),
														TransportProtocol: port.GetTransportProtocol(),
													})
												}
											})
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
