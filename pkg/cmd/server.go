package cmd

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

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

	s := grpc.NewServer()

	helloService := service.NewHelloServiceServer()
	api.RegisterHelloServiceServer(s, helloService)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	go func() {
		<-c
		log.Printf("stopping...")
		s.GracefulStop()
	}()

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		return errors.WithStack(err)
	}
	return nil
}
