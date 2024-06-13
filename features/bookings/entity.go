package bookings

import (
	"time"

	"github.com/labstack/echo/v4"
)

type Booking struct {
	ID               int       `json:"id"`
	EventID          int       `json:"event_id"`
	TicketID         int       `json:"ticket_id"`
	UserID           int       `json:"user_id"`
	FullName         string    `json:"full_name"`
	PhoneNumber      string    `json:"phone_number"`
	Email            string    `json:"email"`
	BookingDate      string    `json:"booking_date"`
	ParseBookingDate time.Time `json:"-"`
	IsBooked         bool      `json:"is_booked"`
	IsPaid           bool      `json:"is_paid"`
	Quantity         int       `json:"quantity"`
	Total            int       `json:"total"`
	Name             string    `json:"name"`
	Username         string    `json:"username"`
	Price            int       `json:"price"`
	EventTitle       string    `json:"event_title"`
}

type BookingDetail struct {
	ID        int `json:"id"`
	BookingID int `json:"booking_id"`
	Price     int `json:"price"`
	Quantity  int `json:"quantity"`
	Total     int `json:"total"`
	TicketID  int `json:"ticket_id"`
}

type QuantityTicket struct {
	Quantity int `json:"quantity"`
}

type Ticket struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	TicketDate string `json:"ticket_date"`
	Quantity   int    `json:"quantity"`
	Price      int    `json:"price"`
}

type BookingHandlerInterface interface {
	CreateBooking() echo.HandlerFunc
	GetAll() echo.HandlerFunc
	DeleteBooking() echo.HandlerFunc
	GetDetail() echo.HandlerFunc
}

type BookingServiceInterface interface {
	CreateBooking(newData Booking) (*Booking, error)
	GetAll(status string, userID int) ([]Booking, error)
	DeleteBooking(id int, userID int) (bool, error)
	GetDetail(id int, userID int) (*Booking, error)
}

type BookingDataInterface interface {
	CheckTicket(id int, quantity int) (*Ticket, error)
	CheckBookingPaid(userId int) ([]Booking, error)
	CreateBooking(newData Booking) (*Booking, error)
	CreateBookingDetail(bookingId int, quantity int, price int) error
	ChangeQuantityTicket(id int, quntity int) error
	GetAll(status string, userID int) ([]Booking, error)
	GetDetail(id int, userID int) (*Booking, error)
	DeleteBooking(id int, userID int) (bool, error)
	CheckBookingDetail(id int) (*BookingDetail, error)
}
