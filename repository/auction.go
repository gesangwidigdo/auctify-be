package repository

import (
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
	if err := a.db.Create(&request).Error; err != nil {
		return err
	}
	return nil
}

// Detail implements interfaces.AuctionRepository.
func (a *auctionRepository) Detail(id uint) (model.Auction, error) {
	var auction model.Auction
	if err := a.db.Preload("Item.User").Preload("Item").Where("id = ?", id).First(&auction).Error; err != nil {
		return model.Auction{}, err
	}
	return auction, nil
}

// List implements interfaces.AuctionRepository.
func (a *auctionRepository) List() ([]model.Auction, error) {
	var auctions []model.Auction
	if err := a.db.Preload("Item.User").Preload("Item").Find(&auctions).Error; err != nil {
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

// Menutup auction tertentu
func (a *auctionRepository) CloseAuction(id uint) error {
	return a.db.Exec(
		"UPDATE auctions SET is_closed = ? WHERE id = ? AND end_time <= NOW() AND is_closed = false AND deleted_at IS NULL",
		true, id,
	).Error
}

// Ambil semua auction yang sudah harus ditutup
func (a *auctionRepository) GetAuctionsToClose() ([]model.Auction, error) {
	var auctions []model.Auction
	err := a.db.Raw(
		"SELECT id, end_time FROM auctions WHERE end_time <= NOW() AND is_closed = false AND deleted_at IS NULL",
	).Scan(&auctions).Error
	return auctions, err
}
