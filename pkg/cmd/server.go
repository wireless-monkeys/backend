package cmd

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	"github.com/pkg/errors"
	"github.com/wireless-monkeys/backend/pkg/config"
	"github.com/wireless-monkeys/backend/pkg/di"
)

func RunServer() error {
	cfg := config.NewConfig()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Port))
	if err != nil {
		return errors.WithStack(err)
	}

	s := di.InitializeServer()

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
