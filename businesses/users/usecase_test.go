package users_test

import (

	"learn-compute-services/app/middlewares"
	"learn-compute-services/businesses/users"
	_userMock "learn-compute-services/businesses/users/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	usersRepository _userMock.Repository
	usersService    users.Usecase

	usersDomain users.Domain
)

func TestMain(m *testing.M) {
	usersService = users.NewUserUseCase(&usersRepository, &middlewares.ConfigJWT{})
	usersDomain = users.Domain{
		Email:    "testing123@testing.com",
		Password: "testing123",
	}
	m.Run()
}

func TestGetAllUsers(t *testing.T) {
	t.Run("GetAllUsers | Valid", func(t *testing.T) {
		usersRepository.On("GetAllUsers").Return([]users.Domain{usersDomain}).Once()

		result := usersService.GetAllUsers()

		assert.Equal(t, 1, len(result))
	})

	t.Run("Get All | InValid", func(t *testing.T) {
		usersRepository.On("GetAllUsers").Return([]users.Domain{}).Once()

		result := usersService.GetAllUsers()

		assert.Equal(t, 0, len(result))
	})
}

func TestCreateUser(t *testing.T) {
	t.Run("CreateUser | Valid", func(t *testing.T) {
		usersRepository.On("CreateUser", &usersDomain).Return(users.Domain{}).Once()
		result := usersService.CreateUser(&usersDomain)

		assert.NotNil(t, result)
	})

	t.Run("CreateUser | InValid", func(t *testing.T) {
		usersRepository.On("CreateUser", &users.Domain{}).Return(users.Domain{}).Once()

		result := usersService.CreateUser(&users.Domain{})

		assert.NotNil(t, result)
	})
}

func TestRegister(t *testing.T) {
	t.Run("Register | Valid", func(t *testing.T) {
		usersRepository.On("Register", &usersDomain).Return(users.Domain{}).Once()
		result := usersService.Register(&usersDomain)

		assert.NotNil(t, result)
	})

	t.Run("Register | InValid", func(t *testing.T) {
		usersRepository.On("Register", &users.Domain{}).Return(users.Domain{}).Once()

		result := usersService.Register(&users.Domain{})

		assert.NotNil(t, result)
	})
}

func TestLogin(t *testing.T) {
	t.Run("Login | Valid", func(t *testing.T) {
		usersRepository.On("Login", &usersDomain).Return(users.Domain{}).Once()

		result := usersService.Login(&usersDomain)

		assert.NotNil(t, result)
	})

	t.Run("Login | InValid", func(t *testing.T) {
		usersRepository.On("Login", &users.Domain{}).Return(users.Domain{}).Once()

		result := usersService.Login(&users.Domain{})

		assert.Empty(t, result)
	})
}

func TestGetUser(t *testing.T) {
	t.Run("Get User | Valid", func(t *testing.T) {
		usersRepository.On("GetUser", "1").Return(usersDomain).Once()
		result := usersService.GetUser("1")
		assert.NotNil(t, result)
	})

	t.Run("Get User | InValid", func(t *testing.T) {
		usersRepository.On("GetUser", "0").Return(users.Domain{}).Once()
		result := usersService.GetUser("0")
		assert.NotNil(t, result)
	})
}
