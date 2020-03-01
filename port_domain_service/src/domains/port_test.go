package domains

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"os"
	"port_domain_service/src/config"
	"port_domain_service/src/db"
	"testing"
)

var portFixtures = []*Port{
	&Port{
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
	&Port{
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
}

func ClearPortCollection() {
	collection := db.Client.Database("port_db").Collection("ports")
	_, _ = collection.DeleteMany(context.Background(), bson.M{})
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

func TestUpsertOne(t *testing.T) {
	defer ClearPortCollection()

	collection := db.Client.Database("port_db").Collection("ports")

	t.Run("success insert", func(tt *testing.T) {
		numInserted, numUpdated, err := UpsertOne(collection, portFixtures[0])
		assert.Nil(t, err)
		assert.Equal(t, int32(1), numInserted)
		assert.Equal(t, int32(0), numUpdated)
	})

	t.Run("success update", func(tt *testing.T) {
		numInserted, numUpdated, err := UpsertOne(collection, portFixtures[0])
		assert.Nil(t, err)
		assert.Equal(t, int32(0), numInserted)
		assert.Equal(t, int32(1), numUpdated)
	})
}

func TestGetOne(t *testing.T) {
	defer ClearPortCollection()

	collection := db.Client.Database("port_db").Collection("ports")
	UpsertOne(collection, portFixtures[0])

	t.Run("success", func(tt *testing.T) {
		port, err := GetOne(collection, portFixtures[0].Abbreviation)
		assert.Nil(tt, err)
		assert.NotNil(tt, port)
		assert.Equal(tt, portFixtures[0].Name, port.Name)
	})

	t.Run("no such port", func(tt *testing.T) {
		port, err := GetOne(collection, "afdgafgadgsdfg")
		assert.NotNil(tt, err)
		assert.Nil(tt, port)
	})
}

func TestUpsertMany(t *testing.T) {
	defer ClearPortCollection()
	collection := db.Client.Database("port_db").Collection("ports")

	t.Run("success insert many", func(tt *testing.T) {
		numInserted, numUpdated, err := UpsertMany(collection, portFixtures)
		assert.Nil(tt, err)
		assert.Equal(tt, int32(2), numInserted)
		assert.Equal(tt, int32(0), numUpdated)
	})

	t.Run("success update many", func(tt *testing.T) {
		numInserted, numUpdated, err := UpsertMany(collection, portFixtures)
		assert.Nil(tt, err)
		assert.Equal(tt, int32(0), numInserted)
		assert.Equal(tt, int32(2), numUpdated)
	})
}
