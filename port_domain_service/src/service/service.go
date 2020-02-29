package service

import (
	"context"
	"io"
	"port_domain_service/src/logger"
	"port_domain_service/src/portpb"
	"port_domain_service/src/storage"
)

var PortService portpb.PortServiceServer = &portService{}

type portService struct{}

func (s *portService) Get(ctx context.Context, r *portpb.GetPortRequest) (*portpb.Port, error) {
	return storage.Storage.Get(r.Abbreviation)
}

func (s *portService) Import(stream portpb.PortService_ImportServer) error {
	var numberInserted int32
	var numberUpdated int32
	for {
		port, err := stream.Recv()
		if err == io.EOF {
			logger.Logger.Infof("successful import. number_inserted: %d. number_updated: %d",
				numberInserted, numberUpdated)
			return stream.SendAndClose(&portpb.ImportResponse{
				NumberInserted:       numberInserted,
				NumberUpdated:        numberUpdated,
			})
		} else if err != nil {
			logger.Logger.Errorw("error retrieving data", "error", err)
			return err
		}
		dInserted, dUpdated := storage.Storage.Upsert(port)
		numberInserted += dInserted
		numberUpdated += dUpdated
	}
}
