package services

import (
	"client_api/src/domains"
	"client_api/src/logger"
	"client_api/src/portpb"
	"context"
	"google.golang.org/grpc"
)

var PortService PortServiceInterface = &portService{}

type PortServiceInterface interface {
	Import(portCh <-chan portpb.Port, conn *grpc.ClientConn) (*domains.ImportResponse, error)
	Get(key string, conn *grpc.ClientConn) (*portpb.Port, error)
}

type portService struct{}

func (s *portService) Import(portCh <-chan portpb.Port, conn *grpc.ClientConn) (*domains.ImportResponse, error) {
	client := portpb.NewPortServiceClient(conn)
	stream, err := client.Import(context.Background())
	if err != nil {
		return nil, err
	}
	for port := range portCh {
		err = stream.Send(&port)
		if err != nil {
			logger.Logger.Errorw("error sending port down the channel", "error", err)
		}
	}
	resp, err := stream.CloseAndRecv()
	if err != nil {
		return nil, err
	}
	return &domains.ImportResponse{
		NumberInserted: resp.NumberInserted,
		NumberUpdated:  resp.NumberUpdated,
	}, nil
}

func (s *portService) Get(abbreviation string, conn *grpc.ClientConn) (*portpb.Port, error) {
	client := portpb.NewPortServiceClient(conn)
	return client.Get(context.Background(), &portpb.GetPortRequest{Abbreviation: abbreviation})
}
