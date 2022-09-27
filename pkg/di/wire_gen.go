// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"github.com/asaskevich/EventBus"
	"github.com/influxdata/influxdb-client-go/v2"
	api2 "github.com/influxdata/influxdb-client-go/v2/api"
	"github.com/wireless-monkeys/backend/pkg/api"
	"github.com/wireless-monkeys/backend/pkg/config"
	"github.com/wireless-monkeys/backend/pkg/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// Injectors from wire.go:

func InitializeServer() (*grpc.Server, error) {
	configConfig := config.NewConfig()
	bus := EventBus.New()
	client := provideInfluxClient(configConfig)
	queryAPI := provideInfluxQueryAPI(configConfig, client)
	dashboardServiceServer := service.NewDashboardServiceServer(configConfig, bus, queryAPI)
	writeAPIBlocking := provideInfluxWriteAPI(configConfig, client)
	edgeServiceServer := service.NewEdgeServiceServer(bus, writeAPIBlocking)
	helloServiceServer := service.NewHelloServiceServer()
	server := provideGrpcServer(dashboardServiceServer, edgeServiceServer, helloServiceServer)
	return server, nil
}

// wire.go:

func provideGrpcServer(
	dashboardService api.DashboardServiceServer,
	edgeService api.EdgeServiceServer,
	helloService api.HelloServiceServer,
) *grpc.Server {
	s := grpc.NewServer()
	api.RegisterDashboardServiceServer(s, dashboardService)
	api.RegisterEdgeServiceServer(s, edgeService)
	api.RegisterHelloServiceServer(s, helloService)
	reflection.Register(s)
	return s
}

func provideInfluxClient(cfg *config.Config) influxdb2.Client {
	influxConfig := cfg.InfluxDBConfig
	client := influxdb2.NewClient(influxConfig.Host, influxConfig.Token)
	return client
}

func provideInfluxWriteAPI(cfg *config.Config, client influxdb2.Client) api2.WriteAPIBlocking {
	influxConfig := cfg.InfluxDBConfig
	writeAPI := client.WriteAPIBlocking(influxConfig.Organization, influxConfig.Bucket)
	return writeAPI
}

func provideInfluxQueryAPI(cfg *config.Config, client influxdb2.Client) api2.QueryAPI {
	queryAPI := client.QueryAPI(cfg.InfluxDBConfig.Organization)
	return queryAPI
}
