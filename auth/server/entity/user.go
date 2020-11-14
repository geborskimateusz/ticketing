package entity

// User is representation of model used to signin or signup
type User struct {
	Email    string `json:"email"  validate:"email"`
	Password string `json:"password" validate:"min=8,max=32,alphanum"`
}
