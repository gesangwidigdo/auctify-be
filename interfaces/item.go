package interfaces

import (
	"github.com/gesangwidigdo/auctify-be/dto"
	"github.com/gesangwidigdo/auctify-be/model"
	"github.com/gin-gonic/gin"
)

type ItemRepository interface {
	Create(request model.Item) error
	List() ([]model.Item, error)
	Detail(id uint) (model.Item, error)
	Update(id uint, userID uint, request model.Item) error
	Delete(id uint, userID uint) error
}

type ItemService interface {
	Create(userID uint, request dto.ItemCreateRequest) error
	List() ([]dto.ItemListResponse, error)
	Detail(id uint) (dto.ItemDetailResponse, error)
	Update(id uint, userID uint, request dto.ItemUpdateRequest) error
	Delete(id uint, userID uint) error
}

type ItemController interface {
	Create(ctx *gin.Context)
	List(ctx *gin.Context)
	Detail(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}
