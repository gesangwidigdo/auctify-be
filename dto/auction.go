package dto

import "time"

type AuctionCreateRequest struct {
	ItemID     uint      `json:"item_id" binding:"required"`
	EndTime    time.Time `json:"end_time" binding:"required"` // format example: '2025-03-05T07:00:00+07:00'
	StartPrice float64   `json:"start_price" binding:"required"`
}

type AuctionItemResponse struct {
	ItemName    string `json:"item_name"`
	Description string `json:"description"`
}

type AuctionUserResponse struct {
	Username string `json:"username"`
}

type AuctionDetailResponse struct {
	ID           uint                `json:"id"`
	Item         AuctionItemResponse `json:"item"`
	Seller       AuctionUserResponse `json:"seller"`
	StartTime    string              `json:"start_time"`
	EndTime      string              `json:"end_time"`
	StartPrice   float64             `json:"start_price"`
	CurrentPrice float64             `json:"current_price"`
	IsClosed     bool                `json:"is_closed"`
}

type AuctionListResponse struct {
	Item         AuctionItemResponse `json:"item"`
	Seller       AuctionUserResponse `json:"seller"`
	EndTime      string              `json:"end_time"`
	CurrentPrice float64             `json:"current_price"`
	IsClosed     bool                `json:"is_closed"`
}

type AuctionUpdateRequest struct {
	EndTime time.Time `json:"end_time"`
}

type AuctionUpdateCurrentPriceRequest struct {
	CurrentPrice float64 `json:"current_price" binding:"required"`
}

type AuctionUpdateCurrentPriceResponse struct {
	Status string `json:"status"`
}
