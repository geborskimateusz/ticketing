package test

import (
	"log"
	"os"
	"sync"
	"testing"

	"github.com/benweissmann/memongo"
	"github.com/geborskimateusz/auth/api/db"
	"go.mongodb.org/mongo-driver/mongo"
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

	db.CONNECTIONSTRING = mongoServer.URI()
	tests := m.Run()

	os.Exit(tests)
}

