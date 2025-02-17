package repository

import (
	"github.com/gesangwidigdo/auctify-be/interfaces"
	"github.com/gesangwidigdo/auctify-be/model"
	"gorm.io/gorm"
)

type offerRepository struct {
	db *gorm.DB
}

func NewOfferRepository(db *gorm.DB) interfaces.OfferRepository {
	return &offerRepository{db}
}

// Create implements interfaces.OfferRepository.
func (o *offerRepository) Create(request model.Offer) error {
	if err := o.db.Create(&request).Error; err != nil {
		return err
	}
	return nil
}
