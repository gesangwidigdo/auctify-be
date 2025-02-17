package service

import (
	"errors"
	"time"

	"github.com/gesangwidigdo/auctify-be/dto"
	"github.com/gesangwidigdo/auctify-be/interfaces"
	"github.com/gesangwidigdo/auctify-be/model"
)

type auctionService struct {
	auctionRepo interfaces.AuctionRepository
}

func NewAuctionService(auctionRepo interfaces.AuctionRepository) interfaces.AuctionService {
	return &auctionService{
		auctionRepo,
	}
}

// Create implements interfaces.AuctionService.
func (a *auctionService) Create(id uint, request dto.AuctionCreateRequest) (dto.AuctionCreateResponse, error) {
	currentTime := time.Now()
	if request.EndTime.Before(currentTime) {
		return dto.AuctionCreateResponse{}, errors.New("insert valid end time")
	}

	if request.EndTime.Before(currentTime.AddDate(0, 0, 3)) {
		return dto.AuctionCreateResponse{}, errors.New("end time must be at least 3 day from now")
	}

	if request.StartPrice <= 0 {
		return dto.AuctionCreateResponse{}, errors.New("start price must be greater than 0")
	}

	if request.ItemName == "" {
		return dto.AuctionCreateResponse{}, errors.New("item name must not be empty")
	}

	newAuction := model.Auction{
		ItemName:     request.ItemName,
		Description:  request.Description,
		StartTime:    currentTime,
		EndTime:      request.EndTime,
		StartPrice:   request.StartPrice,
		CurrentPrice: request.StartPrice,
		IsClosed:     false,
		UserID:       id,
	}

	if err := a.auctionRepo.Create(newAuction); err != nil {
		return dto.AuctionCreateResponse{}, err
	}

	return dto.AuctionCreateResponse{
		ItemName:     newAuction.ItemName,
		Description:  newAuction.Description,
		EndTime:      newAuction.EndTime,
		StartPrice:   newAuction.StartPrice,
		CurrentPrice: newAuction.CurrentPrice,
	}, nil
}

// Detail implements interfaces.AuctionService.
func (a *auctionService) Detail(id uint) (dto.AuctionDetailResponse, error) {
	panic("unimplemented")
}

// List implements interfaces.AuctionService.
func (a *auctionService) List() ([]dto.AuctionListResponse, error) {
	panic("unimplemented")
}

// Update implements interfaces.AuctionService.
func (a *auctionService) Update(id uint, request dto.AuctionUpdateRequest) (dto.AuctionUpdateResponse, error) {
	panic("unimplemented")
}

// UpdateCurrentPrice implements interfaces.AuctionService.
func (a *auctionService) UpdateCurrentPrice(id uint, request dto.AuctionUpdateCurrentPriceRequest) (dto.AuctionUpdateCurrentPriceResponse, error) {
	panic("unimplemented")
}

// CloseAuction implements interfaces.AuctionService.
func (a *auctionService) CloseAuction(id uint) (dto.AuctionCloseResponse, error) {
	panic("unimplemented")
}

// Delete implements interfaces.AuctionService.
func (a *auctionService) Delete(id uint) error {
	panic("unimplemented")
}
