package services

import (
	"context"
	"log"
	"time"

	cmap "github.com/orcaman/concurrent-map"
	"github.com/pojntfx/liwasc/pkg/databases"
	proto "github.com/pojntfx/liwasc/pkg/proto/generated"
	"github.com/pojntfx/liwasc/pkg/scanners"
	nodeWakeModels "github.com/pojntfx/liwasc/pkg/sql/generated/node_wake"
	"github.com/pojntfx/liwasc/pkg/wakers"
	"github.com/ugjka/messenger"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type NodeWakeService struct {
	proto.UnimplementedNodeWakeServiceServer

	device             string
	nodeWakeDatabase   *databases.NodeWakeDatabase
	nodeWakeMessengers cmap.ConcurrentMap
	getIPAddress       func(string) (string, error)
	wakeOnLANWaker     *wakers.WakeOnLANWaker
}

func NewNodeWakeService(
	device string,
	nodeWakeDatabase *databases.NodeWakeDatabase,
	getIPAddress func(string) (string, error),
	wakeOnLANWaker *wakers.WakeOnLANWaker,
) *NodeWakeService {
	return &NodeWakeService{
		device:             device,
		nodeWakeDatabase:   nodeWakeDatabase,
		nodeWakeMessengers: cmap.New(),
		getIPAddress:       getIPAddress,
		wakeOnLANWaker:     wakeOnLANWaker,
	}
}

func (s *NodeWakeService) TriggerNodeWake(ctx context.Context, nodeWakeTriggerMessage *proto.NodeWakeTriggerMessage) (*proto.NodeWakeReferenceMessage, error) {
	nodeWake := &nodeWakeModels.NodeWake{
		Done: 0,
	}
	nodeWakeID, err := s.nodeWakeDatabase.CreateNodeWake(nodeWake)
	if err != nil {
		return nil, status.Errorf(codes.Unknown, "could not create node wake in DB: %v", err.Error())
	}

	if err := s.wakeOnLANWaker.Write(nodeWakeTriggerMessage.GetMACAddress()); err != nil {
		return nil, status.Errorf(codes.Unknown, "could not send Wake-on-LAN packet: %v", err.Error())
	}

	go func() {
		nodeWakeScanner := scanners.NewWakeScanner(
			nodeWakeTriggerMessage.GetMACAddress(),
			s.device,
			time.Millisecond*time.Duration(nodeWakeTriggerMessage.GetTimeout()),
			s.getIPAddress,
		)
		if err := nodeWakeScanner.Open(); err != nil {
			log.Println("could not open node wake scanner", err)

			return
		}

		nodeWakeMessenger := messenger.New(0, true)
		s.nodeWakeMessengers.Set(string(nodeWakeID), nodeWakeMessenger)

		// Send packet
		go func() {
			if err := nodeWakeScanner.Transmit(); err != nil {
				log.Println("could not send Wake-on-LAN packet", err)

				return
			}
		}()

		// Receive packet
		for {
			node := nodeWakeScanner.Read()

			// Wake scan is done
			if node == nil {
				break
			}

			dbNode := &nodeWakeModels.Node{
				MacAddress: node.MacAddress,
				PoweredOn: func() int64 {
					if node.Awake {
						return 1
					}

					return 0
				}(),
			}

			if _, err := s.nodeWakeDatabase.UpsertNode(dbNode, nodeWakeID); err != nil {
				log.Println("could not create node in DB", err)

				break
			}

			nodeWakeMessenger.Broadcast(dbNode)
		}

		nodeWakeMessenger.Reset()

		nodeWake.Done = 1
		if _, err := s.nodeWakeDatabase.UpdateNodeWakeScan(nodeWake); err != nil {
			log.Println("could not update node wake scan in DB", err)

			return
		}

		s.nodeWakeMessengers.Remove(string(nodeWakeID))
	}()

	protoNodeWakeReferenceMessage := &proto.NodeWakeReferenceMessage{
		MACAddress: nodeWakeTriggerMessage.GetMACAddress(),
		NodeWakeID: nodeWakeID,
	}

	return protoNodeWakeReferenceMessage, nil
}
