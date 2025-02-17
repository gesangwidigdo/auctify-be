package model

import (
	"time"

	"gorm.io/gorm"
)

type Auction struct {
	gorm.Model
	ItemName     string    `json:"item_name" gorm:"not null; type:varchar(100)"`
	Description  string    `json:"description" gorm:"type:text"`
	UserID       uint      `json:"user_id" gorm:"not null"`
	StartTime    time.Time `json:"start_time" gorm:"not null"`
	EndTime      time.Time `json:"end_time" gorm:"not null"`
	StartPrice   float64   `json:"start_price" gorm:"not null"`
	CurrentPrice float64   `json:"current_price" gorm:"not null"`
	IsClosed     bool      `json:"is_closed" gorm:"not null"`
}
