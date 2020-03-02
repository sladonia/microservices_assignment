package config

import "github.com/jinzhu/configor"

var Config Configuration

type Configuration struct {
	ServiceName       string `env:"SERVICE_NAME" default:"port_domain_service"`
	Env               string `env:"ENV" default:"dev"`
	LogLevel          string `env:"LOG_LEVEL" default:"debug"`
	Port              string `env:"PORT" default:":50051"`
	SavePortChunkSize int    `env:"SAVE_PORT_CHUNK_SIZE" default:"400"`
	DbConfig          MongodbConfig
}

type MongodbConfig struct {
	User       string `env:"MONGO_USER" default:"user"`
	Password   string `env:"MONGO_PASSWORD" default:"password"`
	Host       string `env:"MONGO_HOST" default:"localhost"`
	Port       string `env:"MONGO_PORT" default:"27017"`
	DbName     string `env:"MONGO_DB_NAME" default:"port_db"`
	Collection string `env:"MONGO_DB_COLLECTION" default:"port"`
}

func Load() error {
	return configor.Load(&Config)
}
