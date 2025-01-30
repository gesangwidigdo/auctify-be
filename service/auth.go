package service

import (
	"errors"
	"strings"

	"github.com/gesangwidigdo/auctify-be/dto"
	"github.com/gesangwidigdo/auctify-be/interfaces"
	"github.com/gesangwidigdo/auctify-be/model"
	"github.com/gesangwidigdo/auctify-be/utils"
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
	userFound, err := a.userRepository.GetByUsername(request.Username)
	if err != nil {
		return dto.UserLoginResponse{}, errors.New("Username does not exist, please register first")
	}

	println(userFound.Password)
	println(request.Password)

	if !utils.CheckPasswordHash(userFound.Password, request.Password) {
		return dto.UserLoginResponse{}, errors.New("Invalid password")
	}

	tokenString, err := utils.CreateToken(userFound)
	if err != nil {
		return dto.UserLoginResponse{}, err
	}
	return dto.UserLoginResponse{
		Username: userFound.Username,
		Token:    tokenString,
	}, nil
}

// Register implements interfaces.AuthService.
func (a *authService) Register(request dto.UserRegisterRequest) (dto.UserRegisterResponse, error) {
	hashedPassword, err := utils.HashPassword(request.Password)
	if err != nil {
		return dto.UserRegisterResponse{}, err
	}
	newUser := model.User{
		Name:     request.Name,
		Email:    request.Email,
		Username: request.Username,
		Password: hashedPassword,
		Role:     "user",
		Address:  request.Address,
	}
	user, err := a.userRepository.Register(newUser)
	if err != nil {
		if strings.Contains(strings.ToLower(err.Error()), "duplicate entry") {
			if strings.Contains(err.Error(), "email") {
				return dto.UserRegisterResponse{}, errors.New("email already registered")
			} else if strings.Contains(err.Error(), "uni_users_username") {
				return dto.UserRegisterResponse{}, errors.New("username already registered")
			}
		}
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
