package db

import (
	"context"
	"log"

	"github.com/geborskimateusz/auth/server/entity"
	"github.com/geborskimateusz/auth/server/validation"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// GetCollection returns collection of User Documents
func GetCollection() (*mongo.Collection, error) {
	client, err := GetMongoClient()
	if err != nil {
		return nil, validation.NewDataBaseConnectionError(err)
	}
	return client.Database(DBNAME).Collection(USERS), nil
}

// FindBy find User by given filter
func FindBy(filter map[string]interface{}) ([]primitive.M, error) {
	collection, err := GetCollection()
	if err != nil {
		return nil, err
	}
	filterCursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	var usersFiltered []bson.M
	if err = filterCursor.All(context.TODO(), &usersFiltered); err != nil {
		log.Fatal(err)
		return nil, err
	}

	return usersFiltered, nil
}

// Create creates User
func Create(user entity.User) (*mongo.InsertOneResult, error) {
	collection, err := GetCollection()
	if err != nil {
		return nil, err
	}
	return collection.InsertOne(context.TODO(), user)
}
