package service

import (
	"errors"

	"github.com/gesangwidigdo/auctify-be/dto"
	"github.com/gesangwidigdo/auctify-be/interfaces"
	"github.com/gesangwidigdo/auctify-be/model"
)

type userService struct {
	userRepository interfaces.UserRepository
}

func NewUserService(userRepository interfaces.UserRepository) interfaces.UserService {
	return &userService{
		userRepository,
	}
}

// Delete implements interfaces.UserService.
func (u *userService) Delete(id uint) (dto.UserDeleteResponse, error) {
	if err := u.userRepository.Delete(id); err != nil {
		return dto.UserDeleteResponse{}, err
	}
	return dto.UserDeleteResponse{
		ID: id,
	}, nil
}

// Detail implements interfaces.UserService.
func (u *userService) Detail(id uint) (dto.UserDetailResponse, error) {
	user, err := u.userRepository.Detail(id)
	if err != nil {
		return dto.UserDetailResponse{}, err
	}
	return dto.UserDetailResponse{
		ID:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		Username: user.Username,
		Address:  user.Address,
	}, nil
}

// List implements interfaces.UserService.
func (u *userService) List(role string) ([]dto.UserListResponse, error) {
	if role == "admin" {
		users, err := u.userRepository.List()
		if err != nil {
			return nil, err
		}
		var response []dto.UserListResponse
		for _, user := range users {
			response = append(response, dto.UserListResponse{
				ID:       user.ID,
				Name:     user.Name,
				Username: user.Username,
				Role:     user.Role,
			})
		}
		return response, nil
	}
	return nil, errors.New("unauthorized: user cannot access list of user")
}

// Update implements interfaces.UserService.
func (u *userService) Update(id uint, user dto.UserUpdateRequest) (dto.UserUpdateResponse, error) {
	request := model.User{
		Name:     user.Name,
		Email:    user.Email,
		Username: user.Username,
		Password: user.Password,
		Address:  user.Address,
	}
	updatedUser, err := u.userRepository.Update(id, request)
	if err != nil {
		return dto.UserUpdateResponse{}, err
	}
	return dto.UserUpdateResponse{
		ID:       updatedUser.ID,
		Name:     updatedUser.Name,
		Email:    updatedUser.Email,
		Username: updatedUser.Username,
		Address:  updatedUser.Address,
	}, nil
}
