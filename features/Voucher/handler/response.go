package handler

// CreateVoucherResponse represents the response structure for creating a voucher.
type CreateVoucherResponse struct {
	Code       string  json:"code"
	Discount   float64 json:"discount"
	ExpiryDate string  json:"expiry_date"
	EventID    int     json:"event_id"
	Status     bool    json:"status"
}

// UpdateVoucherResponse represents the response structure for updating a voucher.
type UpdateVoucherResponse struct {
	Code       string  json:"code"
	Discount   float64 json:"discount"
	ExpiryDate string  json:"expiry_date"
	EventID    int     json:"event_id"
	Status     bool    json:"status"
}

// GetVoucherResponse represents the response structure for fetching a voucher.
type GetVoucherResponse struct {
	ID         int     json:"id"
	Code       string  json:"code"
	Discount   float64 json:"discount"
	ExpiryDate string  json:"expiry_date"
	EventID    int     json:"event_id"
	Status     bool    json:"status"
}

// DeleteVoucherResponse represents the response structure for deleting a voucher.
type DeleteVoucherResponse struct {
	Message string json:"message"
}