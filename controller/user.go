package controller

import (
	"fmt"

	"github.com/gesangwidigdo/auctify-be/dto"
	"github.com/gesangwidigdo/auctify-be/interfaces"
	"github.com/gesangwidigdo/auctify-be/utils"
	"github.com/gin-gonic/gin"
)

type userController struct {
	userService interfaces.UserService
}

func NewUserController(userService interfaces.UserService) interfaces.UserController {
	return &userController{
		userService,
	}
}

// Delete implements interfaces.UserController.
func (u *userController) Delete(ctx *gin.Context) {
	panic("unimplemented")
}

// Detail implements interfaces.UserController.
func (u *userController) Detail(ctx *gin.Context) {
	panic("unimplemented")
}

// List implements interfaces.UserController.
func (u *userController) List(ctx *gin.Context) {
	users, err := u.userService.List()
	if err != nil {
		utils.FailResponse(ctx, 400, err.Error())
	}

	utils.SuccessResponse(ctx, 200, users)
}

// Update implements interfaces.UserController.
func (u *userController) Update(ctx *gin.Context) {
	var userUpdateRequest dto.UserUpdateRequest
	id, ok := ctx.Get("id")
	if !ok {
		utils.FailResponse(ctx, 400, "User not found")
		return
	}

	if err := ctx.ShouldBindJSON(&userUpdateRequest); err != nil {
		utils.FailResponse(ctx, 400, err.Error())
		return
	}

	fmt.Println("goyim1")
	idFloat, ok := id.(float64)
	if !ok {
		utils.FailResponse(ctx, 400, "Invalid user ID type")
		return
	}
	idUint := uint(idFloat)
	updateRes, err := u.userService.Update(idUint, userUpdateRequest)
	if err != nil {
		utils.FailResponse(ctx, 400, err.Error())
		return
	}

	utils.SuccessResponse(ctx, 200, updateRes)
}
