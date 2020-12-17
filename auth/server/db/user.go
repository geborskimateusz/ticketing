package db

import "go.mongodb.org/mongo-driver/bson/primitive"

// User datamodel
type User struct {
	ID       primitive.ObjectID `bson:"_id"`
	Email    string             `bson:"email"`
	Password string             `bson:"password"`
}
