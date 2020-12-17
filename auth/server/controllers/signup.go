package controllers

import (
	"context"
	"errors"
	"log"
	"net/http"

	"github.com/geborskimateusz/auth/server/customerr"
	"github.com/geborskimateusz/auth/server/db"
	"github.com/geborskimateusz/auth/server/entity"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
)

func Signup(c *gin.Context) {
	log.Println("Hit endpoitnt: api/users/signup")

	var user entity.User
	if err := c.ShouldBindBodyWith(&user, binding.JSON); err == nil {
		validate := validator.New()
		if err := validate.Struct(&user); err != nil {
			validationErrors := err.(validator.ValidationErrors)
			c.Error(customerr.NewRequestValidationError(validationErrors))
			return
		}
	}

	client, err := db.GetMongoClient()
	if err != nil {
		c.Error(customerr.NewDataBaseConnectionError(err))
	}

	collection := client.Database(db.DBNAME).Collection(db.USERS)

	filterCursor, err := collection.Find(context.TODO(), bson.M{"email": user.Email})
	if err != nil {
		log.Fatal(err)
	}
	var usersFiltered []bson.M
	if err = filterCursor.All(context.TODO(), &usersFiltered); err != nil {
		log.Fatal(err)
	}

	if len(usersFiltered) != 0 {
		c.Error(errors.New("Email already in use"))
		return
	}

	saved, err := collection.InsertOne(context.TODO(), user)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, saved)

}
