package model

import "gorm.io/gorm"

func Migrate(db *gorm.DB) error {
	err := db.AutoMigrate(
		&User{},
		&Item{},
		&Auction{},
		&Offer{},
	)
	if err != nil {
		return err
	}

	return nil
}
