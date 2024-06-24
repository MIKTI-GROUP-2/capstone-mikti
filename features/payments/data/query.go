package data

import (
	// _bookings "capstone-mikti/features/bookings/data"
	"capstone-mikti/features/payments"
	"time"

	// _tickets "capstone-mikti/features/tickets/data"
	"errors"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type PaymentData struct {
	db *gorm.DB
}

func New(db *gorm.DB) *PaymentData {
	return &PaymentData{
		db: db,
	}
}

func (p *PaymentData) InsertPaymentMaster(newData payments.PaymentMaster) (*payments.PaymentMaster, error) {
	var dbData = new(Payment)
	dbData.BookID = newData.BookID
	dbData.UserID = newData.UserID
	dbData.OrderID = newData.OrderID
	dbData.MidtransID = newData.MidtransID
	dbData.GrandTotal = newData.GrandTotal
	dbData.TransactionDate = newData.TransactionDateParse
	dbData.PaymentStatus = newData.PaymentStatus
	dbData.PaymentType = newData.PaymentType
	dbData.Status = newData.Status

	if err := p.db.Create(dbData).Error; err != nil {
		logrus.Error("Error Create Payment : ", err.Error())
		return nil, errors.New("ERROR Create Payment")
	}

	return &newData, nil
}

func (p *PaymentData) InsertPaymentDetail(newData payments.PaymentDetail) (*payments.PaymentDetail, error) {
	var dbData = new(PaymentDetail)
	dbData.PaymentID = newData.PaymentID
	dbData.VoucherID = newData.VoucherID
	dbData.Price = newData.Price
	dbData.Qty = newData.Qty
	dbData.Total = newData.Total
	dbData.VoucherPrice = newData.VoucherPrice

	if err := p.db.Table("payment_details").Create(dbData).Error; err != nil {
		logrus.Error("Error Create Payment : ", err.Error())
		return nil, errors.New("ERROR Create Payment")
	}

	return &newData, nil
}

func (p *PaymentData) GetByIDMidtrans(midtrans_id string) (uint, error) {
	var listPayment payments.PaymentMaster

	if err := p.db.Table("payments").Where("midtrans_id = ?", midtrans_id).Scan(&listPayment).Error; err != nil {
		logrus.Error("Error Get By ID Midtrans : ", err.Error())
		return 0, errors.New("ERROR Get By ID Midtrans")
	}

	id := listPayment.ID
	return id, nil
}

func (p *PaymentData) GetQtyAndVoucByID(payment_id uint) (int, uint, error) {
	var listPayment payments.PaymentDetail

	if err := p.db.Table("payment_details").Where("payment_id = ?", payment_id).Scan(&listPayment).Error; err != nil {
		logrus.Error("Error Get By ID Midtrans : ", err.Error())
		return 0, 0, errors.New("ERROR Get By ID Midtrans")
	}

	qtyTicket, voucherID := listPayment.Qty, listPayment.VoucherID
	return qtyTicket, voucherID, nil
}

func (p *PaymentData) GetTicketByID(payment_id uint) (uint, error) {
	var ticket payments.TicketInfo

	var qry = p.db.Table("payments").
		Select("bookings.ticket_id").
		Joins("JOIN bookings ON payments.book_id = bookings.id").
		Where("payments.id = ?", payment_id).
		Scan(&ticket)

	if err := qry.Error; err != nil {
		logrus.Error("Error Get By ID Midtrans : ", err.Error())
		return 0, errors.New("ERROR Get By ID Midtrans")
	}

	ticketID := ticket.TicketID
	return ticketID, nil
}

func (p *PaymentData) GetAndUpdate(newData payments.UpdatePayment, id string) (bool, error) {

	var payment Payment
	var paymentDetail PaymentDetail
	// db := ad.db

	_ = p.db.Where("midtrans_id = ?", id).First(&payment)
	paymentID := payment.ID

	_ = p.db.Table("payment_details").Where("payment_id = ?", paymentID).First(&paymentDetail)

	// fmt.Println("This is the new payment status: ", newData.PaymentStatus)

	if newData.PaymentStatus == 2 {
		updatePayment := p.db.Table("payments").Where("id = ?", paymentID).Updates(Payment{
			PaymentStatus: newData.PaymentStatus,
			Status:        "Paid",
		})

		if updatePayment.Error != nil {
			return false, nil
		}
	}

	return true, nil
}

// Admin
func (p *PaymentData) GetAll(status string, limit, offset, timePublication int) ([]payments.PaymentInfo, error) {
	var listPayment = []payments.PaymentInfo{}
	var now time.Time

	var qry = p.db.Table("payments").
		Joins("JOIN bookings ON bookings.id = payments.book_id").
		Joins("JOIN tickets ON tickets.id = bookings.ticket_id").
		Joins("JOIN events ON events.id = tickets.event_id").
		Select("payments.*", "bookings.full_name", "bookings.phone_number", "bookings.email", "bookings.booking_date", "tickets.ticket_date", "events.event_title", "events.city", "events.address").
		Order("payments.created_at DESC")

	now = time.Now()

	if status != "" {
		qry.Where("payments.status = ?", status)
	}

	switch timePublication {
	case 1:
		before := now.AddDate(0, 0, -7)
		qry.Where("payments.created_at BETWEEN ? AND ?", before, now)
	case 2:
		before := now.AddDate(0, 0, -30)
		qry.Where("payments.created_at BETWEEN ? AND ?", before, now)
	}

	if limit != 0 {
		qry.Limit(limit)
	}

	if offset != 0 {
		qry.Offset(offset)
	}

	if err := qry.Scan(&listPayment).Error; err != nil {
		logrus.Error("Error Get All : ", err.Error())
		return nil, errors.New("ERROR Get All")
	}

	return listPayment, nil
}

// User
func (p *PaymentData) GetByUserID(id uint, status string, limit, offset, timePublication int) ([]payments.PaymentInfo, error) {
	var listPayment = []payments.PaymentInfo{}
	var now time.Time

	var qry = p.db.Table("payments").
		Joins("JOIN bookings ON bookings.id = payments.book_id").
		Joins("JOIN tickets ON tickets.id = bookings.ticket_id").
		Joins("JOIN events ON events.id = tickets.event_id").
		Select("payments.*", "bookings.full_name", "bookings.phone_number", "bookings.email", "bookings.booking_date", "tickets.ticket_date", "events.event_title", "events.city", "events.address").
		Order("payments.created_at DESC").
		Where("payments.user_id = ?", id)

	now = time.Now()

	if status != "" {
		qry.Where("payments.status = ?", status)
	}

	switch timePublication {
	case 1:
		before := now.AddDate(0, 0, -7)
		qry.Where("payments.created_at BETWEEN ? AND ?", before, now)
	case 2:
		before := now.AddDate(0, 0, -30)
		qry.Where("payments.created_at BETWEEN ? AND ?", before, now)
	}

	if limit != 0 {
		qry.Limit(limit)
	}

	if offset != 0 {
		qry.Offset(offset)
	}

	if err := qry.Scan(&listPayment).Error; err != nil {
		logrus.Error("Error Get All : ", err.Error())
		return nil, errors.New("ERROR Get All")
	}

	return listPayment, nil
}
