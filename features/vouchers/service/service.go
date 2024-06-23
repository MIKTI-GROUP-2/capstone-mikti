package service

import (
	"capstone-mikti/features/vouchers"
	"errors"
	"time"

	"github.com/sirupsen/logrus"
)

type VoucherService struct {
	data vouchers.VoucherDataInterface
}

func New(d vouchers.VoucherDataInterface) *VoucherService {
	return &VoucherService{
		data: d,
	}
}

func (vs *VoucherService) CreateVoucher(newData vouchers.Voucher) (*vouchers.Voucher, error) {
	layout := "2006-01-02"

	parsedExpiredDate, _ := time.Parse(layout, newData.ExpiryDate)

	newData.ExpiryDateParse = parsedExpiredDate
	newData.Status = true
	result, err := vs.data.Create(newData)
	if err != nil {
		logrus.Error("Service Error : ", err.Error())
		return nil, errors.New("ERROR Create Voucher")
	}
	return result, nil
}

func (vs *VoucherService) GetVoucher(id int) ([]vouchers.VoucherInfo, error) {
	result, err := vs.data.GetByID(id)
	if err != nil {
		logrus.Error("Service Error : ", err.Error())
		return nil, errors.New("ERROR Get Voucher By ID")
	}
	return result, nil
}

func (vs *VoucherService) GetVouchers() ([]vouchers.VoucherInfo, error) {
	result, err := vs.data.GetAll()
	if err != nil {
		logrus.Error("Service Error : ", err.Error())
		return nil, errors.New("ERROR Get All")
	}
	return result, nil
}

func (vs *VoucherService) GetVoucherByCode(code string) (*vouchers.VoucherInfo, error) {
	result, err := vs.data.GetByCode(code)
	if err != nil {
		logrus.Error("Service Error : ", err.Error())
		return nil, errors.New("ERROR Get Voucher By Code")
	}
	return result, nil
}

func (vs *VoucherService) UpdateVoucher(id int, newData vouchers.UpdateVoucher) (bool, error) {
	result, err := vs.data.Update(id, newData)
	if err != nil {
		logrus.Error("Service Error : ", err.Error())
		return false, errors.New("ERROR Update Voucher")
	}
	return result, nil
}

func (vs *VoucherService) ActivateVoucher(id int) (bool, error) {
	res, err := vs.data.Activate(id)
	if err != nil {
		logrus.Error("Service Error : ", err.Error())
		return res, errors.New("ERROR Activate Voucher")
	}
	return res, nil
}

func (vs *VoucherService) DeactivateVoucher(id int) (bool, error) {
	res, err := vs.data.Deactivate(id)
	if err != nil {
		logrus.Error("Service Error : ", err.Error())
		return res, errors.New("ERROR Deactivate Voucher")
	}
	return res, nil
}
