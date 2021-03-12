package db

import (
	"context"

	entity "github.com/geborskimateusz/auth/api/entity"
	common "github.com/geborskimateusz/ticketing-common"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// GetCollection returns collection of User Documents
func GetCollection() (*mongo.Collection, error) {
	client, err := GetMongoClient()
	if err != nil {
		return nil, common.NewDataBaseConnectionError(err)
	}
	return client.Database(DBNAME).Collection(USERS), nil
}

// Filter creates a map used in FindBy query
func Filter(field, value string) primitive.M {
	return bson.M{field: value}
}

// FindBy find User by given filter
func FindBy(filter primitive.M) (entity.UserDoc, error) {
	collection, err := GetCollection()
	if err != nil {
		return entity.UserDoc{}, err
	}

	var result entity.UserDoc
	err = collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return entity.UserDoc{}, nil
	}

	return result, nil
}

// Create UserDoc
func Create(user entity.User) (*entity.UserDoc, error) {
	collection, err := GetCollection()
	if err != nil {
		return nil, err
	}

	userDoc := entity.NewUserDoc(user)
	_, err = collection.InsertOne(context.TODO(), userDoc)
	if err != nil {
		return nil, err
	}

	return &userDoc, nil
}
