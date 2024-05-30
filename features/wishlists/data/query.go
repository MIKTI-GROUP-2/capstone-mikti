package data

import (
	"capstone-mikti/features/wishlists"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// Repository

type WishlistData struct {
	db *gorm.DB
}

func New(db *gorm.DB) *WishlistData {
	return &WishlistData{
		db: db,
	}
}

// GetAll implements wishlists.WishlistDataInterface.
func (wd *WishlistData) GetAll() ([]wishlists.WishlistInfo, error) {
	var listWishlist = []wishlists.WishlistInfo{}

	err := wd.db.Table("wishlists").Joins("JOIN events ON events.id = wishlists.event_id").Joins("JOIN users ON users.id = wishlists.user_id").Scan(&listWishlist).Error

	if err != nil {
		logrus.Error("DATA : GetAll Error : ", err.Error())
		return listWishlist, err
	}

	return listWishlist, nil
}
