package service

import (
	"context"

	api "github.com/wireless-monkeys/backend/pkg/api"
)

//go:generate protoc --go_out=../api --go_opt=paths=source_relative --go-grpc_out=../api --go-grpc_opt=paths=source_relative ../../proto/hello-service.proto --proto_path=../../proto

type helloServiceServer struct {
	api.UnimplementedHelloServiceServer
}

func NewHelloServiceServer() api.HelloServiceServer {
	return &helloServiceServer{}
}

func (s *helloServiceServer) SayHello(ctx context.Context, in *api.HelloRequest) (*api.HelloReply, error) {
	return &api.HelloReply{Message: "Hello " + in.GetName()}, nil
}
