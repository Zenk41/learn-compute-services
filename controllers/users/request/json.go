package request

import (
	"learn-compute-services/businesses/users"

	"github.com/go-playground/validator/v10"
)

type User struct {
	Email    string `json:"email" form:"email" validate:"required,email"`
	Password string `json:"password" form:"password" validate:"required"`
}

func (req *User) ToDomain() *users.Domain {
	return &users.Domain{
		Email:    req.Email,
		Password: req.Password,
	}
}

func (req *User) Validate() error {
	validate := validator.New()

	err := validate.Struct(req)
	return err
}
