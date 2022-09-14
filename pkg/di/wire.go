//go:build wireinject
// +build wireinject

package di

import (
	"github.com/asaskevich/EventBus"
	"github.com/google/wire"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	api2 "github.com/influxdata/influxdb-client-go/v2/api"
	"github.com/wireless-monkeys/backend/pkg/api"
	"github.com/wireless-monkeys/backend/pkg/config"
	"github.com/wireless-monkeys/backend/pkg/service"
	"google.golang.org/grpc"
)

func InitializeServer() (*grpc.Server, error) {
	wire.Build(
		config.NewConfig,
		service.NewDashboardServiceServer,
		service.NewEdgeServiceServer,
		service.NewHelloServiceServer,
		provideGrpcServer,
		EventBus.New,
		provideInfluxWriteAPI,
	)
	return &grpc.Server{}, nil
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

func provideInfluxWriteAPI(cfg *config.Config) api2.WriteAPIBlocking {
	influxConfig := cfg.InfluxDBConfig
	client := influxdb2.NewClient(influxConfig.Host, influxConfig.Token)
	writeAPI := client.WriteAPIBlocking(influxConfig.Organization, influxConfig.Bucket)
	return writeAPI
}
