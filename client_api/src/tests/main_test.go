// Package tests implements end to end API tests
package tests

import (
	"client_api/src/app"
	"client_api/src/config"
	"client_api/src/grpc_client"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var (
	testServer *httptest.Server
	client     *http.Client
)

func TestMain(m *testing.M) {
	if err := config.Load(); err != nil {
		panic(err)
	}
	err := grpc_client.InitGrpcConnection(config.Config.PortDomain.Host,
		config.Config.PortDomain.Port)
	if err != nil {
		panic(err)
	}
	r := app.ConfigureRouter(config.Config.Port)
	testServer = httptest.NewServer(r)
	client = testServer.Client()
	os.Exit(m.Run())
}
