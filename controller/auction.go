package controller

import (
	"github.com/gesangwidigdo/auctify-be/dto"
	"github.com/gesangwidigdo/auctify-be/interfaces"
	"github.com/gesangwidigdo/auctify-be/utils"
	"github.com/gin-gonic/gin"
)

type auctionController struct {
	auctionService interfaces.AuctionService
}

func NewAuctionController(auctionService interfaces.AuctionService) interfaces.AuctionController {
	return &auctionController{
		auctionService,
	}
}

// Create implements interfaces.AuctionController.
func (a *auctionController) Create(ctx *gin.Context) {
	var auctionRequest dto.AuctionCreateRequest
	if err := ctx.ShouldBindJSON(&auctionRequest); err != nil {
		utils.FailResponse(ctx, 400, err.Error())
		return
	}
	id, statCode, err := utils.ExtractID(ctx)
	if err != nil {
		utils.FailResponse(ctx, statCode, err.Error())
		return
	}
	auctionResponse, err := a.auctionService.Create(id, auctionRequest)
	if err != nil {
		utils.FailResponse(ctx, 400, err.Error())
		return
	}
	utils.SuccessResponse(ctx, 201, auctionResponse)
}

// Detail implements interfaces.AuctionController.
func (a *auctionController) Detail(ctx *gin.Context) {
	panic("unimplemented")
}

// List implements interfaces.AuctionController.
func (a *auctionController) List(ctx *gin.Context) {
	panic("unimplemented")
}

// Update implements interfaces.AuctionController.
func (a *auctionController) Update(ctx *gin.Context) {
	panic("unimplemented")
}

// UpdateCurrentPrice implements interfaces.AuctionController.
func (a *auctionController) UpdateCurrentPrice(ctx *gin.Context) {
	panic("unimplemented")
}

// CloseAuction implements interfaces.AuctionController.
func (a *auctionController) CloseAuction(ctx *gin.Context) {
	panic("unimplemented")
}

// Delete implements interfaces.AuctionController.
func (a *auctionController) Delete(ctx *gin.Context) {
	panic("unimplemented")
}
