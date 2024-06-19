package payments

import (
	"time"

	"github.com/labstack/echo/v4"
)

type Payment struct {
	BookID      uint   `json:"book_id"`
	UserID      uint   `json:"user_id"`
	OrderID     string `json:"order_id"`
	MidtransID  string `json:"midtrans_id"`
	VoucherID   uint   `json:"voucher_id"`
	VoucherCode string `json:"voucher_code"`

	Price int `json:"price"`
	Qty   int `json:"qty"`
	Total int `json:"total"`

	GrandTotal int `json:"grand_total"`

	TransactionDateParse time.Time `json:"transaction_date_parse"`
	TransactionDate      string    `json:"transaction_date"`

	PaymentStatus uint   `json:"payment_status"`
	PaymentType   string `json:"payment_type"`
	Status        string `json:"status"`
}

type PaymentMaster struct {
	ID         uint   `json:"id"`
	BookID     uint   `json:"book_id"`
	UserID     uint   `json:"user_id"`
	OrderID    string `json:"order_id"`
	MidtransID string `json:"midtrans_id"`

	GrandTotal int `json:"grand_total"`

	TransactionDateParse time.Time `json:"transaction_date_parse"`
	TransactionDate      string    `json:"transaction_date"`

	PaymentStatus uint   `json:"payment_status"`
	PaymentType   string `json:"payment_type"`
	Status        string `json:"status"`
}

type PaymentDetail struct {
	PaymentID    uint   `json:"payment_id"`
	VoucherID    uint   `json:"voucher_id"`
	VoucherCode  string `json:"voucher_code"`
	Price        int    `json:"price"`
	Qty          int    `json:"qty"`
	Total        int    `json:"total"`
	VoucherPrice int    `json:"voucher_price"`
}

type PaymentInfo struct {
	FullName        string `json:"full_name"`
	PhoneNumber     string `json:"phone_number"`
	Email           string `json:"email"`
	BookingDate     string `json:"booking_date"`
	TransactionDate string `json:"transaction_date"`
	TicketDate      string `json:"ticket_date"`

	EventTitle string `json:"event_title"`
	City       string `json:"city"`
	Address    string `json:"address"`

	OrderID string `json:"order_id"`

	GrandTotal int `json:"grand_total"`

	Status      string `json:"status"`
	PaymentType string `json:"payment_type"`
}

type TicketInfo struct {
	TicketID uint `json:"ticket_id"`
}

type UpdatePayment struct {
	PaymentStatus uint `json:"payment_status"`
}

type PaymentHandlerInterface interface {
	CreatePayment() echo.HandlerFunc
	NotifPayment() echo.HandlerFunc

	GetAll() echo.HandlerFunc
}

type PaymentServiceInterface interface {
	CreatePayment(newData Payment) (*Payment, map[string]interface{}, error)
	UpdatePayment(notificationPayload map[string]interface{}, newData UpdatePayment) (bool, error)

	// Admin
	GetAll(status string, limitInt, offsetInt, timePublication int) ([]PaymentInfo, error)

	// User
	GetByUserID(id uint, status string, limitInt, offsetInt, timePublication int) ([]PaymentInfo, error)
}

type PaymentDataInterface interface {
	GetByIDMidtrans(id string) (uint, error)
	GetAndUpdate(newData UpdatePayment, id string) (bool, error)
	InsertPaymentMaster(newData PaymentMaster) (*PaymentMaster, error)
	InsertPaymentDetail(newData PaymentDetail) (*PaymentDetail, error)
	GetQtyAndVoucByID(payment_id uint) (int, uint, error)
	GetTicketByID(payment_id uint) (uint, error)

	// Admin
	GetAll(status string, limitInt, offsetInt, timePublication int) ([]PaymentInfo, error)

	// User
	GetByUserID(id uint, status string, limitInt, offsetInt, timePublication int) ([]PaymentInfo, error)
}
