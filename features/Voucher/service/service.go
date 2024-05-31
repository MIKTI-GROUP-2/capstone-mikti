package service

import (
	"capstone-mikti/features/vouchers"
	"errors"
)

type VoucherService struct {
	data vouchers.VoucherDataInterface
}

func NewVoucherService(d vouchers.VoucherDataInterface) *VoucherService {
	return &VoucherService{
		data: d,
	}
}

func (vs *VoucherService) CreateVoucher(newData vouchers.Voucher) (*vouchers.Voucher, error) {
	result, err := vs.data.CreateVoucher(newData)
	if err != nil {
		return nil, errors.New("ERROR Create Voucher")
	}
	return result, nil
}

func (vs *VoucherService) GetVoucherByID(id int) (*vouchers.Voucher, error) {
	result, err := vs.data.GetVoucherByID(id)
	if err != nil {
		return nil, errors.New("ERROR Get Voucher By ID")
	}
	return result, nil
}

func (vs *VoucherService) GetVoucherByCode(code string) (*vouchers.Voucher, error) {
	result, err := vs.data.GetVoucherByCode(code)
	if err != nil {
		return nil, errors.New("ERROR Get Voucher By Code")
	}
	return result, nil
}

func (vs *VoucherService) UpdateVoucher(id int, newData vouchers.UpdateVoucher) (bool, error) {
	result, err := vs.data.UpdateVoucher(id, newData)
	if err != nil {
		SSs
		return false, errors.New("ERROR Update Voucher")
	}
	return result, nil
}

func (vs *VoucherService) DeleteVoucher(id int) error {
	err := vs.data.DeleteVoucher(id)
	if err != nil {
		return errors.New("ERROR Delete Voucher")
	}
	return nil
}
