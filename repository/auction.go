package repository

import (
	"fmt"

	"github.com/gesangwidigdo/auctify-be/interfaces"
	"github.com/gesangwidigdo/auctify-be/model"
	"gorm.io/gorm"
)

type auctionRepository struct {
	db *gorm.DB
}

func NewAuctionRepository(db *gorm.DB) interfaces.AuctionRepository {
	return &auctionRepository{
		db,
	}
}

// Create implements interfaces.AuctionRepository.
func (a *auctionRepository) Create(request model.Auction) error {
	fmt.Println(request.EndTime)
	if err := a.db.Exec(
		"INSERT INTO auctions (item_name, description, start_time, end_time, start_price, current_price, is_closed, user_id, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?, NOW(), NOW())",
		request.ItemName,
		request.Description,
		request.StartTime,
		request.EndTime,
		request.StartPrice,
		request.CurrentPrice,
		request.IsClosed,
		request.UserID,
	).Error; err != nil {
		return err
	}
	return nil
}

// Detail implements interfaces.AuctionRepository.
func (a *auctionRepository) Detail(id uint) (model.Auction, error) {
	var auction model.Auction
	if err := a.db.Raw(
		"SELECT id, item_name, description, start_time, end_time, start_price, current_price, is_closed, user_id FROM auctions WHERE id = ? AND deleted_at IS NULL",
		id,
	).Take(&auction).Error; err != nil {
		return model.Auction{}, err
	}
	return auction, nil
}

// List implements interfaces.AuctionRepository.
func (a *auctionRepository) List() ([]model.Auction, error) {
	var auctions []model.Auction
	if err := a.db.Raw(
		"SELECT item_name, description, start_time, end_time, current_price, is_closed FROM auctions WHERE deleted_at IS NULL",
	).Scan(&auctions).Error; err != nil {
		return nil, err
	}
	return auctions, nil
}

// Update implements interfaces.AuctionRepository.
func (a *auctionRepository) Update(id uint, request model.Auction) error {
	if err := a.db.Model(&model.Auction{}).Where("id = ?", id).Updates(request).Error; err != nil {
		return err
	}
	return nil
}

func (a *auctionRepository) UpdateCurrentPrice(id uint, price float64) error {
	if err := a.db.Exec(
		"UPDATE auctions SET current_price = ? WHERE id = ? AND deleted_at IS NULL",
		price,
		id,
	).Error; err != nil {
		return err
	}
	return nil
}

func (a *auctionRepository) CloseAuction(id uint) error {
	if err := a.db.Exec(
		"UPDATE auctions SET is_closed = ? WHERE id = ? AND deleted_at IS NULL",
		true,
		id,
	).Error; err != nil {
		return err
	}
	return nil
}