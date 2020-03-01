package main

import (
	"google.golang.org/grpc"
	"net"
	"os"
	"os/signal"
	"port_domain_service/src/config"
	"port_domain_service/src/logger"
	"port_domain_service/src/portpb"
	"port_domain_service/src/service"
	"syscall"
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

	// Start tcp listener
	lis, err := net.Listen("tcp", config.Config.Port)
	if err != nil {
		logger.Logger.Fatalw("can't start listener", "error", err)
	}
	logger.Logger.Infof("started tcp listener on port %s", config.Config.Port)

	// Init grpc service
	s := grpc.NewServer()
	portpb.RegisterPortServiceServer(s, service.PortService)

	// Start grpc service
	go func() {
		logger.Logger.Infof("starting port_domain grpc service")
		if err := s.Serve(lis); err != nil {
			logger.Logger.Fatalw("unable to start the server", "error", err)
		}
	}()

	// Clean up before shutdown
	go func() {
		<-signalsCh
		logger.Logger.Info("shutdown signal received. waiting for current requests to finish...")

		s.GracefulStop()

		doneCh <- struct{}{}
	}()

	// Shutdown after the clean up
	<-doneCh
	logger.Logger.Info("shutting down gracefully")
	logger.Logger.Sync()
}
