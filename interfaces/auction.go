package interfaces

import (
	"github.com/gesangwidigdo/auctify-be/dto"
	"github.com/gesangwidigdo/auctify-be/model"
	"github.com/gin-gonic/gin"
)

type AuctionRepository interface {
	Create(request model.Auction) error
	Update(id uint, request model.Auction) error
	UpdateCurrentPrice(id uint, price float64) error
	CloseAuction(id uint) error
	List() ([]model.Auction, error)
	Detail(id uint) (model.Auction, error)
}

type AuctionService interface {
	Create(id uint, request dto.AuctionCreateRequest) (dto.AuctionCreateResponse, error)
	List() ([]dto.AuctionListResponse, error)
	Detail(id uint) (dto.AuctionDetailResponse, error)
	Update(id uint, request dto.AuctionUpdateRequest) (error)
	UpdateCurrentPrice(id uint, request dto.AuctionUpdateCurrentPriceRequest) (dto.AuctionUpdateCurrentPriceResponse, error)
	CloseAuction(id uint) (dto.AuctionCloseResponse, error)
}

type AuctionController interface {
	Create(ctx *gin.Context)
	List(ctx *gin.Context)
	Detail(ctx *gin.Context)
	Update(ctx *gin.Context)
	UpdateCurrentPrice(ctx *gin.Context)
	CloseAuction(ctx *gin.Context)
}
