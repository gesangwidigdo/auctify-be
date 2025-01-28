package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `json:"name" gorm:"not null; type:varchar(255)"`
	Email    string `json:"email" gorm:"not null; type:varchar(100); unique"`
	Username string `json:"username" gorm:"not null; type:varchar(100); unique"`
	Password string `json:"password" gorm:"not null; type:varchar(255)"`
	Role     string `json:"role" gorm:"not null; type:varchar(50)"`
	Address  string `json:"address" gorm:"type:text"`
}
