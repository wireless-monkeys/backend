//go:build wireinject
// +build wireinject

package di

import (
	"log"
	"os"
	"strconv"

	"github.com/asaskevich/EventBus"
	"github.com/google/wire"
	"github.com/wireless-monkeys/backend/pkg/api"
	"github.com/wireless-monkeys/backend/pkg/service"
	"google.golang.org/grpc"
)

func InitializeServer() *grpc.Server {
	wire.Build(
		service.NewDashboardServiceServer,
		service.NewEdgeServiceServer,
		service.NewHelloServiceServer,
		parseQdbEnv,
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

func parseQdbEnv() *service.QdbConfig {
	qdbHost := getEnv("QDB_HOST", "localhost")
	qdbPort, err := strconv.Atoi(getEnv("QDB_PORT", "8812"))
	if err != nil {
		log.Fatal(err)
	}
	qdbUser := getEnv("QDB_USER", "admin")
	qdbPassword := getEnv("QDB_PASSWORD", "quest")
	qdbDbname := getEnv("QDB_DBNAME", "qdb")
	qdbSslMode := getEnv("QDB_SSLMODE", "disable")

	return &service.QdbConfig{
		Host:     qdbHost,
		Port:     qdbPort,
		User:     qdbUser,
		Password: qdbPassword,
		Dbname:   qdbDbname,
		SslMode:  qdbSslMode,
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
