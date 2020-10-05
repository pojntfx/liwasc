package services

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/pojntfx/liwasc/pkg/networking"
	proto "github.com/pojntfx/liwasc/pkg/proto/generated"
)

type MetadataService struct {
	proto.UnimplementedMetadataServiceServer

	subnets            []string
	device             string
	interfaceInspector *networking.InterfaceInspector
}

func NewMetadataService(interfaceInspector *networking.InterfaceInspector) *MetadataService {
	return &MetadataService{interfaceInspector: interfaceInspector}
}

func (s *MetadataService) Open() error {
	subnets, err := s.interfaceInspector.GetIPv4Subnets()
	if err != nil {
		return err
	}

	s.subnets = subnets
	s.device = s.interfaceInspector.GetDevice()

	return nil
}

func (s *MetadataService) GetMetadata(ctx context.Context, _ *empty.Empty) (*proto.MetadataMessage, error) {
	protoMetadataMessage := &proto.MetadataMessage{
		Subnets: s.subnets,
		Device:  s.device,
	}

	return protoMetadataMessage, nil
}
