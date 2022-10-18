package mysql_driver

import (
	"learn-compute-services/drivers/mysql/users"
	"learn-compute-services/util"
	"errors"
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ConfigDB struct {
	DB_USERNAME string
	DB_PASSWORD string
	DB_NAME     string
	DB_HOST     string
	DB_PORT     string
}

// configure and connecting database
func (config *ConfigDB) InitDB() *gorm.DB {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.DB_USERNAME,
		config.DB_PASSWORD,
		config.DB_HOST,
		config.DB_PORT,
		config.DB_NAME,
	)

	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Fatalf("error when connecting to a database server: %s", err)
	}

	log.Println("connected to a database server")
	return db
}

// Migrating Struct into Table in Database
func DBMigrate(db *gorm.DB) {
	db.AutoMigrate(&users.User{})
}

// Closing Database
func CloseDB(db *gorm.DB) error {
	database, err := db.DB()
	if err != nil {
		log.Printf("error when getting the database instance : %v", err)
		return err
	}

	if err := database.Close(); err != nil {
		log.Printf("error when closing the database connection : %v", err)
		return err
	}
	log.Println("database connection is closed")
	return nil
}

// Seed for testing

func SeedUser(db *gorm.DB) users.User {
	password, _ := bcrypt.GenerateFromPassword([]byte("testing"), bcrypt.DefaultCost)
	fakeUser, _ := util.CreateFaker[users.User]()

	userRecord := users.User{
		Email:    fakeUser.Email,
		Password: string(password),
	}
	if err := db.Create(&userRecord).Error; err != nil {
		panic(err)
	}
	var lastUser users.User
	db.Last(&lastUser)

	lastUser.Password = "testing"
	return lastUser
}

func CleanSeeds(db *gorm.DB) {
	db.Exec("SET FOREIGN_KEY_CHECKS = 0")
	userResult := db.Exec("DELETE FROM users")

	if userResult != nil {
		panic(errors.New("error when cleaning up users seeders"))
	}
	log.Println("Seeders are cleaned up successfully")
}
