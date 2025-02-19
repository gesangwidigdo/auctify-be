package service

import (
	"errors"
	"fmt"
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
	if request.EndTime.UTC().Before(currentTime) {
		return dto.AuctionCreateResponse{}, errors.New("insert valid end time")
	}

	// if request.EndTime.Before(currentTime.AddDate(0, 0, 3)) {
	// 	return dto.AuctionCreateResponse{}, errors.New("end time must be at least 3 day from now")
	// }

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
	auction, err := a.auctionRepo.Detail(id)
	if err != nil {
		return dto.AuctionDetailResponse{}, err
	}

	return dto.AuctionDetailResponse{
		ID:           auction.ID,
		ItemName:     auction.ItemName,
		Description:  auction.Description,
		StartTime:    auction.StartTime.String(),
		EndTime:      auction.EndTime.String(),
		StartPrice:   auction.StartPrice,
		CurrentPrice: auction.CurrentPrice,
		IsClosed:     auction.IsClosed,
	}, nil
}

// List implements interfaces.AuctionService.
func (a *auctionService) List() ([]dto.AuctionListResponse, error) {
	auctions, err := a.auctionRepo.List()
	if err != nil {
		return nil, err
	}

	var response []dto.AuctionListResponse
	for _, auction := range auctions {
		response = append(response, dto.AuctionListResponse{
			ItemName:     auction.ItemName,
			EndTime:      auction.EndTime.String(),
			CurrentPrice: auction.CurrentPrice,
			IsClosed:     auction.IsClosed,
		})
	}

	return response, nil
}

// Update implements interfaces.AuctionService.
func (a *auctionService) Update(id uint, request dto.AuctionUpdateRequest) error {
	newAuction := model.Auction{
		ItemName:    request.ItemName,
		Description: request.Description,
		EndTime:     request.EndTime,
	}
	if err := a.auctionRepo.Update(id, newAuction); err != nil {
		return err
	}
	return nil
}

// UpdateCurrentPrice implements interfaces.AuctionService.
func (a *auctionService) UpdateCurrentPrice(id uint, request dto.AuctionUpdateCurrentPriceRequest) (dto.AuctionUpdateCurrentPriceResponse, error) {
	panic("unimplemented")
}

// CloseAuction implements interfaces.AuctionService.
func (a *auctionService) CloseAuction(id uint) error {
	auction, err := a.auctionRepo.Detail(id)
	if err != nil {
		return err
	}

	if auction.EndTime.Before(time.Now()) {
		if err := a.auctionRepo.CloseAuction(id); err != nil {
			return err
		}
		fmt.Println("Auction", id, "closed successfully")
	} else {
		fmt.Println("Auction", id, "is still active")
	}
	return nil
}

// Auto-close semua auction yang sudah lewat end_time
func (a *auctionService) StartAuctionAutoClose() {
	ticker := time.NewTicker(1 * time.Minute)

	go func() {
		for range ticker.C {
			fmt.Println("Running auction auto-close task at", time.Now())

			// Ambil semua auction yang sudah lewat end_time dan belum closed
			auctions, err := a.auctionRepo.GetAuctionsToClose()
			if err != nil {
				fmt.Println("Error fetching auctions to close:", err)
				continue
			}

			// Close semua auction yang memenuhi syarat
			for _, auction := range auctions {
				if err := a.auctionRepo.CloseAuction(auction.ID); err != nil {
					fmt.Println("Failed to close auction", auction.ID, ":", err)
				} else {
					fmt.Println("Auction", auction.ID, "closed successfully")
				}
			}
		}
	}()
}
