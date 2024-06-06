package vouchers

import (
	"time"

	"github.com/labstack/echo/v4"
)

type Voucher struct {
	ID              uint      `json:"id"`
	Code            string    `json:"code"`
	Name            string    `json:"name"`
	Quantity        int       `json:"quantity"`
	Price           int       `json:"price"`
	ExpiryDateParse time.Time `json:"expired_date"`
	ExpiryDate      string    `json:"expired_date_parse"`
	EventID         uint      `json:"event_id"`
	Status          bool      `json:"status"`
}

type UpdateVoucher struct {
	Code            string    `json:"code"`
	Name            string    `json:"name"`
	Quantity        int       `json:"quantity"`
	Price           int       `json:"price"`
	ExpiryDateParse time.Time `json:"expired_date"`
	ExpiryDate      string    `json:"expired_date_parse"`
	EventID         uint      `json:"event_id"`
	Status          bool      `json:"status"`
}

type VoucherInfo struct {
	Code        string `json:"code"`
	Name        string `json:"name"`
	Quantity    int    `json:"quantity"`
	Price       int    `json:"price"`
	ExpiredDate string `json:"expired_date"`
	EventTitle  string `json:"event_title"`
	Status      bool   `json:"status"`
}

type VoucherHandlerInterface interface {
	GetVouchers() echo.HandlerFunc
	GetVoucher() echo.HandlerFunc
	GetVoucherByCode() echo.HandlerFunc
	CreateVoucher() echo.HandlerFunc
	UpdateVoucher() echo.HandlerFunc
	ActivateVoucher() echo.HandlerFunc
	DeactivateVoucher() echo.HandlerFunc
}

type VoucherServiceInterface interface {
	GetVouchers() ([]VoucherInfo, error)
	GetVoucher(id int) ([]VoucherInfo, error)
	GetVoucherByCode(code string) ([]VoucherInfo, error)
	CreateVoucher(newData Voucher) (*Voucher, error)
	UpdateVoucher(id int, newData UpdateVoucher) (bool, error)
	ActivateVoucher(id int) (bool, error)
	DeactivateVoucher(id int) (bool, error)
}

type VoucherDataInterface interface {
	GetAll() ([]VoucherInfo, error)
	GetByID(id int) ([]VoucherInfo, error)
	GetByCode(code string) ([]VoucherInfo, error)
	Create(newData Voucher) (*Voucher, error)
	Update(id int, newData UpdateVoucher) (bool, error)
	Activate(id int) (bool, error)
	Deactivate(id int) (bool, error)
}
