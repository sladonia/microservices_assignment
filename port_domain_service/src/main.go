package main

import (
	"google.golang.org/grpc"
	"net"
	"os"
	"os/signal"
	"port_domain_service/src/app"
	"port_domain_service/src/config"
	"port_domain_service/src/logger"
	"port_domain_service/src/portpb"
	"port_domain_service/src/service"
	"syscall"
)

func main() {
	if err := app.ConfigureApp(); err != nil {
		panic(err)
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)

	lis, err := net.Listen("tcp", config.Config.Port)
	if err != nil {
		logger.Logger.Fatalw("can't start listener", "error", err)
	}
	logger.Logger.Infof("started tcp listener on port %s", config.Config.Port)

	s := grpc.NewServer()
	portpb.RegisterPortServiceServer(s, service.PortService)

	go func() {
		logger.Logger.Infof("starting port_domain grpc service")
		if err := s.Serve(lis); err != nil {
			logger.Logger.Fatalw("unable to start the server", "error", err)
		}
	}()

	<-done
	logger.Logger.Info("shutting down gracefully")
	logger.Logger.Sync()
}
