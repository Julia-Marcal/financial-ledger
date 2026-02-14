package server

import (
	"context"
	"financial-ledger/internal/infrastructure/mongodb"
	"financial-ledger/internal/router"
	"log/slog"
	"net/http"
	"time"
)

func Start(ctx context.Context) error {
	cleanup, err := initMongo(ctx)
	if err != nil {
		slog.Debug("failed to initialize mongodb", "error", err)
		cleanup = nil
	}

	// Initialize and start the HTTP server
	routerSrv := router.Router()
	serverErr := make(chan error, 1)
	go func() {
		if err := routerSrv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			serverErr <- err
		} else {
			serverErr <- nil
		}
	}()

	// Wait for either context cancellation or server error
	select {
	case <-ctx.Done():
		slog.Debug("Graceful Shutdown requested")
	case err := <-serverErr:
		if err != nil {
			return err
		}
	}

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := routerSrv.Shutdown(shutdownCtx); err != nil {
		slog.Debug("Server Shutdown Failed: %+v", err)
	}

	if cleanup != nil {
		cleanup()
	}

	slog.Debug("Graceful shutdown complete")
	return nil
}

func initMongo(parentCtx context.Context) (func(), error) {
	connCtx, cancel := context.WithTimeout(parentCtx, 5*time.Second)
	client, err := mongodb.Connect(connCtx)
	cancel()
	if err != nil {
		return nil, err
	}

	cleanup := func() {
		if err := mongodb.Disconnect(context.Background(), client); err != nil {
			slog.Debug("error disconnecting mongodb", "error", err)
		}
	}

	return cleanup, nil
}
