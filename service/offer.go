package service

import (
	"errors"
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

	if err := o.offerRepo.Create(offerResponse); err != nil {
		return dto.OfferCreateResponse{}, err
	}

	if err := o.auctionRepo.UpdateCurrentPrice(auction.ID, offer.OfferAmount); err != nil {
		return dto.OfferCreateResponse{}, err
	}

	return dto.OfferCreateResponse{
		UserID:      offerResponse.UserID,
		AuctionID:   offerResponse.AuctionID,
		OfferAmount: offerResponse.OfferAmount,
		OfferTime:   offerResponse.OfferTime.String(),
	}, nil
}
