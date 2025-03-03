package repository

import (
	"errors"

	"github.com/gesangwidigdo/auctify-be/interfaces"
	"github.com/gesangwidigdo/auctify-be/model"
	"gorm.io/gorm"
)

type itemRepository struct {
	db *gorm.DB
}

func NewItemRepository(db *gorm.DB) interfaces.ItemRepository {
	return &itemRepository{
		db,
	}
}

// Create implements interfaces.ItemRepository.
func (i itemRepository) Create(request model.Item) error {
	if err := i.db.Create(&request).Error; err != nil {
		return err
	}
	return nil
}

// Delete implements interfaces.ItemRepository.
func (i itemRepository) Delete(id uint, userID uint) error {
	if err := i.db.Where("user_id = ?", userID).Delete(&model.Item{}, id); err != nil {
		if err.Error != nil {
			return err.Error
		}

		if err.RowsAffected == 0 {
			return errors.New("no item found")
		}
	}
	return nil
}

// Detail implements interfaces.ItemRepository.
func (i itemRepository) Detail(id uint) (model.Item, error) {
	var item model.Item
	if err := i.db.Preload("User").First(&item, id).Error; err != nil {
		return model.Item{}, err
	}
	return item, nil
}

// List implements interfaces.ItemRepository.
func (i itemRepository) List() ([]model.Item, error) {
	var items []model.Item
	if err := i.db.Preload("User").Find(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}

// Update implements interfaces.ItemRepository.
func (i itemRepository) Update(id uint, userID uint, request model.Item) error {
	if err := i.db.Model(model.Item{}).Where("id = ? AND user_id = ?", id, userID).Updates(&request); err != nil {
		if err.Error != nil {
			return err.Error
		}

		if err.RowsAffected == 0 {
			return errors.New("no item found")
		}
	}
	return nil
}
