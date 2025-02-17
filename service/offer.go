package service

import (
	"errors"
	"sync"
	"time"

	"github.com/gesangwidigdo/auctify-be/dto"
	"github.com/gesangwidigdo/auctify-be/interfaces"
	"github.com/gesangwidigdo/auctify-be/model"
)

type offerService struct {
	offerRepo   interfaces.OfferRepository
	auctionRepo interfaces.AuctionRepository
}

func NewOfferService(offerRepo interfaces.OfferRepository, auctionRepo interfaces.AuctionRepository) interfaces.OfferService {
	return &offerService{offerRepo, auctionRepo}
}

// Create implements interfaces.OfferService.
func (o *offerService) Create(userId uint, offer dto.OfferCreateRequest) (dto.OfferCreateResponse, error) {
	currentTime := time.Now()

	auction, err := o.auctionRepo.Detail(offer.AuctionID)
	if err != nil {
		return dto.OfferCreateResponse{}, err
	}

	if auction.UserID == userId {
		return dto.OfferCreateResponse{}, errors.New("cannot make an offer on your own auction")
	}

	if auction.IsClosed {
		return dto.OfferCreateResponse{}, errors.New("auction is closed")
	}

	if offer.OfferAmount <= auction.CurrentPrice {
		return dto.OfferCreateResponse{}, errors.New("offer amount must be higher than the current auction price")
	}

	offerResponse := model.Offer{
		UserID:      userId,
		AuctionID:   offer.AuctionID,
		OfferAmount: offer.OfferAmount,
		OfferTime:   currentTime,
	}

	var wg sync.WaitGroup
	errChan := make(chan error, 2)

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := o.offerRepo.Create(offerResponse); err != nil {
			errChan <- err
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := o.auctionRepo.UpdateCurrentPrice(auction.ID, offer.OfferAmount); err != nil {
			errChan <- err
		}
	}()

	wg.Wait()
	close(errChan)

	for err := range errChan {
		if err != nil {
			return dto.OfferCreateResponse{}, err
		}
	}

	return dto.OfferCreateResponse{
		UserID:      offerResponse.UserID,
		AuctionID:   offerResponse.AuctionID,
		OfferAmount: offerResponse.OfferAmount,
		OfferTime:   offerResponse.OfferTime.String(),
	}, nil
}

// List implements interfaces.OfferService.
func (o *offerService) List(auctionId uint) ([]dto.OfferListResponse, error) {
	offers, err := o.offerRepo.List(auctionId)
	if err != nil {
		return nil, err
	}

	offerListResponse := make([]dto.OfferListResponse, 0)
	for _, offer := range offers {
		offerListResponse = append(offerListResponse, dto.OfferListResponse{
			User: dto.UserOfferListResponse{
				Name:     offer.User.Name,
				Username: offer.User.Username,
			},
			OfferAmount: offer.OfferAmount,
			OfferTime:   offer.OfferTime.String(),
		})
	}

	return offerListResponse, nil
}
