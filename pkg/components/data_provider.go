package components

import (
	"context"
	"fmt"
	"io"
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

	go func() {
		log.Println("subscribing to node scans")

		nodeScanStream, err := c.NodeAndPortScanServiceClient.SubscribeToNodeScans(c.getAuthenticatedContext(), &emptypb.Empty{})
		if err != nil {
			log.Println("could not subscribe to node scans", err)

			c.invalidateConnection()

			return
		}

		for {
			nodeScanMessage, err := nodeScanStream.Recv()
			if err != nil {
				log.Println("could not receive node scan message", err)

				c.invalidateConnection()

				break
			}

			log.Printf("received node scan message: %v\n", nodeScanMessage)

			go func(nsm *proto.NodeScanMessage) {
				nodeMessageStream, err := c.NodeAndPortScanServiceClient.SubscribeToNodes(c.getAuthenticatedContext(), nsm)
				if err != nil {
					log.Println("could not subscribe to nodes", err)

					c.invalidateConnection()

					return
				}

				for {
					nodeMessage, err := nodeMessageStream.Recv()
					if err != nil {
						if err == io.EOF {
							log.Println("no more nodes for this scan, continuing to the next scan")

							break
						}

						log.Println("could not receive node message", err)

						c.invalidateConnection()

						break
					}

					log.Printf("received node message: %v\n", nodeMessage)

					go func(nm *proto.NodeMessage) {
						portScanStream, err := c.NodeAndPortScanServiceClient.SubscribeToPortScans(c.getAuthenticatedContext(), nm)
						if err != nil {
							log.Println("could not subscribe to port scans", err)

							c.invalidateConnection()

							return
						}

						for {
							portScanMessage, err := portScanStream.Recv()
							if err != nil {
								if err == io.EOF {
									log.Println("no more port scans for this node, continuing to the next node")

									break
								}

								log.Println("could not receive node port scan message", err)

								c.invalidateConnection()

								break
							}

							log.Printf("received port scan message: %v\n", portScanMessage)

							go func(psm *proto.PortScanMessage) {
								portMessageStream, err := c.NodeAndPortScanServiceClient.SubscribeToPorts(c.getAuthenticatedContext(), psm)
								if err != nil {
									log.Println("could not subscribe to ports", err)

									c.invalidateConnection()

									return
								}

								for {
									portMessage, err := portMessageStream.Recv()
									if err != nil {
										if err == io.EOF {
											log.Println("no more ports for this scan, continuing to the next scan")

											break
										}

										log.Println("could not receive port message", err)

										c.invalidateConnection()

										break
									}

									log.Printf("received port message: %v\n", portMessage)
								}
							}(portScanMessage)
						}
					}(nodeMessage)
				}
			}(nodeScanMessage)
		}
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
