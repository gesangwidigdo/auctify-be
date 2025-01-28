package interfaces

import (
	"github.com/gesangwidigdo/auctify-be/dto"
	"github.com/gesangwidigdo/auctify-be/model"
	"github.com/gin-gonic/gin"
)

type UserRepository interface {
	Register(request model.User) (model.User, error)
	GetByUsername(username string) (model.User, error)
	Update(id uint, request model.User) (model.User, error)
	Delete(id uint) (error)
	List() ([]model.User, error)
	Detail(id uint) (model.User, error)
}

type UserService interface {
	Update(id uint, user dto.UserUpdateRequest) (dto.UserUpdateResponse, error)
	Delete(id uint) (dto.UserDeleteResponse, error)
	List() ([]dto.UserListResponse, error)
	Detail(id uint) (dto.UserDetailResponse, error)
}

type UserController interface {
	Update(*gin.Context) 
	Delete(*gin.Context)
	List(*gin.Context)
	Detail(*gin.Context)
}