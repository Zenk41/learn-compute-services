package users

import (
	"fmt"
	"learn-compute-services/businesses/users"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type userRepository struct {
	conn *gorm.DB
}

func NewMySQLRepository(conn *gorm.DB) users.Repository {
	return &userRepository{
		conn: conn,
	}
}

func (ur *userRepository) GetAllUsers() []users.Domain {
	var rec []User

	ur.conn.Find(&rec)

	userDomain := []users.Domain{}

	for _, user := range rec {
		userDomain = append(userDomain, user.ToDomain())
	}
	return userDomain
}

func (ur *userRepository) CreateUser(userDomain *users.Domain) users.Domain {
	password, _ := bcrypt.GenerateFromPassword([]byte(userDomain.Password), bcrypt.DefaultCost)
	rec := FromDomain(userDomain)
	rec.Password = string(password)
	result := ur.conn.Create(&rec)

	result.Last(&rec)

	return rec.ToDomain()
}

func (ur *userRepository) Register(userDomain *users.Domain) users.Domain {
	password, _ := bcrypt.GenerateFromPassword([]byte(userDomain.Password), bcrypt.DefaultCost)
	rec := FromDomain(userDomain)
	rec.Password = string(password)
	result := ur.conn.Create(&rec)

	result.Last(&rec)

	return rec.ToDomain()
}

func (ur *userRepository) Login(usersDomain *users.Domain) users.Domain {
	var user User
	ur.conn.First(&user, "email=?", usersDomain.Email)

	if user.ID == 0 {
		fmt.Println("user not found")
		return users.Domain{}
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(usersDomain.Password)); err != nil {
		fmt.Println("password failed")
		return users.Domain{}
	}
	return user.ToDomain()
}

func (ur *userRepository) GetUser(id string) users.Domain {
	var user User
	ur.conn.First(&user, "id=?", id)
	return user.ToDomain()
}
