package model

import "gorm.io/gorm"

type Item struct {
	gorm.Model
	ItemName    string `json:"item_name" gorm:"not null; type:varchar(100)"`
	Description string `json:"description" gorm:"type:text"`
	UserID      uint   `json:"user_id" gorm:"not null"`

	// Relationship
	User User `json:"user"`
}
