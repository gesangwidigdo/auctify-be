package service

import (
	"github.com/gesangwidigdo/auctify-be/dto"
	"github.com/gesangwidigdo/auctify-be/interfaces"
	"github.com/gesangwidigdo/auctify-be/model"
)

type authService struct {
	userRepository interfaces.UserRepository
}

func NewAuthService(userRepository interfaces.UserRepository) interfaces.AuthService {
	return &authService{
		userRepository,
	}
}

// Login implements interfaces.AuthService.
func (a *authService) Login(request dto.UserLoginRequest) (dto.UserLoginResponse, error) {
	panic("unimplemented")
}

// Register implements interfaces.AuthService.
func (a *authService) Register(request dto.UserRegisterRequest) (dto.UserRegisterResponse, error) {
	newUser := model.User{
		Name:     request.Name,
		Email:    request.Email,
		Username: request.Username,
		Password: request.Password,
		Role:     "user",
		Address:  request.Address,
	}
	user, err := a.userRepository.Register(newUser)
	if err != nil {
		return dto.UserRegisterResponse{}, err
	}
	return dto.UserRegisterResponse{
		ID:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		Username: user.Username,
		Address:  user.Address,
	}, nil
}
