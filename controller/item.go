package controller

import (
	"strconv"

	"github.com/gesangwidigdo/auctify-be/dto"
	"github.com/gesangwidigdo/auctify-be/interfaces"
	"github.com/gesangwidigdo/auctify-be/utils"
	"github.com/gin-gonic/gin"
)

type itemController struct {
	itemService interfaces.ItemService
}

// Create implements interfaces.ItemController.
// @Summary Create Item
// @Description Create new item
// @Tags item
// @Accept json
// @Produce json
// @Param request body dto.ItemCreateRequest true "Item Create Data"
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Router /item [post]
func (i itemController) Create(ctx *gin.Context) {
	id, statCode, err := utils.ExtractID(ctx)
	if err != nil {
		utils.FailResponse(ctx, statCode, err.Error())
	}

	var request dto.ItemCreateRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		utils.FailResponse(ctx, 400, err.Error())
		return
	}
	if err := i.itemService.Create(id, request); err != nil {
		utils.FailResponse(ctx, 400, err.Error())
		return
	}
	utils.SuccessResponse(ctx, 201, "success create item")
}

// Delete implements interfaces.ItemController.
// @Summary Delete Item
// @Description Delete item data
// @Tags item
// @Accept json
// @Produce json
// @Param id path string true "Item ID"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Router /item/{id} [delete]
func (i itemController) Delete(ctx *gin.Context) {
	userID, statCode, err := utils.ExtractID(ctx)
	if err != nil {
		utils.FailResponse(ctx, statCode, err.Error())
	}

	paramId := ctx.Param("id")
	itemID, err := strconv.ParseUint(paramId, 10, 64)
	if err != nil {
		utils.FailResponse(ctx, 400, "invalid id format")
		return
	}

	//
	if err := i.itemService.Delete(uint(itemID), userID); err != nil {
		utils.FailResponse(ctx, 400, err.Error())
		return
	}

	utils.SuccessResponse(ctx, 200, "success delete item")
}

// List implements interfaces.ItemController.
// @Summary List Item
// @Description Get item list
// @Tags item
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Router /item [get]
func (i itemController) List(ctx *gin.Context) {
	items, err := i.itemService.List()
	if err != nil {
		utils.FailResponse(ctx, 400, err.Error())
		return
	}

	utils.SuccessResponse(ctx, 200, items)
}

// Detail implements interfaces.ItemController.
// @Summary Get Item Detail
// @Description get item detail data
// @Tags item
// @Accept json
// @Produce json
// @Param id path string true "Item ID"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Router /item/{id} [get]
func (i itemController) Detail(ctx *gin.Context) {
	paramId := ctx.Param("id")
	id, err := strconv.ParseUint(paramId, 10, 64)
	if err != nil {
		utils.FailResponse(ctx, 400, "invalid id format")
		return
	}

	item, err := i.itemService.Detail(uint(id))
	if err != nil {
		utils.FailResponse(ctx, 400, "invalid id format")
		return
	}

	utils.SuccessResponse(ctx, 200, item)
}

// Update implements interfaces.ItemController.
// @Summary Update Item
// @Description Update item data
// @Tags item
// @Accept json
// @Produce json
// @Param id path string true "Item Data"
// @Param request body dto.ItemUpdateRequest true "Item Update Data"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Router /item/{id} [put]
func (i itemController) Update(ctx *gin.Context) {
	userID, statCode, err := utils.ExtractID(ctx)
	if err != nil {
		utils.FailResponse(ctx, statCode, err.Error())
	}

	paramId := ctx.Param("id")
	itemID, err := strconv.ParseUint(paramId, 10, 64)
	if err != nil {
		utils.FailResponse(ctx, 400, "invalid id format")
		return
	}

	var request dto.ItemUpdateRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		utils.FailResponse(ctx, 400, err.Error())
		return
	}

	if err := i.itemService.Update(uint(itemID), userID, request); err != nil {
		utils.FailResponse(ctx, 400, err.Error())
		return
	}

	utils.SuccessResponse(ctx, 200, "success update item")
}

func NewItemController(itemService interfaces.ItemService) interfaces.ItemController {
	return itemController{
		itemService,
	}
}
