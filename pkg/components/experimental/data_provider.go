package experimental

import (
	"context"
	"io"
	"sync"
	"time"

	"github.com/maxence-charriere/go-app/v7/pkg/app"
	proto "github.com/pojntfx/liwasc-frontend-web/pkg/proto/generated"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type ScannerMetadata struct {
	Subnets []string
	Device  string
}

type Port struct {
	// Internal metadata
	createdAt time.Time
	priority  int64

	// Data
	PortNumber        int64
	TransportProtocol string

	// Public metadata
	ServiceName             string
	Description             string
	Assignee                string
	Contact                 string
	RegistrationDate        string
	ModificationDate        string
	Reference               string
	ServiceCode             string
	UnauthorizedUseReported string
	AssignmentNotes         string
}

type Node struct {
	// Internal metadata
	createdAt time.Time
	priority  int64

	// Data
	MACAddress       string
	IPAddress        string
	PoweredOn        bool
	LastPortScanDate time.Time
	Ports            []Port

	// Public metadata
	Vendor       string
	Registry     string
	Organization string
	Address      string
	Visible      bool
}

type Network struct {
	ScannerMetadata  ScannerMetadata
	NodeScanRunning  bool
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
			NodeScanRunning:  false,
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
			if nodeScanCreatedAt.After(c.network.LastNodeScanDate) || nodeScanCreatedAt.Equal(c.network.LastNodeScanDate) {
				c.dispatch(func() {
					// Set the new latest node scan date
					c.network.LastNodeScanDate = nodeScanCreatedAt

					// Update the node scan indicator
					if nodeScan.Done {
						c.network.NodeScanRunning = false
					} else {
						c.network.NodeScanRunning = true
					}
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

						// Get the node's metadata
						nodeMetadata, err := c.MetadataService.GetMetadataForNode(c.AuthenticatedContext, &proto.NodeMetadataReferenceMessage{
							MACAddress: node.GetMACAddress(),
						})
						if err != nil {
							if status, ok := status.FromError(err); ok && status.Code() == codes.NotFound {
								nodeMetadata = &proto.NodeMetadataMessage{
									MACAddress:   node.GetMACAddress(),
									Vendor:       "",
									Registry:     "",
									Organization: "",
									Address:      "",
									Visible:      true, // The majority are visible, so set it as the default value
								}
							} else {
								panic(err)
							}
						}

						c.dispatch(func() {
							// Only continue if this node is newer and has a higher priority
							lastKnownNodeIndex := -1
							for i, knownNode := range c.network.Nodes {
								if knownNode.MACAddress == node.GetMACAddress() {
									if (nodeCreatedAt.After(knownNode.createdAt) || nodeCreatedAt.Equal(knownNode.createdAt)) && knownNode.priority > node.GetPriority() {
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

								Vendor:       nodeMetadata.GetVendor(),
								Registry:     nodeMetadata.GetRegistry(),
								Organization: nodeMetadata.GetOrganization(),
								Address:      nodeMetadata.GetAddress(),
								Visible:      nodeMetadata.GetVisible(),
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
										if currentNode.MACAddress == node.GetMACAddress() && (portScanCreatedAt.After(currentNode.LastPortScanDate) || portScanCreatedAt.Equal(currentNode.LastPortScanDate)) {
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

											// Get the port's metadata
											portMetadata, err := c.MetadataService.GetMetadataForPort(c.AuthenticatedContext, &proto.PortMetadataReferenceMessage{
												PortNumber:        port.GetPortNumber(),
												TransportProtocol: port.GetTransportProtocol(),
											})
											if err != nil {
												if status, ok := status.FromError(err); ok && status.Code() == codes.NotFound {
													portMetadata = &proto.PortMetadataMessage{
														ServiceName:             "",
														PortNumber:              port.GetPortNumber(),
														TransportProtocol:       port.GetTransportProtocol(),
														Description:             "",
														Assignee:                "",
														Contact:                 "",
														RegistrationDate:        "",
														ModificationDate:        "",
														Reference:               "",
														ServiceCode:             "",
														UnauthorizedUseReported: "",
														AssignmentNotes:         "",
													}
												} else {
													panic(err)
												}
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
																if (portCreatedAt.After(knownPort.createdAt) || portCreatedAt.Equal(knownPort.createdAt)) && knownPort.priority > port.GetPriority() {
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

														ServiceName:             portMetadata.ServiceName,
														Description:             portMetadata.Description,
														Assignee:                portMetadata.Assignee,
														Contact:                 portMetadata.Contact,
														RegistrationDate:        portMetadata.RegistrationDate,
														ModificationDate:        portMetadata.ModificationDate,
														Reference:               portMetadata.Reference,
														ServiceCode:             portMetadata.ServiceCode,
														UnauthorizedUseReported: portMetadata.UnauthorizedUseReported,
														AssignmentNotes:         portMetadata.AssignmentNotes,
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
