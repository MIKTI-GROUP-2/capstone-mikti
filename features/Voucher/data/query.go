package data

import (
	"capstone-mikti/features/vouchers"
	"errors"
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type VoucherData struct {
	db *gorm.DB
}

func New(db *gorm.DB) *VoucherData {
	return &VoucherData{
		db: db,
	}
}

func (vd *VoucherData) CreateVoucher(newData vouchers.Voucher) (*vouchers.Voucher, error) {
	var dbData = new(Voucher)
	dbData.Code = newData.Code
	dbData.Discount = newData.Discount
	dbData.ExpiryDate = newData.ExpiryDate
	dbData.EventID = newData.EventID
	dbData.Status = newData.Status

	if err := vd.db.Create(dbData).Error; err != nil {
		logrus.Error("DATA : CreateVoucher Error : ", err.Error())
		return nil, err
	}

	return &newData, nil
}

func (vd *VoucherData) GetVoucherByID(id int) (vouchers.Voucher, error) {
	var voucher vouchers.Voucher
	var qry = vd.db.Table("vouchers").Select("vouchers.*").
		Where("vouchers.id = ?", id).
		Where("vouchers.status = ?", true).
		Scan(&voucher)

	if err := qry.Error; err != nil {
		logrus.Error("DATA : Error GetVoucherByID : ", err.Error())
		return voucher, err
	}

	return voucher, nil
}

func (vd *VoucherData) GetVoucherByCode(code string) (*vouchers.Voucher, error) {
	var dbData = new(Voucher)
	dbData.Code = code

	var qry = vd.db.Where("code = ?", dbData.Code).First(dbData)

	if err := qry.Error; err != nil {
		logrus.Error("DATA : Error GetVoucherByCode : ", err.Error())
		return nil, err
	}

	var result = new(vouchers.Voucher)
	result.ID = dbData.ID
	result.Code = dbData.Code
	result.Discount = dbData.Discount
	result.ExpiryDate = dbData.ExpiryDate
	result.EventID = dbData.EventID
	result.Status = dbData.Status

	return result, nil
}

func (vd *VoucherData) UpdateVoucher(id int, newData vouchers.UpdateVoucher) (bool, error) {
	var qry = vd.db.Table("vouchers").Where("id = ?", id).Updates(Voucher{
		Code:       newData.Code,
		Discount:   newData.Discount,
		ExpiryDate: newData.ExpiryDate,
		EventID:    newData.EventID,
		Status:     newData.Status,
	})

	if err := qry.Error; err != nil {
		logrus.Error("DATA : Error UpdateVoucher : ", err.Error())
		return false, err
	}

	if dataCount := qry.RowsAffected; dataCount < 1 {
		logrus.Error("DATA : No Row Affected")
		return false, nil
	}

	return true, nil
}

func (vd *VoucherData) DeleteVoucher(id int) error {
	if err := vd.db.Table("vouchers").Where("id = ?", id).Delete(&Voucher{}).Error; err != nil {
		logrus.Error("DATA : Error DeleteVoucher : ", err.Error())
		return err
	}

	return nil
}

type Voucher struct {
	ID         int       gorm:"primaryKey"
	Code       string    gorm:"not null"
	Discount   float64   gorm:"not null"
	ExpiryDate time.Time gorm:"not null"
	EventID    int       gorm:"not null"
	Status     bool      gorm:"not null"
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func (Voucher) TableName() string {
	return "vouchers"
}