package services

import (
	"context"
	"log"
	"time"

	"github.com/pojntfx/liwasc/pkg/databases"
	proto "github.com/pojntfx/liwasc/pkg/proto/generated"
	"github.com/pojntfx/liwasc/pkg/scanners"
	models "github.com/pojntfx/liwasc/pkg/sql/generated/node_wake_neo"
	"github.com/pojntfx/liwasc/pkg/wakers"
	"github.com/ugjka/messenger"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type NodeWakeNeoService struct {
	proto.UnimplementedNodeWakeNeoServiceServer

	device         string
	wakeOnLANWaker *wakers.WakeOnLANWaker

	nodeWakeDatabase *databases.NodeWakeNeoDatabase
	getIPAddress     func(macAddress string) (ipAddress string, err error)

	nodeWakeMessenger *messenger.Messenger
}

func NewNodeWakeNeoService(
	device string,
	wakeOnLANWaker *wakers.WakeOnLANWaker,

	nodeWakeDatabase *databases.NodeWakeNeoDatabase,
	getIPAddress func(macAddress string) (ipAddress string, err error),
) *NodeWakeNeoService {
	return &NodeWakeNeoService{
		device:         device,
		wakeOnLANWaker: wakeOnLANWaker,

		nodeWakeDatabase: nodeWakeDatabase,
		getIPAddress:     getIPAddress,

		nodeWakeMessenger: messenger.New(0, true),
	}
}

func (s *NodeWakeNeoService) StartNodeWake(_ context.Context, nodeWakeStartMessage *proto.NodeWakeStartNeoMessage) (*proto.NodeWakeNeoMessage, error) {
	// Create and broadcast node wake in DB
	dbNodeWake := &models.NodeWakesNeo{
		Done:       0,
		PoweredOn:  0,
		MacAddress: nodeWakeStartMessage.GetMACAddress(),
	}
	if err := s.nodeWakeDatabase.CreateNodeWake(dbNodeWake); err != nil {
		log.Printf("could not create node wake in DB: %v\n", err)

		return nil, status.Errorf(codes.Unknown, "could not create node wake in DB")
	}
	s.nodeWakeMessenger.Broadcast(dbNodeWake)

	// Wake the node
	go func() {
		if err := s.wakeOnLANWaker.Write(dbNodeWake.MacAddress); err != nil {
			log.Printf("could not wake node: %v\n", err)

			return
		}
	}()

	// Transmit and receive node wakes
	go func() {
		for i := 0; i < 5; i++ {
			timeout := time.Millisecond * time.Duration(nodeWakeStartMessage.GetNodeWakeTimeout()/5)

			// Create and open wake scanner
			wakeScanner := scanners.NewWakeScanner(
				dbNodeWake.MacAddress,
				s.device,
				timeout,
				s.getIPAddress,
			)
			if err := wakeScanner.Open(); err != nil {
				log.Printf("could not open wake scanner: %v\n", err)

				return
			}

			go func() {
				if err := wakeScanner.Transmit(); err != nil {
					log.Printf("could not transmit from wake scanner: %v\n", err)

					return
				}
			}()

			for {
				node := wakeScanner.Read()

				// Wake scan is done
				if node == nil {
					dbNodeWake.Done = 1
				}

				// Update and broadcast node wake in DB
				if node.Awake {
					dbNodeWake.PoweredOn = 1
					dbNodeWake.Done = 1
				} else {
					dbNodeWake.PoweredOn = 0
				}
				if err := s.nodeWakeDatabase.UpdateNodeWake(dbNodeWake); err != nil {
					log.Printf("could not update node wake in DB: %v\n", err)

					return
				}
				s.nodeWakeMessenger.Broadcast(dbNodeWake)

				if dbNodeWake.Done == 1 {
					break
				}
			}
		}
	}()

	protoNodeWake := &proto.NodeWakeNeoMessage{
		CreatedAt: dbNodeWake.CreatedAt.String(),
		Done: func() bool {
			if dbNodeWake.Done == 1 {
				return true
			}

			return false
		}(),
		ID:         dbNodeWake.ID,
		MACAddress: dbNodeWake.MacAddress,
		PoweredOne: func() bool {
			if dbNodeWake.PoweredOn == 1 {
				return true
			}

			return false
		}(),
	}

	return protoNodeWake, nil
}
