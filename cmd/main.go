package main

import (
	"log"

	"github.com/AlexCorn999/users/internal/config"
	"github.com/AlexCorn999/users/internal/transport"
)

func main() {
	server := transport.NewAPIServer(config.NewConfig())
	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
}
