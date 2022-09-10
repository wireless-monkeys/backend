package cmd

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"strconv"

	"github.com/asaskevich/EventBus"
	"github.com/pkg/errors"
	"github.com/wireless-monkeys/backend/pkg/api"
	"github.com/wireless-monkeys/backend/pkg/service"
	"google.golang.org/grpc"
)

type Config struct {
	Port int
}

func RunServer() error {
	// TODO: load config from env
	cfg := Config{
		Port: 4000,
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *&cfg.Port))
	if err != nil {
		return errors.WithStack(err)
	}

	bus := EventBus.New()

	s := grpc.NewServer()

	helloService := service.NewHelloServiceServer()
	api.RegisterHelloServiceServer(s, helloService)

	edgeService := service.NewEdgeServiceServer(bus)
	api.RegisterEdgeServiceServer(s, edgeService)

	qdbConfig := parseQdbEnv()
	dashboardService := service.NewDashboardServiceServer(qdbConfig, bus)
	api.RegisterDashboardServiceServer(s, dashboardService)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	go func() {
		<-c
		log.Printf("stopping...")
		s.Stop()
	}()

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
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
