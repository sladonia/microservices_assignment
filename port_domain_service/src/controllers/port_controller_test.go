package controllers

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/metadata"
	"io"
	"os"
	"port_domain_service/src/config"
	"port_domain_service/src/db"
	"port_domain_service/src/portpb"
	"port_domain_service/src/services"
	"testing"
)

var recvResCh = GetNextRecvResult()

type recvRes struct {
	Port *portpb.Port
	err  error
}

func GetNextRecvResult() chan recvRes {
	ch := make(chan recvRes)

	recvMockData := []recvRes{
		{
			Port: &portpb.Port{
				Abbreviation: "QFSAS",
				Name:         "NamaNama",
				Coordinates:  []float64{123.3, 23.112},
				City:         "Tul",
				Province:     "Rayma",
				Country:      "Tuntur",
				Alias:        []string{},
				Regions:      []string{},
				Timezone:     "Afg/sdf",
				Unlocs:       []string{},
			},
			err: nil,
		},
		{
			Port: &portpb.Port{
				Abbreviation: "LDIHS",
				Name:         "EoXo",
				Coordinates:  []float64{123.3, 23.112},
				City:         "Tul",
				Province:     "Rayma",
				Country:      "Tuntur",
				Alias:        []string{},
				Regions:      []string{},
				Timezone:     "Afg/sdf",
				Unlocs:       []string{},
			},
			err: nil,
		},
		{
			nil,
			io.EOF,
		},
	}

	go func() {
		for _, arg := range recvMockData {
			ch <- arg
		}
		close(ch)
	}()

	return ch
}

func TestMain(m *testing.M) {
	err := godotenv.Load("../../.env.test")
	if err != nil {
		fmt.Println("unable to load test env")
		os.Exit(1)
	}
	if err := config.Load(); err != nil {
		panic(err)
	}

	err = db.Connect(config.Config.DbConfig)
	if err != nil {
		panic(err)
	}
	os.Exit(m.Run())
}

func Test_portController_Get(t *testing.T) {
	defer services.ClearPortCollection()

	mockStream := &MochPortService_ImportServer{}
	PortController.Import(mockStream)

	t.Run("port exists after import", func(tt *testing.T) {
		request := &portpb.GetPortRequest{
			Abbreviation: "QFSAS",
		}
		port, err := PortController.Get(context.Background(), request)
		assert.Nil(tt, err)
		assert.NotNil(tt, port)
		assert.Equal(tt, "NamaNama", port.Name)
	})
}

func Test_portController_Import(t *testing.T) {
	defer services.ClearPortCollection()

	t.Run("no such port in db", func(tt *testing.T) {
		request := &portpb.GetPortRequest{
			Abbreviation: "AEDXB",
		}
		port, err := PortController.Get(context.Background(), request)
		assert.Nil(tt, port)
		assert.NotNil(tt, err)
	})
}

type MochPortService_ImportServer struct{}

func (m *MochPortService_ImportServer) SendAndClose(*portpb.ImportResponse) error {
	return nil
}

func (m *MochPortService_ImportServer) Recv() (*portpb.Port, error) {
	res := <-recvResCh
	return res.Port, res.err
}

func (m *MochPortService_ImportServer) SetHeader(metadata.MD) error {
	return nil
}

func (m *MochPortService_ImportServer) SendHeader(metadata.MD) error {
	return nil
}

func (m *MochPortService_ImportServer) SetTrailer(metadata.MD) {}

func (m *MochPortService_ImportServer) Context() context.Context {
	return context.Background()
}

func (m *MochPortService_ImportServer) SendMsg(msg interface{}) error {
	return nil
}

func (m *MochPortService_ImportServer) RecvMsg(msg interface{}) error {
	return nil
}
