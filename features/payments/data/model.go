package data

import (
	"time"

	"gorm.io/gorm"
)

type Payment struct {
	*gorm.Model
	BookID          uint          `gorm:"column:book_id"`
	UserID          uint          `gorm:"column:user_id"`
	OrderID         string        `gorm:"column:order_id"`
	MidtransID      string        `gorm:"column:midtrans_id"`
	GrandTotal      int           `gorm:"column:grand_total;type:int"`
	TransactionDate time.Time     `gorm:"column:transaction_date;type:date"`
	PaymentStatus   uint          `gorm:"column:payment_status"`
	PaymentType     string        `gorm:"column:payment_type;type:varchar(255)"`
	Status          string        `gorm:"column:status"`
	PaymentDetail   PaymentDetail `gorm:"foreignKey:PaymentID"`
}

type PaymentDetail struct {
	*gorm.Model
	PaymentID    uint `gorm:"column:payment_id"`
	VoucherID    uint `gorm:"column:voucher_id;default:null"`
	Price        int  `gorm:"column:price;type:int"`
	Qty          int  `gorm:"column:qty;type:int"`
	VoucherPrice int  `gorm:"column:voucher_price;type:int"`
	Total        int  `gorm:"column:total;type:int"`
}
