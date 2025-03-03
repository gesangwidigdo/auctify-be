package controller

import (
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
// @Summary Delete user
// @Description Delete user data
// @Tags user
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Router /user [delete]
func (u *userController) Delete(ctx *gin.Context) {
	id, statCode, err := utils.ExtractID(ctx)
	if err != nil && statCode != 0 {
		utils.FailResponse(ctx, statCode, err.Error())
		return
	}
	deleteRes, err := u.userService.Delete(id)
	if err != nil {
		utils.FailResponse(ctx, 400, err.Error())
		return
	}

	ctx.SetCookie("token", "", -1, "/", "localhost", false, true)

	utils.SuccessResponse(ctx, 200, deleteRes)
}

// Detail implements interfaces.UserController.
// @Summary Get logged in user
// @Description Get user detail data
// @Tags user
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Router /user/me [get]
func (u *userController) Detail(ctx *gin.Context) {
	id, statCode, err := utils.ExtractID(ctx)
	if err != nil && statCode != 0 {
		utils.FailResponse(ctx, statCode, err.Error())
		return
	}
	user, err := u.userService.Detail(id)
	if err != nil {
		utils.FailResponse(ctx, 400, err.Error())
		return
	}

	utils.SuccessResponse(ctx, 200, user)
}

// List implements interfaces.UserController.
// @Summary List user
// @Description Get list of user data
// @Tags user
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Router /user [get]
func (u *userController) List(ctx *gin.Context) {
	role, ok := ctx.Get("role")
	if !ok {
		utils.FailResponse(ctx, 400, "no role found")
		return
	}

	roleStr, ok := role.(string)
	if !ok {
		utils.FailResponse(ctx, 400, "role is not a string")
		return
	}

	users, err := u.userService.List(roleStr)
	if err != nil {
		utils.FailResponse(ctx, 400, err.Error())
		return
	}

	utils.SuccessResponse(ctx, 200, users)
}

// Update implements interfaces.UserController.
// @Summary Update user
// @Description Update user data
// @Tags user
// @Accept json
// @Produce json
// @Param request body dto.UserUpdateRequest true "User Update Data"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Router /user [put]
func (u *userController) Update(ctx *gin.Context) {
	var userUpdateRequest dto.UserUpdateRequest
	if err := ctx.ShouldBindJSON(&userUpdateRequest); err != nil {
		utils.FailResponse(ctx, 400, err.Error())
		return
	}

	id, statCode, err := utils.ExtractID(ctx)
	if err != nil && statCode != 0 {
		utils.FailResponse(ctx, statCode, err.Error())
		return
	}
	updateRes, err := u.userService.Update(id, userUpdateRequest)
	if err != nil {
		utils.FailResponse(ctx, 400, err.Error())
		return
	}

	utils.SuccessResponse(ctx, 200, updateRes)
}
