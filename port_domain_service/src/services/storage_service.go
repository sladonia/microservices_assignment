package services

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"port_domain_service/src/db"
	"port_domain_service/src/domains"
	"port_domain_service/src/portpb"
)

var StorageService StorageServiceInterface = &storageService{}

type StorageServiceInterface interface {
	// insert single Port record to the db
	UpsertOne(*mongo.Collection, *domains.Port) (int32, int32, error)
	// retrieve single Port record form the db
	GetOne(*mongo.Collection, string) (*portpb.Port, error)
	// perform bulk Port data upsert
	UpsertMany(*mongo.Collection, []*domains.Port) (int32, int32, error)
}

type storageService struct{}

func (s *storageService) UpsertOne(collection *mongo.Collection, port *domains.Port) (int32, int32, error) {
	op := options.Update().SetUpsert(true)
	filter := bson.M{"abbreviation": port.Abbreviation}
	update := bson.M{"$set": port}

	res, err := collection.UpdateOne(context.Background(), filter, update, op)
	if err != nil {
		return 0, 0, err
	}

	numModified := res.ModifiedCount
	if res.UpsertedCount == 0 && res.ModifiedCount == 0 {
		numModified = res.MatchedCount
	}

	return int32(res.UpsertedCount), int32(numModified), nil
}

func (s *storageService) GetOne(collection *mongo.Collection, abbreviation string) (*portpb.Port, error) {
	filter := bson.M{"abbreviation": abbreviation}

	var port portpb.Port
	if err := collection.FindOne(context.Background(), filter).Decode(&port); err != nil {
		return nil, err
	}
	return &port, nil
}

func (s *storageService) UpsertMany(collection *mongo.Collection, ports []*domains.Port) (int32, int32, error) {
	var operations []mongo.WriteModel
	for _, port := range ports {
		op := mongo.NewUpdateOneModel()
		op.Filter = bson.M{"abbreviation": port.Abbreviation}
		op.Update = bson.M{"$set": port}
		op.SetUpsert(true)
		operations = append(operations, op)
	}
	res, err := collection.BulkWrite(context.Background(), operations)
	if err != nil {
		return 0, 0, err
	}
	return int32(res.UpsertedCount), int32(res.MatchedCount), nil
}

func ClearPortCollection() {
	collection := db.Client.Database("port_db").Collection("ports")
	_, _ = collection.DeleteMany(context.Background(), bson.M{})
}