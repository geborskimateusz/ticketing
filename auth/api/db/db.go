package db

import (
	"context"
	"log"

	"github.com/geborskimateusz/auth/api/entity"
	"github.com/geborskimateusz/auth/api/validation"
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

// Filter creates a map used in FindBy query
func Filter(field, value string) primitive.M {
	return bson.M{field: value}
}

// FindBy find User by given filter
func FindBy(filter primitive.M) ([]primitive.M, error) {
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

// Create UserDoc
func Create(user entity.User) (*entity.User, error) {
	collection, err := GetCollection()
	if err != nil {
		return nil, err
	}

	userDoc := entity.NewUserDoc(user)
	_, err = collection.InsertOne(context.TODO(), userDoc)
	if err != nil {
		return nil, err
	}

	return &entity.User{Email: userDoc.Email, Password: userDoc.Password}, nil
}
