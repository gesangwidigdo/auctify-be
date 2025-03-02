package dto

import "time"

type AuctionCreateRequest struct {
	ItemName    string   `json:"item_name" binding:"required"`
	Description string   `json:"description"`
	EndTime     time.Time `json:"end_time" binding:"required"` // format example: '2025-03-05T07:00:00+07:00'
	StartPrice  float64  `json:"start_price" binding:"required"`
}

type AuctionCreateResponse struct {
	ItemName     string   `json:"item_name"`
	Description  string   `json:"description"`
	EndTime      time.Time `json:"end_time"`
	StartPrice   float64  `json:"start_price"`
	CurrentPrice float64  `json:"current_price"`
}

type AuctionDetailResponse struct {
	ID           uint    `json:"id"`
	ItemName     string  `json:"item_name"`
	Description  string  `json:"description"`
	StartTime    string  `json:"start_time"`
	EndTime      string  `json:"end_time"`
	StartPrice   float64 `json:"start_price"`
	CurrentPrice float64 `json:"current_price"`
	IsClosed     bool    `json:"is_closed"`
}

type AuctionListResponse struct {
	ItemName     string  `json:"item_name"`
	EndTime      string  `json:"end_time"`
	CurrentPrice float64 `json:"current_price"`
	IsClosed     bool    `json:"is_closed"`
}

type AuctionUpdateRequest struct {
	ItemName    string `json:"item_name"`
	Description string `json:"description"`
	EndTime     time.Time `json:"end_time"`
}

type AuctionUpdateCurrentPriceRequest struct {
	CurrentPrice float64 `json:"current_price" binding:"required"`
}

type AuctionUpdateCurrentPriceResponse struct {
	Status string `json:"status"`
}
