package data

import (
	"capstone-mikti/features/bookings"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type BookingData struct {
	db *gorm.DB
}

func New(db *gorm.DB) *BookingData {
	return &BookingData{
		db: db,
	}
}

func (b *BookingData) CheckBookingPaid(userId int) ([]bookings.Booking, error) {
	var dbData = []bookings.Booking{}

	err := b.db.Table("bookings").
		Where("user_id = ? ", userId).
		Where("is_paid = false").
		Where("deleted_at is null").
		Find(&dbData).Error

	if err != nil {
		logrus.Error("Data : Error get user booking status unpaid")
		return nil, err
	}

	return dbData, nil

}

func (b *BookingData) CheckTicket(id int, quantity int) (*bookings.Ticket, error) {
	var dbData bookings.Ticket

	var qry = b.db.Table("tickets").
		Where("id = ?", id).
		Where("deleted_at is null")

	if quantity > 0 {
		qry.Where("quantity >= ? ", quantity)
	}

	if err := qry.Find(&dbData).Error; err != nil {
		logrus.Error("DATA : Error Get By Title : ", err.Error())
		return nil, err
	}

	return &dbData, nil
}

func (b *BookingData) CreateBooking(newData bookings.Booking) (*bookings.Booking, error) {
	var dbData = new(Booking)
	dbData.TicketID = newData.TicketID
	dbData.UserID = newData.UserID
	dbData.FullName = newData.FullName
	dbData.PhoneNumber = newData.PhoneNumber
	dbData.Email = newData.Email
	dbData.BookingDate = newData.ParseBookingDate
	dbData.IsBooked = newData.IsBooked
	dbData.IsPaid = newData.IsPaid

	if err := b.db.Table("bookings").Create(dbData).Error; err != nil {
		logrus.Error("DATA : Create Error : ", err.Error())
		return nil, err
	}

	result := &bookings.Booking{
		ID:               int(dbData.ID),
		TicketID:         dbData.TicketID,
		UserID:           dbData.UserID,
		FullName:         dbData.FullName,
		PhoneNumber:      dbData.PhoneNumber,
		Email:            dbData.Email,
		BookingDate:      dbData.BookingDate.Format("2006-01-02"),
		IsBooked:         dbData.IsBooked,
		IsPaid:           dbData.IsPaid,
		ParseBookingDate: dbData.BookingDate,
	}

	return result, nil
}

func (b *BookingData) CreateBookingDetail(bookingId int, quantity int, price int) (bool, error) {

	total := quantity * price

	var dbData = new(Booking_Detail)
	dbData.BookingID = bookingId
	dbData.Price = price
	dbData.Quantity = quantity
	dbData.Total = total

	if err := b.db.Table("booking_details").Create(dbData).Error; err != nil {
		logrus.Error("DATA : Create Error : ", err.Error())
		return false, err
	}

	return true, nil
}

func (b *BookingData) ChangeQuantityTicket(id int, quantity int) (bool, error) {

	var qry = b.db.Model(&bookings.Ticket{}).
		Where("id = ?", id).
		Where("deleted_at is null").
		Update("quantity", quantity)

	if err := qry.Error; err != nil {
		logrus.Error("Data : Error Update Quantity Ticket")
		return false, err
	}

	return true, nil
}

func (b *BookingData) GetAll(status string) ([]bookings.Booking, error) {
	var result = []bookings.Booking{}

	var qry = b.db.Table("bookings").
		Select("bookings.*, users.username, tickets.name, booking_details.booking_id, booking_details.quantity, booking_details.total").
		Joins("left join users on users.id = bookings.user_id").
		Joins("left join tickets on tickets.id = bookings.ticket_id").
		Joins("left join booking_details on booking_details.booking_id = bookings.id").
		Where("bookings.deleted_at is null")

	switch status {
	case "unpaid":
		qry.Where("bookings.is_paid = false")
	case "paid":
		qry.Where("bookings.is_paid = true")
	}

	if err := qry.Scan(&result).Error; err != nil {
		logrus.Error("Data : Get All Error : ", err.Error())
		return nil, err
	}

	return result, nil

}

func (b *BookingData) CheckBookingDetail(id int) (*bookings.BookingDetail, error) {
	var dbData bookings.BookingDetail

	var qry = b.db.Table("booking_details").
		Select("booking_details.*, bookings.ticket_id").
		Joins("join bookings on bookings.id = booking_details.booking_id").
		Where("booking_details.booking_id = ? ", id).
		Where("booking_details.deleted_at is null").
		First(&dbData)

	if err := qry.Error; err != nil {
		logrus.Error("Data : Get Booking Detail Error : ", err.Error())
		return nil, err
	}

	return &dbData, nil
}

func (b *BookingData) DeleteBooking(id int) (bool, error) {

	var booking Booking

	if err := b.db.Preload("BookingDetails").Where("is_paid = false").
		Where("deleted_at is null").First(&booking, id).Error; err != nil {
		logrus.Error("Data : Get Booking Error : ", err.Error())
		return false, err
	}

	if err := b.db.Select(clause.Associations).
		Where("is_paid = false").
		Where("deleted_at is null").
		Delete(&booking).Error; err != nil {
		logrus.Error("Data : Delete Booking Error : ", err.Error())
		return false, err
	}

	return true, nil
}

func (b *BookingData) GetDetail(id int) (*bookings.Booking, error) {
	var booking bookings.Booking

	var qry = b.db.Table("bookings").
		Select("bookings.*, users.username, tickets.name, tickets.price, tickets.event_id, booking_details.quantity, booking_details.total, events.event_title").
		Joins("left join users on users.id = bookings.user_id").
		Joins("left join tickets on tickets.id = bookings.ticket_id").
		Joins("left join booking_details on booking_details.booking_id = bookings.id").
		Joins("left join events on events.id = tickets.event_id").
		Where("bookings.id = ? ", id).
		Where("bookings.deleted_at is null").
		First(&booking)

	if err := qry.Error; err != nil {
		logrus.Error("Data : Get Detail Error : ", err.Error())
		return nil, err
	}

	return &booking, nil
}
