package handler

type InputRequest struct {
	BookID          uint   `json:"book_id" validate:"required"`
	VoucherID       uint   `json:"voucher_id"`
	VoucherCode     string `json:"voucher_code"`
	Price           int    `json:"price" validate:"required"`
	Qty             int    `json:"qty" validate:"required"`
	TransactionDate string `json:"transaction_date" validate:"required"`
	PaymentType     string `json:"payment_type" validate:"required"`
}
