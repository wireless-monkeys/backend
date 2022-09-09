package main

import (
	"log"

	"github.com/wireless-monkeys/backend/pkg/cmd"
)

func main() {
	err := cmd.RunServer()
	if err != nil {
		log.Fatal(err)
	}
}
