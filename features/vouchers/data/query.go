package data

import (
	"capstone-mikti/features/vouchers"

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

func (vd *VoucherData) Create(newData vouchers.Voucher) (*vouchers.Voucher, error) {
	var dbData = new(Voucher)
	dbData.Code = newData.Code
	dbData.Name = newData.Name
	dbData.Quantity = newData.Quantity
	dbData.Price = newData.Price
	dbData.ExpiredDate = newData.ExpiryDateParse
	dbData.EventID = newData.EventID
	dbData.Status = newData.Status

	if err := vd.db.Create(dbData).Error; err != nil {
		logrus.Error("DATA : Create Error : ", err.Error())
		return nil, err
	}

	return &newData, nil
}

func (vd *VoucherData) GetAll() ([]vouchers.VoucherInfo, error) {
	var voucher = []vouchers.VoucherInfo{}

	var qry = vd.db.Table("vouchers").Select("vouchers.*", "events.event_title").
		Joins("LEFT JOIN events ON events.id = vouchers.event_id").
		Where("vouchers.status = ?", true).
		Scan(&voucher)

	if err := qry.Error; err != nil {
		logrus.Error("DATA : Error GetVoucherByID : ", err.Error())
		return voucher, err
	}

	return voucher, nil
}

func (vd *VoucherData) GetByID(id int) ([]vouchers.VoucherInfo, error) {
	var voucher = []vouchers.VoucherInfo{}

	var qry = vd.db.Table("vouchers").Select("vouchers.*", "events.event_title").
		Joins("LEFT JOIN events ON events.id = vouchers.event_id").
		Where("vouchers.id = ?", id).
		Where("vouchers.status = ?", true).
		Scan(&voucher)

	if err := qry.Error; err != nil {
		logrus.Error("DATA : Error GetVoucherByID : ", err.Error())
		return voucher, err
	}

	return voucher, nil
}

func (vd *VoucherData) GetByCode(code string) ([]vouchers.VoucherInfo, error) {
	var voucher = []vouchers.VoucherInfo{}

	var qry = vd.db.Table("vouchers").Select("vouchers.*", "events.event_title").
		Joins("LEFT JOIN events ON events.id = vouchers.event_id").
		Where("vouchers.code = ?", code).
		Where("vouchers.status = ?", true).
		Scan(&voucher)

	if err := qry.Error; err != nil {
		logrus.Error("DATA : Error GetVoucherByCode : ", err.Error())
		return nil, err
	}

	return voucher, nil
}

func (vd *VoucherData) Update(id int, newData vouchers.UpdateVoucher) (bool, error) {
	var qry = vd.db.Table("vouchers").Where("id = ?", id).Updates(Voucher{
		Code:        newData.Code,
		Name:        newData.Name,
		Quantity:    newData.Quantity,
		Price:       newData.Price,
		ExpiredDate: newData.ExpiryDateParse,
		EventID:     newData.EventID,
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

func (v *VoucherData) Deactivate(id int) (bool, error) {
	var qry = v.db.Model(&Voucher{}).Where("id = ?", id).Updates(map[string]interface{}{
		"Status": false,
	})

	if err := qry.Error; err != nil {
		return false, err
	}

	return true, nil
}

func (v *VoucherData) Activate(id int) (bool, error) {
	var qry = v.db.Table("vouchers").Where("id = ?", id).Updates(Voucher{Status: true})

	if err := qry.Error; err != nil {
		return false, err
	}

	return true, nil
}
