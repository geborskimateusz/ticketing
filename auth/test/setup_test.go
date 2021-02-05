package test

import (
	"context"
	"log"
	"os"
	"sync"
	"testing"

	"github.com/benweissmann/memongo"
	"github.com/geborskimateusz/auth/api/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/* Used to create a singleton object of MongoDB client.
Initialized and exposed through  GetMongoClient().*/
var clientInstance *mongo.Client

//Used during creation of singleton client object in GetMongoClient().
var clientInstanceError error

//Used to execute client creation procedure only once.
var mongoOnce sync.Once

var mongoServer memongo.Server

// TestMain can control flow of startup and cleanup
func TestMain(m *testing.M) {
	mongoServer, err := memongo.Start("4.0.5")
	if err != nil {
		log.Fatal(err)
	}
	defer mongoServer.Stop()

	tests := m.Run()

	os.Exit(tests)
}

func GetMongoTestClient() (*mongo.Client, error) {
	//Perform connection creation operation only once.
	//Creates Singleton
	mongoOnce.Do(func() {
		// Set client options
		clientOptions := options.Client().ApplyURI(mongoServer.URI())
		// Connect to MongoDB
		client, err := mongo.Connect(context.TODO(), clientOptions)
		if err != nil {
			clientInstanceError = err
		}
		// Check the connection
		err = client.Ping(context.TODO(), nil)
		if err != nil {
			clientInstanceError = err
		}
		clientInstance = client
	})
	return clientInstance, clientInstanceError
}

func CleanupCollection() {
	collection, err := db.GetCollection()
	if err != nil {
		log.Println(err)
	}

	collection.DeleteMany(context.TODO(), bson.M{})
}
