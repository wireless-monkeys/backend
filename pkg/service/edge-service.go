package service

import (
	"context"

	"github.com/asaskevich/EventBus"
	api "github.com/wireless-monkeys/backend/pkg/api"
)

type edgeServiceServer struct {
	api.UnimplementedEdgeServiceServer

	bus EventBus.Bus
}

func NewEdgeServiceServer(bus EventBus.Bus) api.EdgeServiceServer {
	return &edgeServiceServer{
		bus: bus,
	}
}

func (s *edgeServiceServer) Heartbeat(ctx context.Context, in *api.Empty) (*api.Empty, error) {
	return &api.Empty{}, nil
}

func (s *edgeServiceServer) SetData(ctx context.Context, in *api.SetDataRequest) (*api.Empty, error) {
	s.bus.Publish("edge:setdata", in)
	return &api.Empty{}, nil
}
