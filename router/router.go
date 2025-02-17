package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Router(r *gin.Engine, db *gorm.DB) {
	authRoutes := r.Group("/api/auth")
	AuthRoute(authRoutes, db)

	userRoutes := r.Group("/api/user")
	UserRoute(userRoutes, db)

	auctionRoutes := r.Group("/api/auction")
	AuctionRouter(auctionRoutes, db)

	offerRoutes := r.Group("/api/offer")
	OfferRouter(offerRoutes, db)
}
