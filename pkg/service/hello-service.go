package service

import (
	"context"

	api "github.com/wireless-monkeys/backend/pkg/api"
)

type helloServiceServer struct {
	api.UnimplementedHelloServiceServer
}

func NewHelloServiceServer() api.HelloServiceServer {
	return &helloServiceServer{}
}

func (s *helloServiceServer) SayHello(ctx context.Context, in *api.HelloRequest) (*api.HelloReply, error) {
	return &api.HelloReply{Message: "Hello " + in.GetName()}, nil
}
