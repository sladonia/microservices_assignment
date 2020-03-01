package db

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"port_domain_service/src/config"
)

var Client *mongo.Client

func getOptions(username, password, host, port, dbName string) *options.ClientOptions {
	return options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s@%s:%s/%s",
		username, password, host, port, dbName))
}

func Connect(cfg config.MongodbConfig) (err error) {
	opts := getOptions(cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DbName)
	Client, err = mongo.Connect(context.Background(), opts)
	if err != nil {
		return
	}
	return Client.Ping(context.Background(), nil)
}
