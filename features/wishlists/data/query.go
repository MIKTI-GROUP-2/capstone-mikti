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

// CheckUnique
func (wd *WishlistData) CheckUnique(user_id, event_id int) ([]wishlists.Wishlist, error) {
	// Get Entity
	var wishlist = []wishlists.Wishlist{}

	// Query
	err := wd.db.Table("wishlists").
		Select("*").
		Where("wishlists.user_id = ? AND wishlists.event_id = ?", user_id, event_id).
		Find(&wishlist).Error

	// Error Handling
	if err != nil {
		logrus.Error("DATA : CheckUnique Error : ", err.Error())
		return nil, err
	}

	return wishlist, nil
}

// Create
func (wd *WishlistData) Create(user_id int, new_data wishlists.Wishlist) (*wishlists.Wishlist, error) {
	// Get Entity
	wishlist := &wishlists.Wishlist{
		UserID:  user_id,
		EventID: new_data.EventID,
	}

	// Query
	err := wd.db.Create(wishlist).Error

	// Error Handling
	if err != nil {
		logrus.Error("DATA : Create Error : ", err.Error())
		return nil, err
	}

	return wishlist, nil
}

// GetAll
func (wd *WishlistData) GetAll(user_id int) ([]wishlists.WishlistInfo, error) {
	// Get Entity
	var wishlist = []wishlists.WishlistInfo{}

	// Query
	err := wd.db.Table("wishlists").
		Select("wishlists.id, wishlists.user_id, wishlists.event_id, events.event_title").
		Joins("JOIN events ON events.id = wishlists.event_id").
		Where("wishlists.user_id = ?", user_id).
		Find(&wishlist).Error

	// Error Handling
	if err != nil {
		logrus.Error("DATA : GetAll Error : ", err.Error())
		return nil, err
	}

	return wishlist, nil
}

// GetByID
func (wd *WishlistData) GetByID(user_id, id int) ([]wishlists.WishlistInfo, error) {
	// Get Entity
	var wishlist = []wishlists.WishlistInfo{}

	// Query
	err := wd.db.Table("wishlists").
		Select("wishlists.id, wishlists.user_id, wishlists.event_id, events.event_title").
		Joins("JOIN events ON events.id = wishlists.event_id").
		Where("wishlists.user_id = ? AND wishlists.id = ?", user_id, id).
		Find(&wishlist).Error

	// Error Handling
	if err != nil {
		logrus.Error("DATA : GetByID Error : ", err.Error())
		return nil, err
	}

	return wishlist, nil
}

// Delete
func (wd *WishlistData) Delete(user_id int, event_id int) error {
	// Get Entity
	var wishlist = &wishlists.Wishlist{}

	// Query
	err := wd.db.Where("user_id = ? AND event_id = ?", user_id, event_id).
		Delete(wishlist)

	// Error Handling
	if err.Error != nil {
		logrus.Error("DATA : Delete Error : ", err.Error)
		return err.Error
	}

	if err.RowsAffected == 0 {
		logrus.Warn("DATA : Record Not Found : ")
		return gorm.ErrRecordNotFound
	}

	return nil
}
