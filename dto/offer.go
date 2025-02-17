package dto

type OfferCreateRequest struct {
	AuctionID   uint    `json:"auction_id"`
	OfferAmount float64 `json:"offer_amount"`
}

type OfferCreateResponse struct {
	UserId      uint    `json:"user_id"`
	AuctionID   uint    `json:"auction_id"`
	OfferAmount float64 `json:"offer_amount"`
	OfferTime   string  `json:"offer_time"`
}
