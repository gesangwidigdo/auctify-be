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

func OfferRouter(r *gin.RouterGroup, db *gorm.DB) {
	var offerRepo interfaces.OfferRepository = repository.NewOfferRepository(db)
	var auctionRepo interfaces.AuctionRepository = repository.NewAuctionRepository(db)
	var offerService interfaces.OfferService = service.NewOfferService(offerRepo, auctionRepo)
	var offerController interfaces.OfferController = controller.NewOfferController(offerService)

	r.POST("/", middleware.AuthMiddleware, offerController.Create)
	r.GET("/:auction_id", offerController.List)
}
