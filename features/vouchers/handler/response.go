package handler

// CreateVoucherInput represents the input structure for creating a voucher.
type InputResponse struct {
	Code       string `json:"code" form:"code" validate:"required"`
	Name       string `json:"name" form:"name" validate:"required"`
	Quantity   int    `json:"quantity" form:"quantity" validate:"required"`
	Price      int    `json:"price" form:"price" validate:"required"`
	ExpiryDate string `json:"expiry_date" form:"expiry_date" validate:"required"`
	EventID    uint   `json:"event_id" form:"event_id" validate:"required"`
	Status     bool   `json:"status" form:"status" validate:"required"`
}

// UpdateVoucherInput represents the input structure for updating a voucher.
type UpdateResponse struct {
	Code       string `json:"code" form:"code" validate:"required"`
	Name       string `json:"name" form:"name" validate:"required"`
	Quantity   int    `json:"quantity" form:"quantity" validate:"required"`
	Price      uint   `json:"price" form:"price" validate:"required"`
	ExpiryDate string `json:"expiry_date" form:"expiry_date" validate:"required"`
	EventID    int    `json:"event_id" form:"event_id" validate:"required"`
}
