package config

import "github.com/jinzhu/configor"

var Config Configuration

type Configuration struct {
	ServiceName string `env:"SERVICE_NAME" default:"client_api"`
	Env         string `env:"ENV" default:"dev"`
	LogLevel    string `env:"LOG_LEVEL" default:"debug"`
	Port        string `env:"PORT" default:":8080"`
	PortDomain  PortDomain
}

type PortDomain struct {
}

func Load() error {
	return configor.Load(&Config)
}
