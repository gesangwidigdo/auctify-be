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

	if auction.EndTime.Before(time.Now()) {
		o.auctionRepo.CloseAuction(auction.ID)
		return dto.OfferCreateResponse{}, errors.New("auction has been closed")
	}

	if auction.Item.UserID == userId {
		return dto.OfferCreateResponse{}, errors.New("cannot make an offer on your own auction")
	}

	if auction.IsClosed {
		return dto.OfferCreateResponse{}, errors.New("auction is closed")
	}

	if offer.OfferAmount <= auction.CurrentPrice {
		return dto.OfferCreateResponse{}, errors.New("offer amount must be higher than the current auction price")
	}

	newOffer := model.Offer{
		UserID:      userId,
		AuctionID:   offer.AuctionID,
		OfferAmount: offer.OfferAmount,
		OfferTime:   currentTime,
	}

	if err := o.offerRepo.Create(newOffer); err != nil {
		return dto.OfferCreateResponse{}, err
	}

	if err := o.auctionRepo.UpdateCurrentPrice(auction.ID, offer.OfferAmount); err != nil {
		return dto.OfferCreateResponse{}, err
	}

	return dto.OfferCreateResponse{
		UserID:      newOffer.UserID,
		AuctionID:   newOffer.AuctionID,
		OfferAmount: newOffer.OfferAmount,
		OfferTime:   newOffer.OfferTime.String(),
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
