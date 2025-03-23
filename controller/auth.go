package controller

import (
	"net/http"

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

// Register mendaftarkan user baru
// @Summary Register user baru
// @Description Mendaftarkan user baru dengan email dan password
// @Tags auth
// @Accept json
// @Produce json
// @Param request body dto.UserRegisterRequest true "User Registration Data"
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /auth/register [post]
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

// Login user ke sistem
// @Summary Login user
// @Description Login user dengan email dan password
// @Tags auth
// @Accept json
// @Produce json
// @Param request body dto.UserLoginRequest true "User Login Data"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Router /auth/login [post]
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

	ctx.SetSameSite(http.SameSiteNoneMode)
	ctx.SetCookie("token", loginResponse.Token, 3600*24, "/", "localhost", false, true)

	utils.SuccessResponse(ctx, 200, "login success")
}

// Logout user dari sistem
// @Summary Logout user
// @Description Menghapus token autentikasi
// @Tags auth
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Router /auth/logout [post]
func (a *authController) Logout(ctx *gin.Context) {
	_, err := ctx.Cookie("token")
	if err != nil {
		utils.FailResponse(ctx, 400, "no token found")
		return
	}
	ctx.SetCookie("token", "", -1, "/", "localhost", false, true)
	utils.SuccessResponse(ctx, 200, "logout success")
}
