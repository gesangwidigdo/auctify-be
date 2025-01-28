package interfaces

import (
	"github.com/gesangwidigdo/auctify-be/dto"
	"github.com/gin-gonic/gin"
)

type AuthController interface {
	Register(ctx *gin.Context)
	Login(ctx *gin.Context)
	Logout(ctx *gin.Context)
}

type AuthService interface {
	Register(request dto.UserRegisterRequest) (dto.UserRegisterResponse, error)
	Login(request dto.UserLoginRequest) (dto.UserLoginResponse, error)
}
