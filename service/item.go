package service

import (
	"github.com/gesangwidigdo/auctify-be/dto"
	"github.com/gesangwidigdo/auctify-be/interfaces"
	"github.com/gesangwidigdo/auctify-be/model"
)

type itemService struct {
	itemRepository interfaces.ItemRepository
}

// Create implements interfaces.ItemService.
func (i *itemService) Create(userID uint, request dto.ItemCreateRequest) error {
	newItem := model.Item{
		ItemName: request.ItemName,
		Description: request.Description,
		UserID: userID,
	}

	if err := i.itemRepository.Create(newItem); err != nil {
		return err
	}
	return nil
}

// Delete implements interfaces.ItemService.
func (i *itemService) Delete(itemID uint, userID uint) error {
	if err := i.itemRepository.Delete(itemID, userID); err != nil {
		return err
	}
	return nil
}

// Detail implements interfaces.ItemService.
func (i *itemService) Detail(id uint) (dto.ItemDetailResponse, error) {
	item, err := i.itemRepository.Detail(id)
	if err != nil {
		return dto.ItemDetailResponse{}, err
	}

	return dto.ItemDetailResponse{
		User: dto.UserItem{
			Username: item.User.Username,
		},
		ItemName: item.ItemName,
		Description: item.Description,
	}, nil
}

// List implements interfaces.ItemService.
func (i *itemService) List() ([]dto.ItemListResponse, error) {
	items, err := i.itemRepository.List()
	if err != nil {
		return nil, err
	}

	var response []dto.ItemListResponse
	for _, item := range items {
		response = append(response, dto.ItemListResponse{
			User: dto.UserItem{
				Username: item.User.Username,
			},
			ItemName: item.ItemName,
		})
	}

	return response, nil
}

// Update implements interfaces.ItemService.
func (i *itemService) Update(itemID uint, userID uint, request dto.ItemUpdateRequest) error {
	newItemData := model.Item{
		ItemName: request.ItemName,
		Description: request.Description,
	}
	if err := i.itemRepository.Update(itemID, userID, newItemData); err != nil {
		return err
	}
	return nil
}

func NewItemService(itemRepository interfaces.ItemRepository) interfaces.ItemService {
	return &itemService{
		itemRepository,
	}
}
