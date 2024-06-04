package database

import (
	dataTicket "capstone-mikti/features/tickets/data"
	dataUser "capstone-mikti/features/users/data"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(dataUser.User{})
	db.AutoMigrate(dataUser.UserResetPass{})
	db.AutoMigrate(dataTicket.Ticket{})
}
