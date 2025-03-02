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
func (a *auctionService) Create(id uint, request dto.AuctionCreateRequest) error {
	currentTime := time.Now()
	if request.EndTime.UTC().Before(currentTime) {
		return errors.New("insert valid end time")
	}

	if request.StartPrice <= 0 {
		return errors.New("start price must be greater than 0")
	}

	newAuction := model.Auction{
		ItemID:       request.ItemID,
		StartTime:    currentTime,
		EndTime:      request.EndTime,
		StartPrice:   request.StartPrice,
		CurrentPrice: request.StartPrice,
		IsClosed:     false,
	}

	if err := a.auctionRepo.Create(newAuction); err != nil {
		return err
	}

	return nil
}

// Detail implements interfaces.AuctionService.
func (a *auctionService) Detail(id uint) (dto.AuctionDetailResponse, error) {
	auction, err := a.auctionRepo.Detail(id)
	if err != nil {
		return dto.AuctionDetailResponse{}, err
	}

	return dto.AuctionDetailResponse{
		ID:           auction.ID,
		Seller: dto.AuctionUserResponse{
			Username: auction.Item.User.Username,
		},
		Item: dto.AuctionItemResponse{
			ItemName: auction.Item.ItemName,
			Description: auction.Item.Description,
		},
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
			Seller: dto.AuctionUserResponse{
				Username: auction.Item.User.Username,
			},
			Item: dto.AuctionItemResponse{
				ItemName: auction.Item.ItemName,
				Description: auction.Item.Description,
			},
			EndTime:      auction.EndTime.String(),
			CurrentPrice: auction.CurrentPrice,
			IsClosed:     auction.IsClosed,
		})
	}

	return response, nil
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
	ticker := time.NewTicker(30 * time.Minute)

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
