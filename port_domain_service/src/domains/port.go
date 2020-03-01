package domains

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"port_domain_service/src/portpb"
)

type Port struct {
	Abbreviation string    `json:"abbreviation"`
	Name         string    `json:"name"`
	City         string    `json:"city"`
	Country      string    `json:"country"`
	Alias        []string  `json:"alias"`
	Regions      []string  `json:"regions"`
	Coordinates  []float64 `json:"coordinates"`
	Province     string    `json:"province"`
	Timezone     string    `json:"timezone"`
	Unlocs       []string  `json:"unlocs"`
	Code         string    `json:"code"`
}

func PortDomainFromPBPort(p *portpb.Port) *Port {
	return &Port{
		Abbreviation: p.Abbreviation,
		Name:         p.Name,
		City:         p.City,
		Country:      p.Country,
		Alias:        p.Alias,
		Regions:      p.Regions,
		Coordinates:  p.Coordinates,
		Province:     p.Province,
		Timezone:     p.Timezone,
		Unlocs:       p.Unlocs,
	}
}

func UpsertOne(collection *mongo.Collection, port *Port) (int32, int32, error) {
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

func GetOne(collection *mongo.Collection, abbreviation string) (*portpb.Port, error) {
	filter := bson.M{"abbreviation": abbreviation}

	var port portpb.Port
	if err := collection.FindOne(context.Background(), filter).Decode(&port); err != nil {
		return nil, err
	}
	return &port, nil
}

func UpsertMany(collection *mongo.Collection, ports []*Port) (int32, int32, error) {
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
