package controller

import (
	"github.com/gesangwidigdo/auctify-be/dto"
	"github.com/gesangwidigdo/auctify-be/interfaces"
	"github.com/gesangwidigdo/auctify-be/utils"
	"github.com/gin-gonic/gin"
)

type authController struct {
	authService interfaces.AuthService
}

func NewAuthController(authService interfaces.AuthService) interfaces.AuthController {
	return &authController{
		authService,
	}
}

// Login implements interfaces.AuthController.
func (a *authController) Login(ctx *gin.Context) {
	panic("unimplemented")
}

// Register implements interfaces.AuthController.
func (a *authController) Register(ctx *gin.Context) {
	var userRequest dto.UserRegisterRequest
	if err := ctx.ShouldBindJSON(&userRequest); err != nil {
		utils.FailResponse(ctx, 400, err.Error())
		return
	}

	userResponse, err := a.authService.Register(userRequest)
	if err != nil {
		utils.FailResponse(ctx, 500, err.Error())
		return
	}

	utils.SuccessResponse(ctx, 201, userResponse)
}
