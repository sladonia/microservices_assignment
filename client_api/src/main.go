package main

import (
	"client_api/src/app"
	"client_api/src/config"
	"client_api/src/grpc_client"
	"client_api/src/logger"
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// Load config from env vars
	if err := config.Load(); err != nil {
		panic(err)
	}
	// Init logger
	err := logger.InitLogger(config.Config.ServiceName, config.Config.LogLevel)
	if err != nil {
		panic(err)
	}

	// Init channels to handle shutdown
	signalsCh := make(chan os.Signal, 1)
	doneCh := make(chan struct{})
	signal.Notify(signalsCh, os.Interrupt, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)

	// Connect to port_domain grpc service
	err = grpc_client.InitGrpcConnection(config.Config.PortDomain.Host,
		config.Config.PortDomain.Port)
	if err != nil {
		panic(err)
	}

	// Map routes
	r := app.ConfigureRouter(config.Config.Port)
	// Init http server
	server := http.Server{Addr: config.Config.Port, Handler: r}

	// Start http server
	go func() {
		logger.Logger.Infof("start listening on port %s", config.Config.Port)
		if err := server.ListenAndServe(); err != nil {
			logger.Logger.Infow("unable to start the server", "error", err)
		}
	}()

	// Clean up before shutdown
	// Stop process new incoming requests
	// Wait current requests to finish or shutdown by timeout
	go func() {
		<-signalsCh
		logger.Logger.Info("shutdown signal received. cleaning up...")

		ctx, cancel := context.WithTimeout(context.Background(),
			time.Duration(config.Config.ShutdownTimeout)*time.Second)
		defer cancel()
		err := server.Shutdown(ctx)
		if err != nil {
			logger.Logger.Errorw("unable to perform graceful shutdown", "error", err)
		}

		doneCh <- struct{}{}
	}()

	// Shutdown after the clean up
	<-doneCh
	logger.Logger.Info("shutting down gracefully")
	logger.Logger.Sync()
}
