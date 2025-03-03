package model

import (
	"time"

	"gorm.io/gorm"
)

type Auction struct {
	gorm.Model
	StartTime    time.Time `json:"start_time" gorm:"not null"`
	EndTime      time.Time `json:"end_time" gorm:"not null"`
	StartPrice   float64   `json:"start_price" gorm:"not null"`
	CurrentPrice float64   `json:"current_price" gorm:"not null"`
	IsClosed     bool      `json:"is_closed" gorm:"not null"`
	ItemID       uint      `json:"item_id" gorm:"not null;unique"`

	// Relationship
	Item   Item    `json:"item" gorm:"foreignKey:ItemID"`
	Offers []Offer `json:"offers"`
}
