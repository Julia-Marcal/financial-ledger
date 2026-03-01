package main

import (
	"context"
	"financial-ledger/internal/infraestructure/server"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	err := server.Start(ctx)
	if err != nil {
		slog.Error("Error when starting server: %v", err.Error())
		os.Exit(1)
	}
}
