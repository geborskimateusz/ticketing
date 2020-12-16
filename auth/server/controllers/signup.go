package controllers

import (
	"context"
	"errors"
	"net/http"

	"github.com/geborskimateusz/auth/server/customerr"
	"github.com/geborskimateusz/auth/server/db"
	"github.com/geborskimateusz/auth/server/entity"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Signup(c *gin.Context) {

	var user entity.User
	if err := c.ShouldBindBodyWith(&user, binding.JSON); err == nil {
		validate := validator.New()
		if err := validate.Struct(&user); err != nil {
			validationErrors := err.(validator.ValidationErrors)
			c.Error(customerr.NewRequestValidationError(validationErrors))
			return
		}
	}

	found := db.User{}
	filter := bson.D{primitive.E{Key: "username", Value: user.Email}}

	client, err := db.GetMongoClient()
	if err != nil {
		c.Error(customerr.NewDataBaseConnectionError(err))
	}

	collection := client.Database(db.DBNAME).Collection(db.USERS)
	err = collection.FindOne(context.TODO(), filter).Decode(&found)
	if err != nil {
		c.Error(err)
	}

	if (db.User{}) == found {
		c.Error(errors.New("Email already in use"))
	}

	saved, err := collection.InsertOne(context.TODO(), user)
	if err != nil {
		c.Error(err)
	}

	c.JSON(http.StatusCreated, saved)

}
