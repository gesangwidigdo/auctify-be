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

func ItemRouter(r *gin.RouterGroup, db *gorm.DB) {
	var itemRepo interfaces.ItemRepository = repository.NewItemRepository(db)
	var itemService interfaces.ItemService = service.NewItemService(itemRepo)
	var itemController interfaces.ItemController = controller.NewItemController(itemService)

	r.POST("/", middleware.AuthMiddleware, itemController.Create)
	r.GET("/", itemController.List)
	r.GET("/:id", itemController.Detail)
	r.PUT("/:id", middleware.AuthMiddleware, itemController.Update)
	r.DELETE("/:id", middleware.AuthMiddleware, itemController.Delete)
}
