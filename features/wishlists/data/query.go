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

// Create
func (wd *WishlistData) Create(new_data wishlists.Wishlist) (*wishlists.Wishlist, error) {
	wishlist := &wishlists.Wishlist{
		UserID:  new_data.UserID,
		EventID: new_data.EventID,
	}

	if err := wd.db.Create(wishlist).Error; err != nil {
		logrus.Error("DATA : Create Error : ", err.Error())
		return nil, err
	}

	return wishlist, nil
}

// GetAll
func (wd *WishlistData) GetAll() ([]wishlists.WishlistInfo, error) {
	var result = []wishlists.WishlistInfo{}

	err := wd.db.Table("wishlists").
		Select("wishlists.id, wishlists.user_id, users.username, wishlists.event_id, events.event_title").
		Joins("JOIN users ON users.id = wishlists.user_id").
		Joins("JOIN events ON events.id = wishlists.event_id").
		Scan(&result).Error

	if err != nil {
		logrus.Error("DATA : GetAll Error : ", err.Error())
		return result, err
	}

	return result, nil
}

// GetByUserID
func (wd *WishlistData) GetByUserID(user_id int) ([]wishlists.WishlistInfo, error) {
	var result = []wishlists.WishlistInfo{}

	err := wd.db.Table("wishlists").
		Select("wishlists.id, wishlists.user_id, users.username, wishlists.event_id, events.event_title").
		Joins("JOIN users ON users.id = wishlists.user_id").
		Joins("JOIN events ON events.id = wishlists.event_id").
		Where("wishlists.user_id = ?", user_id).
		Scan(&result).Error

	if err != nil {
		logrus.Error("DATA : GetByUserID Error : ", err.Error())
		return result, err
	}

	return result, nil
}

// Delete
func (wd *WishlistData) Delete(id uint) error {
	delete := wd.db.Delete(&wishlists.Wishlist{}, id)
	if delete.Error != nil {
		logrus.Error("DATA : Delete Error : ", delete.Error.Error())
		return delete.Error
	}

	if delete.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}
