package users

import (
	"time"

	"gorm.io/gorm"
)

type Domain struct {
	ID        uint
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
type Usecase interface {
	GetAllUsers() []Domain
	CreateUser(userDomain *Domain) Domain
	Register(userDomain *Domain) Domain
	Login(userDomain *Domain) string
	GetUser(id string) Domain
}

type Repository interface {
	GetAllUsers() []Domain
	CreateUser(userDomain *Domain) Domain
	Register(userDomain *Domain) Domain
	Login(userDomain *Domain) Domain
	GetUser(id string) Domain
}
