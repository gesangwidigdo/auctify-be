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

// Login implements interfaces.AuthController.
func (a *authController) Login(ctx *gin.Context) {
	var loginAttempt dto.UserLoginRequest
	if err := ctx.ShouldBindJSON(&loginAttempt); err != nil {
		utils.FailResponse(ctx, 400, err.Error())
		return
	}

	loginResponse, err := a.authService.Login(loginAttempt)
	if err != nil {
		utils.FailResponse(ctx, 401, err.Error())
		return
	}

	ctx.SetCookie("token", loginResponse.Token, 3600*24, "/", "localhost", false, true)

	utils.SuccessResponse(ctx, 200, "login success")
}

// Logout implements interfaces.AuthController.
func (a *authController) Logout(ctx *gin.Context) {
	_, err := ctx.Cookie("token")
	if err != nil {
		utils.FailResponse(ctx, 400, "no token found")
		return
	}
	ctx.SetCookie("token", "", -1, "/", "localhost", false, true)
	utils.SuccessResponse(ctx, 200, "logout success")
}