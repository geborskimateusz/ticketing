package entity

// User is representation of model used to signin or signup
type User struct {
	Username        string `json:"username" validate:"required"`
	Email           string `json:"email"  validate:"email"`
	Password        string `json:"password" validate:"min=8,max=32,alphanum"`
	ConfirmPassword string `json:"confirm_password" validate:"eqfield=Password,required"`
}
