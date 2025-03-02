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
// @Summary Create Offer
// @Description Create offer data
// @Tags offer
// @Accept json
// @Produce json
// @Param request body dto.OfferCreateRequest true "Offer Create Request"
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Router /offer [post]
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
// @Summary Offer List By Auction
// @Description Get list offer data by auction id
// @Tags offer
// @Accept json
// @Produce json
// @Param id path string true "Auction ID"
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Router /offer/{id} [get]
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
