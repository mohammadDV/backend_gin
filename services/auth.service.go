package services

import (
	"github.com/mohammadDV/backend_gin/entity"
	"github.com/mohammadDV/backend_gin/repository"
	"github.com/mohammadDV/backend_gin/serilizers"
)

type AuthService interface {
	AddUserService(registerRequest serilizers.RegisterRequest) (string, error)
}

type authService struct {
	authRepository repository.AuthRepository
}

func NewAuthRepository(authRepository repository.AuthRepository) AuthService {
	return &authService{
		authRepository: authRepository,
	}
}

func (a *authService) AddUserService(registerRequest serilizers.RegisterRequest) (string, error) {

	user := entity.User{}
	user.FirstName = registerRequest.FirstName
	user.LastName = registerRequest.LastName
	user.Email = registerRequest.Email
	user.Password = registerRequest.Password

	addUser ,err :=  a.authRepository.AddUser(user)
	if(err != nil){
		return "", err
	}

	return addUser,nil
}
