package controllers

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"io"
	"port_domain_service/src/config"
	"port_domain_service/src/db"
	"port_domain_service/src/domains"
	"port_domain_service/src/logger"
	"port_domain_service/src/portpb"
	"port_domain_service/src/services"
)

var PortController portpb.PortServiceServer = &portController{}

type portController struct{}

// Get calls the storage service package functions to retrieve the Port data form the db
func (s *portController) Get(ctx context.Context, r *portpb.GetPortRequest) (*portpb.Port, error) {
	collection := db.Client.Database(config.Config.DbConfig.DbName).Collection(
		config.Config.DbConfig.Collection)
	return services.StorageService.GetOne(collection, r.Abbreviation)
}

// Import gathers streamed Port data in chunks and calls storage service package functions
// to perform bulk save to the db
// use SAVE_PORT_CHUNK_SIZE env var to set the desired chunk size
func (s *portController) Import(stream portpb.PortService_ImportServer) error {
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
	collection := db.Client.Database(config.Config.DbConfig.DbName).Collection(
		config.Config.DbConfig.Collection)
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
	nIns, nUpd, err := services.StorageService.UpsertMany(col, ports)
	resp.NumberInserted += nIns
	resp.NumberUpdated += nUpd
	if err != nil {
		logger.Logger.Infow("db error while importing ports", "error", err)
		resp.EncounterErrors = true
	}
}
