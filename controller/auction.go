package controller

import (
	"strconv"

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
// @Summary Create Auction
// @Description Create new auction
// @Tags auction
// @Accept json
// @Produce json
// @Param request body dto.AuctionCreateRequest true "Auction Create Data"
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Router /auction [post]
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
	if err := a.auctionService.Create(id, auctionRequest); err != nil {
		utils.FailResponse(ctx, 400, err.Error())
		return
	}
	utils.SuccessResponse(ctx, 201, "success create auction")
}

// Detail implements interfaces.AuctionController.
// @Summary Auction Detail
// @Description Get auction detail
// @Tags auction
// @Accept json
// @Produce json
// @Param id path string true "Auction ID"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Router /auction/{id} [get]
func (a *auctionController) Detail(ctx *gin.Context) {
	id := ctx.Param("id")
	auctionID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		utils.FailResponse(ctx, 400, err.Error())
		return
	}

	auctionResponse, err := a.auctionService.Detail(uint(auctionID))
	if err != nil {
		utils.FailResponse(ctx, 400, err.Error())
		return
	}

	utils.SuccessResponse(ctx, 200, auctionResponse)
}

// List implements interfaces.AuctionController.
// @Summary Auction List
// @Description Get auction list
// @Tags auction
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Router /auction/ [get]
func (a *auctionController) List(ctx *gin.Context) {
	auctions, err := a.auctionService.List()
	if err != nil {
		utils.FailResponse(ctx, 400, err.Error())
		return
	}

	utils.SuccessResponse(ctx, 200, auctions)
}
