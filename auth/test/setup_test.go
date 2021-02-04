package test

import (
	"log"
	"os"
	"testing"

	"github.com/benweissmann/memongo"
)

var mongoServer memongo.Server

// TestMain can control flow of startup and cleanup
func TestMain(m *testing.M) {
	mongoServer, err := memongo.Start("4.0.5")
	if err != nil {
		log.Fatal(err)
	}

	defer mongoServer.Stop()
	os.Exit(m.Run())
}
