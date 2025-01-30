package router

import (
	"github.com/gesangwidigdo/auctify-be/controller"
	"github.com/gesangwidigdo/auctify-be/interfaces"
	"github.com/gesangwidigdo/auctify-be/middleware"
	"github.com/gesangwidigdo/auctify-be/repository"
	"github.com/gesangwidigdo/auctify-be/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UserRoute(r *gin.RouterGroup, db *gorm.DB) {
	var userRepo interfaces.UserRepository = repository.NewUserRepository(db)
	var userService interfaces.UserService = service.NewUserService(userRepo)
	var userController interfaces.UserController = controller.NewUserController(userService)

	r.GET("/", userController.List)
	r.PUT("/", middleware.AuthMiddleware, userController.Update)
	r.GET("/me", middleware.AuthMiddleware, userController.Detail)
}
