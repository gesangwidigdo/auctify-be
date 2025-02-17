package controller

import (
	"strconv"

	"github.com/gesangwidigdo/auctify-be/dto"
	"github.com/gesangwidigdo/auctify-be/interfaces"
	"github.com/gesangwidigdo/auctify-be/utils"
	"github.com/gin-gonic/gin"
)

type offerController struct {
	offerService interfaces.OfferService
}

func NewOfferController(offerService interfaces.OfferService) interfaces.OfferController {
	return &offerController{offerService}
}

// Create implements interfaces.OfferController.
func (o *offerController) Create(ctx *gin.Context) {
	id, statCode, err := utils.ExtractID(ctx)
	if err != nil {
		utils.FailResponse(ctx, statCode, err.Error())
		return
	}

	var offerRequest dto.OfferCreateRequest
	if err := ctx.ShouldBindJSON(&offerRequest); err != nil {
		utils.FailResponse(ctx, 400, err.Error())
		return
	}

	offerResponse, err := o.offerService.Create(id, offerRequest)
	if err != nil {
		utils.FailResponse(ctx, 400, err.Error())
		return
	}

	utils.SuccessResponse(ctx, 201, offerResponse)
}

// List implements interfaces.OfferController.
func (o *offerController) List(ctx *gin.Context) {
	auctionID := ctx.Param("auction_id")
	uintAuctionID, err := strconv.ParseUint(auctionID, 10, 64)
	if err != nil {
		utils.FailResponse(ctx, 400, err.Error())
		return
	}

	offers, err := o.offerService.List(uint(uintAuctionID))
	if err != nil {
		utils.FailResponse(ctx, 400, err.Error())
		return
	}

	utils.SuccessResponse(ctx, 200, offers)
}
