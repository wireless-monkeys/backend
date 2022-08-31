package service

import (
	"context"

	api "github.com/wireless-monkeys/backend/pkg/api"
)

type edgeServiceServer struct {
	api.UnimplementedEdgeServiceServer
}

func NewEdgeServiceServer() api.EdgeServiceServer {
	return &edgeServiceServer{}
}

func (s *edgeServiceServer) Heartbeat(ctx context.Context, in *api.Empty) (*api.Empty, error) {
	return &api.Empty{}, nil
}
