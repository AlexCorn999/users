package main

import (
	"context"
	"log"
	"os/signal"
	"syscall"

	"github.com/AlexCorn999/users/internal/config"
	"github.com/AlexCorn999/users/internal/transport"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	server := transport.NewAPIServer(config.NewConfig())
	if err := server.Start(ctx); err != nil {
		log.Fatal(err)
	}
}
