package interfaces

import (
	"github.com/gesangwidigdo/auctify-be/dto"
	"github.com/gesangwidigdo/auctify-be/model"
	"github.com/gin-gonic/gin"
)

type OfferRepository interface {
	Create(request model.Offer) (error)
	List(auctionId uint) ([]model.Offer, error)
}
type OfferService interface {
	Create(userId uint, offer dto.OfferCreateRequest) (dto.OfferCreateResponse, error)
	List(auctionId uint) ([]dto.OfferListResponse, error)
}
type OfferController interface {
	Create(ctx *gin.Context)
	List(ctx *gin.Context)
}
