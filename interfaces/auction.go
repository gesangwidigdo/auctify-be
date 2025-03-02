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
	GetAuctionsToClose() ([]model.Auction, error)
}

type AuctionService interface {
	Create(id uint, request dto.AuctionCreateRequest) error
	List() ([]dto.AuctionListResponse, error)
	Detail(id uint) (dto.AuctionDetailResponse, error)
	StartAuctionAutoClose()
}

type AuctionController interface {
	Create(ctx *gin.Context)
	List(ctx *gin.Context)
	Detail(ctx *gin.Context)
}
