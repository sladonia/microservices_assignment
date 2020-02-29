package config

import "github.com/jinzhu/configor"

var Config Configuration

type Configuration struct {
	ServiceName string `env:"SERVICE_NAME" default:"port_domain_service"`
	Env         string `env:"ENV" default:"dev"`
	LogLevel    string `env:"LOG_LEVEL" default:"debug"`
	Port        string `env:"PORT" default:":50051"`
}

func Load() error {
	return configor.Load(&Config)
}
