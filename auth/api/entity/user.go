package entity

import (
	"time"

	"github.com/geborskimateusz/auth/api/util"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User is representation of model used to signin or signup
type User struct {
	Email    string `json:"email"  validate:"email" bson:"email"`
	Password string `json:"password" validate:"min=8,max=32,alphanum" bson:"password"`
}

// UserDoc is representation of model used to signin or signup
type UserDoc struct {
	ID        primitive.ObjectID `bson:"_id"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
	User      `bson:",inline"`
}

// NewUserDoc creates new User mongo document
func NewUserDoc(user User) UserDoc {
	return UserDoc{
		ID:        primitive.NewObjectID(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		User: User{
			Email:    user.Email,
			Password: util.HashAndSalt(user.Password),
		},
	}
}

func (u *UserDoc) AsJSON() (ret struct {
	ID    string
	Email string
}) {
	ret.ID = u.ID.Hex()
	ret.Email = u.Email
	return
}
