package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User is representation of model used to signin or signup
type User struct {
	Email    string `json:"email"  validate:"email" bson:"email"`
	Password string `json:"password" validate:"min=8,max=32,alphanum" bson:"password"`
}

// User is representation of model used to signin or signup
type UserDoc struct {
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id"`
	CreatedAt time.Time          `json:"createdAt,omitempty" bson:"created_at"`
	UpdatedAt time.Time          `json:"updatedAt,omitempty" bson:"updated_at"`
	*User
}

func NewUserDoc(user User) *UserDoc {
	return &UserDoc{
		ID:        primitive.NewObjectID(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		User: &User{
			Email:    user.Email,
			Password: user.Password,
		},
	}
}
