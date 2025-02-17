package dto

type OfferCreateRequest struct {
	AuctionID   uint    `json:"auction_id"`
	OfferAmount float64 `json:"offer_amount"`
}

type OfferCreateResponse struct {
	UserID      uint    `json:"user_id"`
	AuctionID   uint    `json:"auction_id"`
	OfferAmount float64 `json:"offer_amount"`
	OfferTime   string  `json:"offer_time"`
}

type OfferListResponse struct {
	User        UserOfferListResponse `json:"user"`
	OfferAmount float64               `json:"offer_amount"`
	OfferTime   string                `json:"offer_time"`
}

type UserOfferListResponse struct {
	Name     string `json:"name"`
	Username string `json:"username"`
}
