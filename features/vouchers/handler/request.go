package handler

// CreateVoucherInput represents the input structure for creating a voucher.
type InputRequest struct {
	Code        string `json:"code" form:"code" validate:"required"`
	Name        string `json:"name" form:"name" validate:"required"`
	Quantity    int    `json:"quantity" form:"quantity" validate:"required"`
	Price       int    `json:"price" form:"price" validate:"required"`
	ExpiredDate string `json:"expired_date" form:"expired_date" validate:"required"`
	EventID     uint   `json:"event_id" form:"event_id" validate:"required"`
}

// UpdateVoucherInput represents the input structure for updating a voucher.
type UpdateRequest struct {
	Code        string `json:"code" form:"code"`
	Name        string `json:"name" form:"name"`
	Quantity    int    `json:"quantity" form:"quantity"`
	Price       int    `json:"price" form:"price"`
	ExpiredDate string `json:"expired_date" form:"expired_date"`
	EventID     uint   `json:"event_id" form:"event_id"`
}
