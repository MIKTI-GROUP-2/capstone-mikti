package handler

type InputResponse struct {
	BookID          uint                   `json:"book_id" `
	Price           int                    `json:"price" `
	Qty             int                    `json:"qty" `
	GrandTotal      int                    `json:"grand_total" `
	TransactionDate string                 `json:"transaction_date" `
	CallbackURL     map[string]interface{} `json:"callback_url"`
}
