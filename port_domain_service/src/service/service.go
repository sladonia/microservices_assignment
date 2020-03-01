package service

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"io"
	"port_domain_service/src/config"
	"port_domain_service/src/db"
	"port_domain_service/src/domains"
	"port_domain_service/src/logger"
	"port_domain_service/src/portpb"
)

var PortService portpb.PortServiceServer = &portService{}

type portService struct{}

func (s *portService) Get(ctx context.Context, r *portpb.GetPortRequest) (*portpb.Port, error) {
	//return storage.Storage.Get(r.Abbreviation)
	collection := db.Client.Database("port_db").Collection("ports")
	return domains.GetOne(collection, r.Abbreviation)
}

func (s *portService) Import(stream portpb.PortService_ImportServer) error {
	portsCh := make(chan *portpb.Port, config.Config.SavePortChunkSize)
	respCh := make(chan *portpb.ImportResponse)

	go GatherAndSave(portsCh, respCh)

	for {
		port, err := stream.Recv()
		if err == io.EOF {
			close(portsCh)
			resp := <-respCh
			logger.Logger.Infof("successful import. number_inserted: %d. number_updated: %d",
				resp.NumberInserted, resp.NumberUpdated)
			return stream.SendAndClose(resp)
		} else if err != nil {
			logger.Logger.Errorw("error retrieving data", "error", err)
			return err
		}
		portsCh <- port
	}
}

func GatherAndSave(portsCh <-chan *portpb.Port, respCh chan<- *portpb.ImportResponse) {
	resp := &portpb.ImportResponse{}
	collection := db.Client.Database("port_db").Collection("ports")
	ports := make([]*domains.Port, 0, config.Config.SavePortChunkSize)
	for port := range portsCh {
		p := domains.PortDomainFromPBPort(port)
		ports = append(ports, p)
		if len(ports) >= config.Config.SavePortChunkSize {
			saveAndUpdateResponse(ports, resp, collection)
			ports = ports[:0]
		}
	}
	saveAndUpdateResponse(ports, resp, collection)
	respCh <- resp
}

func saveAndUpdateResponse(ports []*domains.Port, resp *portpb.ImportResponse, col *mongo.Collection) {
	nIns, nUpd, err := domains.UpsertMany(col, ports)
	resp.NumberInserted += nIns
	resp.NumberUpdated += nUpd
	if err != nil {
		logger.Logger.Infow("db error while importing ports", "error", err)
		resp.EncounterErrors = true
	}
}
