package config

import "github.com/jinzhu/configor"

var Config Configuration

type Configuration struct {
	ServiceName     string `env:"SERVICE_NAME" default:"client_api"`
	Env             string `env:"ENV" default:"dev"`
	LogLevel        string `env:"LOG_LEVEL" default:"debug"`
	Port            string `env:"PORT" default:":8080"`
	ShutdownTimeout int    `env:"SHUTDOWN_TIMEOUT" default:"20"` // seconds
	PortDomain      PortDomain
}

type PortDomain struct {
	Host string `env:"PORT_DOMAIN_HOST" default:"localhost"`
	Port string `env:"PORT_DOMAIN_PORT" default:"50051"`
}

func Load() error {
	return configor.Load(&Config)
}
