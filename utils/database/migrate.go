package database

import (
	dataBooking "capstone-mikti/features/bookings/data"
	dataCategory "capstone-mikti/features/categories/data"
	dataEvent "capstone-mikti/features/events/data"
	dataPayment "capstone-mikti/features/payments/data"
	dataTicket "capstone-mikti/features/tickets/data"
	dataUser "capstone-mikti/features/users/data"
	dataVoucher "capstone-mikti/features/vouchers/data"
	dataWishlist "capstone-mikti/features/wishlists/data"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(dataUser.User{})
	db.AutoMigrate(dataUser.UserResetPass{})
	db.AutoMigrate(dataUser.UserVerification{})
	db.AutoMigrate(dataTicket.Ticket{})
	db.AutoMigrate(dataCategory.Category{})
	db.AutoMigrate(dataEvent.Event{})
	db.AutoMigrate(dataVoucher.Voucher{})
	db.AutoMigrate(dataWishlist.Wishlist{})
	db.AutoMigrate(dataTicket.Ticket{})
	db.AutoMigrate(dataBooking.Booking{})
	db.AutoMigrate(dataBooking.Booking_Detail{})
	db.AutoMigrate(dataPayment.Payment{})
	db.AutoMigrate(dataPayment.PaymentDetail{})
}
