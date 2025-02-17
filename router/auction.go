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

func AuctionRouter(r *gin.RouterGroup, db *gorm.DB) {
	var auctionRepo interfaces.AuctionRepository = repository.NewAuctionRepository(db)
	var auctionService interfaces.AuctionService = service.NewAuctionService(auctionRepo)
	var auctionController interfaces.AuctionController = controller.NewAuctionController(auctionService)

	r.POST("/", middleware.AuthMiddleware, auctionController.Create)
	r.GET("/", auctionController.List)
	r.GET("/:id", auctionController.Detail)
	r.PUT("/:id", middleware.AuthMiddleware, auctionController.Update)
}
