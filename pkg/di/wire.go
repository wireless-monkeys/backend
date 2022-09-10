//go:build wireinject
// +build wireinject

package di

import (
	"github.com/asaskevich/EventBus"
	"github.com/google/wire"
	"github.com/wireless-monkeys/backend/pkg/api"
	"github.com/wireless-monkeys/backend/pkg/config"
	"github.com/wireless-monkeys/backend/pkg/service"
	"google.golang.org/grpc"
)

func InitializeServer() *grpc.Server {
	wire.Build(
		config.NewConfig,
		service.NewDashboardServiceServer,
		service.NewEdgeServiceServer,
		service.NewHelloServiceServer,
		provideGrpcServer,
		EventBus.New,
	)
	return &grpc.Server{}
}

func provideGrpcServer(
	dashboardService api.DashboardServiceServer,
	edgeService api.EdgeServiceServer,
	helloService api.HelloServiceServer,
) *grpc.Server {
	s := grpc.NewServer()
	api.RegisterDashboardServiceServer(s, dashboardService)
	api.RegisterEdgeServiceServer(s, edgeService)
	api.RegisterHelloServiceServer(s, helloService)
	return s
}
