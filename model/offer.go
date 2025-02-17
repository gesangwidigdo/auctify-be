package model

import (
	"time"

	"gorm.io/gorm"
)

type Offer struct {
	gorm.Model
	UserID      uint      `json:"user_id"`
	AuctionID   uint      `json:"auction_id"`
	OfferAmount float64   `json:"offer_amount"`
	OfferTime   time.Time `json:"offer_time"`

	User    User    `json:"user"`
	Auction Auction `json:"auction"`
}
