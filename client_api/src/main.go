package main

import (
	"client_api/src/app"
	"client_api/src/config"
	"client_api/src/grpc_client"
	"client_api/src/logger"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	if err := app.ConfigureApp(); err != nil {
		panic(err)
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)

	err := grpc_client.InitGrpcClient(config.Config.PortDomain.Host,
		config.Config.PortDomain.Port)
	if err != nil {
		panic(err)
	}

	r := app.InitApp()
	logger.Logger.Infof("start listening on port %s", config.Config.Port)
	go func() {
		if err := http.ListenAndServe(":8080", r); err != nil {
			logger.Logger.Fatalw("unable to start the server", "error", err)
		}
	}()

	<-done
	logger.Logger.Info("shutting down gracefully")
	logger.Logger.Sync()
}
