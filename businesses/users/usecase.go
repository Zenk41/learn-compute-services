package users

import "learn-compute-services/app/middlewares"

type UserUseCase struct {
	userRepository Repository
	jwtAuth        *middlewares.ConfigJWT
}

func NewUserUseCase(ur Repository, jwtAuth *middlewares.ConfigJWT) Usecase {
	return &UserUseCase{
		userRepository: ur,
		jwtAuth:        jwtAuth,
	}
}

func (uu *UserUseCase) GetAllUsers() []Domain {
	return uu.userRepository.GetAllUsers()
}

func (uu *UserUseCase) CreateUser(userDomain *Domain) Domain {
	return uu.userRepository.CreateUser(userDomain)
}
func (uu *UserUseCase) Register(userDomain *Domain) Domain {
	return uu.userRepository.Register(userDomain)
}

func (uu *UserUseCase) Login(userDomain *Domain) string {
	user := uu.userRepository.Login(userDomain)

	if user.ID == 0 {
		return ""
	}
	token := uu.jwtAuth.GenerateToken(int(user.ID))

	return token
}

func (uu *UserUseCase) GetUser(id string) Domain {
	return uu.userRepository.GetUser(id)
}
