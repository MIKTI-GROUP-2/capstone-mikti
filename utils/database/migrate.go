package database

import (
	dataEvent "capstone-mikti/features/events/data"
	dataUser "capstone-mikti/features/users/data"
	dataWishlist "capstone-mikti/features/wishlists/data"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(dataUser.User{})
	db.AutoMigrate(dataUser.UserResetPass{})
	db.AutoMigrate(dataEvent.Event{})
	db.AutoMigrate(dataWishlist.Wishlist{})
}
