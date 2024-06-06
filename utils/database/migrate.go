package database

import (
	dataCategory "capstone-mikti/features/categories/data"
	dataEvent "capstone-mikti/features/events/data"
	dataUser "capstone-mikti/features/users/data"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(dataUser.User{})
	db.AutoMigrate(dataUser.UserResetPass{})
	db.AutoMigrate(dataCategory.Category{})
	db.AutoMigrate(dataEvent.Event{})
}
