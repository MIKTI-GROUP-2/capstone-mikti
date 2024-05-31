package vouchers

import (
	"time"
)

type Voucherpackage vouchers

import (
	"time"
)

type Voucher struct {
	ID         uint      json:"id"
	Code       string    json:"code"
	Discount   float64   json:"discount"
	ExpiryDate time.Time json:"expiry_date"
	EventID    uint      json:"event_id"
	Status     bool      json:"status"
}

type VoucherHandlerInterface interface {
	CreateVoucher() echo.HandlerFunc
	GetVoucherByID() echo.HandlerFunc
	UpdateVoucher() echo.HandlerFunc
	DeleteVoucher() echo.HandlerFunc
}

type VoucherServiceInterface interface {
	CreateVoucher(newData Voucher) (*Voucher, error)
	GetVoucherByID(id int) (*Voucher, error)
	UpdateVoucher(id int, newData UpdateVoucher) (bool, error)
	DeleteVoucher(id int) error
}

type VoucherDataInterface interface {
	CreateVoucher(newData Voucher) (*Voucher, error)
	GetVoucherByID(id int) (*Voucher, error)
	UpdateVoucher(id int, newData UpdateVoucher) (bool, error)
	DeleteVoucher(id int) error
}struct {
	ID         uint      json:"id"
	Code       string    json:"code"
	Discount   float64   json:"discount"
	ExpiryDate time.Time json:"expiry_date"
	EventID    uint      json:"event_id"
	Status     bool      json:"status"
}

type VoucherHandlerInterface interface {
	CreateVoucher() echo.HandlerFunc
	GetVoucherByID() echo.HandlerFunc
	UpdateVoucher() echo.HandlerFunc
	DeleteVoucher() echo.HandlerFunc
}

type VoucherServiceInterface interface {
	CreateVoucher(newData Voucher) (*Voucher, error)
	GetVoucherByID(id int) (*Voucher, error)
	UpdateVoucher(id int, newData UpdateVoucher) (bool, error)
	DeleteVoucher(id int) error
}

type VoucherDataInterface interface {
	CreateVoucher(newData Voucher) (*Voucher, error)
	GetVoucherByID(id int) (*Voucher, error)
	UpdateVoucher(id int, newData UpdateVoucher) (bool, error)
	DeleteVoucher(id int) error
}