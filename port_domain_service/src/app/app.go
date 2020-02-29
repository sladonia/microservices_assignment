package app

import (
	"port_domain_service/src/config"
	"port_domain_service/src/logger"
)

func ConfigureApp() error {
	if err := config.Load(); err != nil {
		return err
	}
	if err := logger.InitLogger(config.Config.ServiceName, config.Config.LogLevel); err != nil {
		return err
	}
	return nil
}
