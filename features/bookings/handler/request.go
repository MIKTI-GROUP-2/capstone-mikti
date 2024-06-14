package handler

type BookingRequest struct {
	TicketID    int    `json:"ticket_id" form:"ticket_id" validate:"required"`
	FullName    string `json:"full_name" form:"full_name" validate:"required"`
	PhoneNumber string `json:"phone_number" form:"phone_number" validate:"required"`
	Email       string `json:"email" form:"email" validate:"required"`
	Quantity    int    `json:"quantity" form:"quantity" validate:"required"`
}
