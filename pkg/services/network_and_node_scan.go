package services

//go:generate sh -c "mkdir -p ../proto/generated && protoc --go_out=paths=source_relative,plugins=grpc:../proto/generated -I=../proto ../proto/*.proto"

import (
	"context"
	"log"
	"time"

	cmap "github.com/orcaman/concurrent-map"
	"github.com/pojntfx/liwasc/pkg/databases"
	proto "github.com/pojntfx/liwasc/pkg/proto/generated"
	"github.com/pojntfx/liwasc/pkg/scanners"
	liwascModels "github.com/pojntfx/liwasc/pkg/sql/generated/liwasc"
	mac2vendorModels "github.com/pojntfx/liwasc/pkg/sql/generated/mac2vendor"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type NetworkAndNodeScanService struct {
	proto.UnimplementedNetworkAndNodeScanServiceServer

	device             string
	networkScanners    cmap.ConcurrentMap
	mac2VendorDatabase *databases.MAC2VendorDatabase
	liwascDatabase     *databases.LiwascDatabase
}

func NewNetworkAndNodeScanService(device string, mac2VendorDatabase *databases.MAC2VendorDatabase, liwascDatabase *databases.LiwascDatabase) *NetworkAndNodeScanService {
	return &NetworkAndNodeScanService{device: device, networkScanners: cmap.New(), mac2VendorDatabase: mac2VendorDatabase, liwascDatabase: liwascDatabase}
}

func (s *NetworkAndNodeScanService) TriggerNetworkScan(ctx context.Context, scanTriggerMessage *proto.NetworkScanTriggerMessage) (*proto.NetworkScanReferenceMessage, error) {
	scanID, err := s.liwascDatabase.CreateScan(&liwascModels.Scan{})
	if err != nil {
		return nil, status.Errorf(codes.Unknown, "could not create scan in DB: %v", err.Error())
	}

	networkScanner := scanners.NewNetworkScanner(s.device)
	err, _ = networkScanner.Open()
	if err != nil {
		return nil, status.Errorf(codes.Unknown, "could not open network scanner: %v", err.Error())
	}
	s.networkScanners.Set(string(scanID), networkScanner)

	// Receive packets
	go func() {
		receiveCtx, cancel := context.WithTimeout(context.Background(), time.Millisecond*time.Duration(scanTriggerMessage.GetTimeout()))
		defer cancel()

		if err := networkScanner.Receive(receiveCtx); err != nil {
			log.Println("could not receive from network scanner", err)

			return
		}
	}()

	// Transmit packets ("start a scan")
	go func() {
		if err := networkScanner.Transmit(); err != nil {
			log.Println("could not transmit from network scanner", err)

			return
		}
	}()

	// Read packets
	go func() {
		for {
			node := networkScanner.Read()

			// Network scan is done
			if node == nil {
				break
			}

			// Lookup vendor information for node
			vendor, err := s.mac2VendorDatabase.GetVendor(node.MACAddress.String())
			if err != nil {
				vendor = &mac2vendorModels.Vendordb{}
			}

			dbNode := &liwascModels.Node{
				ScanID:       scanID,
				PoweredOn:    0,
				MacAddress:   node.MACAddress.String(),
				IPAddress:    node.IPAddress.String(),
				Vendor:       vendor.Vendor.String,
				Registry:     vendor.Registry,
				Organization: vendor.Organization.String,
				Address:      vendor.Address.String,
				Visible:      vendor.Visibility,
			}

			if _, err := s.liwascDatabase.CreateNode(dbNode, scanID); err != nil {
				log.Println("could not create node in DB", err)

				return
			}
		}
	}()

	return &proto.NetworkScanReferenceMessage{NetworkScanID: scanID}, nil
}

func (s *NetworkAndNodeScanService) SubscribeToNewNodes(*proto.NetworkScanReferenceMessage, proto.NetworkAndNodeScanService_SubscribeToNewNodesServer) error {
	return status.Errorf(codes.Unimplemented, "method SubscribeToNewNodes not implemented")
}
