package db

import "go.mongodb.org/mongo-driver/bson/primitive"

// User datamodel
type User struct {
	ID       primitive.ObjectID
	Username string `valid:"alphanum,required"`
	Password string `valid:"alphanum,required"`
}
