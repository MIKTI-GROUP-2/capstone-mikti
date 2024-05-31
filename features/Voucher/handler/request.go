package handler

import "time"

// CreateVoucherInput represents the input structure for creating a voucher.
type CreateVoucherInput struct {
	Code       string    json:"code" form:"code" validate:"required"
	Discount   float64   json:"discount" form:"discount" validate:"required"
	ExpiryDate time.Time json:"expiry_date" form:"expiry_date" validate:"required"
	EventID    int       json:"event_id" form:"event_id" validate:"required"
	Status     bool      json:"status" form:"status" validate:"required"
}

// UpdateVoucherInput represents the input structure for updating a voucher.
type UpdateVoucherInput struct {
	Code       string    json:"code" form:"code" validate:"required"
	Discount   float64   json:"discount" form:"discount" validate:"required"
	ExpiryDate time.Time json:"expiry_date" form:"expiry_date" validate:"required"
	EventID    int       json:"event_id" form:"event_id" validate:"required"
	Status     bool      json:"status" form:"status" validate:"required"
}

// GetVoucherByCodeInput represents the input structure for fetching a voucher by code.
type GetVoucherByCodeInput struct {
	Code string json:"code" form:"code" validate:"required"
}

// DeleteVoucherInput represents the input structure for deleting a voucher.
type DeleteVoucherInput struct {
	ID int json:"id" form:"id" validate:"required"
}