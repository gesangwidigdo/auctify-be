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

// List implements interfaces.OfferRepository.
func (o *offerRepository) List(auctionId uint) ([]model.Offer, error) {
	var offers []model.Offer
	if err := o.db.Preload("User").Where("auction_id = ?", auctionId).Find(&offers).Error; err != nil {
		return nil, err
	}
	return offers, nil
}
