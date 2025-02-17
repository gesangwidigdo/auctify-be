package interfaces

import (
	"github.com/gesangwidigdo/auctify-be/dto"
	"github.com/gesangwidigdo/auctify-be/model"
	"github.com/gin-gonic/gin"
)

type OfferRepository interface {
	Create(request model.Offer) dto.OfferCreateResponse
}
type OfferService interface {
	Create(offer dto.OfferCreateRequest) dto.OfferCreateResponse
}
type OfferController interface {
	Create(ctx *gin.Context)
}
