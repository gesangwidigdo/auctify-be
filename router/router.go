package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router(r *gin.Engine, db *gorm.DB) {
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	authRoutes := r.Group("/api/auth")
	AuthRoute(authRoutes, db)

	userRoutes := r.Group("/api/user")
	UserRoute(userRoutes, db)

	auctionRoutes := r.Group("/api/auction")
	AuctionRouter(auctionRoutes, db)

	offerRoutes := r.Group("/api/offer")
	OfferRouter(offerRoutes, db)
}
