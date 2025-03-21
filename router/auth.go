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

func AuthRoute(r *gin.RouterGroup, db *gorm.DB) {
	var userRepo interfaces.UserRepository = repository.NewUserRepository(db)
	var authService interfaces.AuthService = service.NewAuthService(userRepo)
	var authController interfaces.AuthController = controller.NewAuthController(authService)

	r.POST("/register", authController.Register)
	r.POST("/login", authController.Login)
	r.POST("/logout", middleware.AuthMiddleware, authController.Logout)
}
