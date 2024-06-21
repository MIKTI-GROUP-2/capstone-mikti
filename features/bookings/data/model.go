package data

import (
	paymentData "capstone-mikti/features/payments/data"
	ticketData "capstone-mikti/features/tickets/data"
	userData "capstone-mikti/features/users/data"
	"time"

	"gorm.io/gorm"
)

type Booking struct {
	*gorm.Model
	UserID         int                 `gorm:"column:user_id;type:int"`
	TicketID       int                 `gorm:"column:ticket_id;type:int"`
	FullName       string              `gorm:"column:full_name;type:varchar(255)"`
	PhoneNumber    string              `gorm:"column:phone_number;type:varchar(255)"`
	Email          string              `gorm:"column:email;type:varchar(255)"`
	BookingDate    time.Time           `gorm:"column:booking_date;type:date"`
	IsBooked       bool                `gorm:"column:is_booked;type:bool"`
	IsPaid         bool                `gorm:"column:is_paid;type:bool"`
	User           userData.User       `gorm:"foreignKey:UserID;references:ID"`
	Ticket         ticketData.Ticket   `gorm:"foreignKey:TicketID;references:ID"`
	Payment        paymentData.Payment `gorm:"foreignKey:BookID"`
	BookingDetails []Booking_Detail    `gorm:"foreignKey:BookingID"`
}

type Booking_Detail struct {
	*gorm.Model
	BookingID int     `gorm:"column:booking_id;type:int"`
	Price     int     `gorm:"column:price;type:int"`
	Quantity  int     `gorm:"column:quantity;type:int"`
	Total     int     `gorm:"column:total;type:int"`
	Booking   Booking `gorm:"foreignKey:BookingID;references:ID"`
}
