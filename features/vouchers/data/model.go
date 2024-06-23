package data

import (
	"capstone-mikti/features/payments/data"
	"time"

	"gorm.io/gorm"
)

type Voucher struct {
	*gorm.Model
	EventID       uint                 `gorm:"column:event_id"`
	Code          string               `gorm:"column:code;type:varchar(255);unique;not null"`
	Name          string               `gorm:"column:name;type:varchar(255);not null"`
	Quantity      int                  `gorm:"column:quantity;type:int;not null"`
	Price         int                  `gorm:"column:price;type:int;not null"`
	Description   string               `gorm:"column:description;type:varchar(255);not null"`
	ExpiredDate   time.Time            `gorm:"column:expired_date;type:date"`
	Status        bool                 `gorm:"column:status;type:bool"`
	PaymentDetail []data.PaymentDetail `gorm:"foreignKey:VoucherID"`
}
