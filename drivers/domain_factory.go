package drivers

import (
	userDomain "learn-compute-services/businesses/users"
	userDB "learn-compute-services/drivers/mysql/users"

	"gorm.io/gorm"
)

func NewUserRepository(conn *gorm.DB) userDomain.Repository {
	return userDB.NewMySQLRepository(conn)
}
