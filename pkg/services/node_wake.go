package services

import (
	"context"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/golang/protobuf/ptypes/empty"
	api "github.com/pojntfx/liwasc/pkg/api/proto/v1"
	models "github.com/pojntfx/liwasc/pkg/db/sqlite/node_wake"
	"github.com/pojntfx/liwasc/pkg/persisters"
	"github.com/pojntfx/liwasc/pkg/scanners"
	"github.com/pojntfx/liwasc/pkg/validators"
	"github.com/pojntfx/liwasc/pkg/wakers"
	"github.com/ugjka/messenger"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type NodeWakeService struct {
	api.UnimplementedNodeWakeServiceServer

	device         string
	wakeOnLANWaker *wakers.WakeOnLANWaker

	nodeWakePersister *persisters.NodeWakePersister
	getIPAddress      func(macAddress string) (ipAddress string, err error)

	nodeWakeMessenger *messenger.Messenger

	contextValidator *validators.ContextValidator
}

func NewNodeWakeService(
	device string,
	wakeOnLANWaker *wakers.WakeOnLANWaker,

	nodeWakePersister *persisters.NodeWakePersister,
	getIPAddress func(macAddress string) (ipAddress string, err error),

	contextValidator *validators.ContextValidator,
) *NodeWakeService {
	return &NodeWakeService{
		device:         device,
		wakeOnLANWaker: wakeOnLANWaker,

		nodeWakePersister: nodeWakePersister,
		getIPAddress:      getIPAddress,

		nodeWakeMessenger: messenger.New(0, true),

		contextValidator: contextValidator,
	}
}

func (s *NodeWakeService) StartNodeWake(ctx context.Context, nodeWakeStartMessage *api.NodeWakeStartMessage) (*api.NodeWakeMessage, error) {
	// Authorize
	valid, err := s.contextValidator.Validate(ctx)
	if err != nil || !valid {
		return nil, status.Errorf(codes.Unauthenticated, "could not authorize: %v", err)
	}

	// Validate
	if nodeWakeStartMessage.GetNodeWakeTimeout() < 1 {
		return nil, status.Error(codes.InvalidArgument, "node wake timeout can't be lower than 1")
	}

	// Create and broadcast node wake in DB
	dbNodeWake := &models.NodeWake{
		Done:       0,
		PoweredOn:  0,
		MacAddress: nodeWakeStartMessage.GetMACAddress(),
	}
	if err := s.nodeWakePersister.CreateNodeWake(dbNodeWake); err != nil {
		log.Printf("could not create node wake in DB: %v\n", err)

		return nil, status.Errorf(codes.Unknown, "could not create node wake in DB")
	}
	s.nodeWakeMessenger.Broadcast(dbNodeWake)

	// Wake the node
	if err := s.wakeOnLANWaker.Write(dbNodeWake.MacAddress); err != nil {
		log.Printf("could not wake node: %v\n", err)

		return nil, status.Errorf(codes.Unknown, "could not wake node")
	}

	successfulFirstOpen := make(chan error)

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

				// Send first error message to client
				if i == 0 {
					successfulFirstOpen <- err
				}

				return
			}

			// Send first error message to client
			if i == 0 {
				successfulFirstOpen <- nil
			}

			go func() {
				if err := wakeScanner.Transmit(); err != nil {
					log.Printf("could not transmit from wake scanner: %v\n", err)

					return
				}
			}()

			for {
				node := wakeScanner.Read()

				// Update and broadcast node wake in DB
				if node != nil && node.Awake {
					dbNodeWake.PoweredOn = 1
					dbNodeWake.Done = 1
				} else {
					dbNodeWake.PoweredOn = 0

					// Wake scan is done
					if node == nil {
						dbNodeWake.Done = 1
					}
				}
				if err := s.nodeWakePersister.UpdateNodeWake(dbNodeWake); err != nil {
					log.Printf("could not update node wake in DB: %v\n", err)

					return
				}
				s.nodeWakeMessenger.Broadcast(dbNodeWake)

				if dbNodeWake.Done == 1 {
					break
				}
			}

			if dbNodeWake.Done == 1 {
				break
			}
		}
	}()

	err = <-successfulFirstOpen
	if err != nil {
		if strings.Contains(err.Error(), "sql: no rows in result set") {
			return nil, status.Errorf(codes.NotFound, "could not find node to wake. Did you run a network scan yet?")
		}

		return nil, status.Errorf(codes.Unknown, "could not wake node")
	}

	protoNodeWake := &api.NodeWakeMessage{
		CreatedAt: dbNodeWake.CreatedAt.Format(time.RFC3339),
		Done: func() bool {
			if dbNodeWake.Done == 1 {
				return true
			}

			return false
		}(),
		ID:         dbNodeWake.ID,
		MACAddress: dbNodeWake.MacAddress,
		PoweredOn: func() bool {
			if dbNodeWake.PoweredOn == 1 {
				return true
			}

			return false
		}(),
	}

	return protoNodeWake, nil
}

func (s *NodeWakeService) SubscribeToNodeWakes(_ *empty.Empty, stream api.NodeWakeService_SubscribeToNodeWakesServer) error {
	// Authorize
	valid, err := s.contextValidator.Validate(stream.Context())
	if err != nil || !valid {
		return status.Errorf(codes.Unauthenticated, "could not authorize: %v", err)
	}

	var wg sync.WaitGroup

	wg.Add(2)

	// Get node wakes from messenger (priority 2)
	go func() {
		dbNodeWakes, err := s.nodeWakeMessenger.Sub()
		if err != nil {
			log.Printf("could not get node wakes from messenger: %v\n", err)

			return
		}
		defer s.nodeWakeMessenger.Unsub(dbNodeWakes)

		for dbNodeWake := range dbNodeWakes {
			protoNodeWake := &api.NodeWakeMessage{
				CreatedAt: dbNodeWake.(*models.NodeWake).CreatedAt.Format(time.RFC3339),
				Done: func() bool {
					if dbNodeWake.(*models.NodeWake).Done == 1 {
						return true
					}

					return false
				}(),
				ID:         dbNodeWake.(*models.NodeWake).ID,
				MACAddress: dbNodeWake.(*models.NodeWake).MacAddress,
				PoweredOn: func() bool {
					if dbNodeWake.(*models.NodeWake).PoweredOn == 1 {
						return true
					}

					return false
				}(),
				Priority: 2,
			}

			if err := stream.Send(protoNodeWake); err != nil {
				log.Printf("could send node wake %v to client: %v\n", protoNodeWake.ID, err)

				return
			}
		}

		wg.Done()
	}()

	// Get lookback node wakes from persister (priority 1)
	go func() {
		dbNodeWakes, err := s.nodeWakePersister.GetNodeWakes()
		if err != nil {
			log.Printf("could not get node wakes from DB: %v\n", err)

			return
		}

		for _, dbNodeWake := range dbNodeWakes {
			protoNodeWake := &api.NodeWakeMessage{
				CreatedAt: dbNodeWake.CreatedAt.Format(time.RFC3339),
				Done: func() bool {
					if dbNodeWake.Done == 1 {
						return true
					}

					return false
				}(),
				ID:         dbNodeWake.ID,
				MACAddress: dbNodeWake.MacAddress,
				PoweredOn: func() bool {
					if dbNodeWake.PoweredOn == 1 {
						return true
					}

					return false
				}(),
				Priority: 1,
			}

			if err := stream.Send(protoNodeWake); err != nil {
				log.Printf("could send node wake %v to client: %v\n", protoNodeWake.ID, err)

				return
			}
		}

		wg.Done()
	}()

	wg.Wait()

	return nil
}
