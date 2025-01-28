package router

// import (
// 	"github.com/gesangwidigdo/auctify-be/controller"
// 	"github.com/gesangwidigdo/auctify-be/interfaces"
// 	"github.com/gesangwidigdo/auctify-be/repository"
// 	"github.com/gesangwidigdo/auctify-be/service"
// 	"github.com/gin-gonic/gin"
// 	"gorm.io/gorm"
// )

// func UserRoute(r *gin.RouterGroup, db *gorm.DB) {
// 	var userRepo interfaces.UserRepository = repository.NewUserRepository(db)
// 	var userService interfaces.AuthService = service.NewAuthService(userRepo)
// 	var userController interfaces.AuthController = controller.NewAuthController(authService)

// 	r.POST("/register", userController.Register)
// 	// r.POST("/login", nil)
// }
