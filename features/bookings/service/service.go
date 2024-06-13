package service

import (
	"capstone-mikti/features/bookings"
	"capstone-mikti/helper/jwt"
	"errors"
	"time"

	"github.com/sirupsen/logrus"
)

type BookingService struct {
	data bookings.BookingDataInterface
	jwt  jwt.JWTInterface
}

func New(b bookings.BookingDataInterface, j jwt.JWTInterface) *BookingService {
	return &BookingService{
		data: b,
		jwt:  j,
	}
}

func (b *BookingService) CreateBooking(newData bookings.Booking) (*bookings.Booking, error) {

	checkBookingPaid, err := b.data.CheckBookingPaid(newData.UserID)
	if err != nil {
		logrus.Error("Service : Error get data booking user")
		return nil, errors.New("Error Get data booking user")
	}

	if len(checkBookingPaid) > 0 {
		logrus.Error("Service : Booking already exist with status unpaid")
		return nil, errors.New("Booking already exist with status unpaid")
	}

	checkTicketData, err := b.data.CheckTicket(newData.TicketID, newData.Quantity)
	if err != nil {
		logrus.Error("Service : Error get data ticket")
		return nil, errors.New("Error Get data ticket")
	}

	if checkTicketData.ID == 0 {
		logrus.Error("Service : Ticket Not Found")
		return nil, errors.New("Ticket Not Found")
	}

	if checkTicketData.Quantity == 0 {
		logrus.Error("Service : Quantity too much")
		return nil, errors.New("Ticket Quantity is too much")
	}

	//Waktu
	layout := "2006-01-02"

	parseBookingDate, _ := time.Parse(layout, checkTicketData.TicketDate)
	newData.ParseBookingDate = parseBookingDate

	newData.IsBooked = true
	newData.IsPaid = false

	resultBooking, err := b.data.CreateBooking(newData)
	if err != nil {
		logrus.Error("Create Booking Failed")
		return nil, errors.New("Create Booking Failed")
	}

	err = b.data.CreateBookingDetail(resultBooking.ID, newData.Quantity, checkTicketData.Price)
	if err != nil {
		logrus.Error("Create Booking Detail Failed")
		return nil, errors.New("Create Booking Detail Failed")
	}

	quantityNow := checkTicketData.Quantity - newData.Quantity
	err = b.data.ChangeQuantityTicket(newData.TicketID, quantityNow)
	if err != nil {
		logrus.Error("Change Quantity Ticket Failed")
		return nil, errors.New("Change Quantity Ticket Failed")
	}

	return resultBooking, nil
}

func (b *BookingService) GetAll(status string, userID int) ([]bookings.Booking, error) {
	result, err := b.data.GetAll(status, userID)

	if err != nil {
		logrus.Error("Get Booking Failed")
		return nil, errors.New("Get Booking Failed")
	}

	return result, nil
}

func (b *BookingService) DeleteBooking(id int, userID int) (bool, error) {
	checkBookingDetail, err := b.data.CheckBookingDetail(id)

	if err != nil {
		logrus.Error("Get Booking Detail Failed")
		return false, errors.New("Get Booking Detail Failed")
	}

	checkTicket, err := b.data.CheckTicket(checkBookingDetail.TicketID, 0)

	quantityNow := checkBookingDetail.Quantity + checkTicket.Quantity

	err = b.data.ChangeQuantityTicket(checkTicket.ID, quantityNow)

	if err != nil {
		logrus.Error("Update Quantity Ticket Failed")
		return false, errors.New("Update Quantity Ticket Failed")
	}

	result, err := b.data.DeleteBooking(id, userID)

	if err != nil {
		logrus.Error("Service : Delete Booking Failed")
		return false, errors.New("Error Delete Booking")
	}

	return result, nil
}

func (b *BookingService) GetDetail(id int, userID int) (*bookings.Booking, error) {
	bookingDetail, err := b.data.GetDetail(id, userID)

	if err != nil {
		logrus.Error("Service : Get Detail Booking Failed")
		return nil, errors.New("Error Get Detail Booking")
	}

	return bookingDetail, nil
}
