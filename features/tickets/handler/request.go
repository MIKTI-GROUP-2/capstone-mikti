package handler

type CreateTicketRequest struct {
	EventID    int    `json:"event_id" form:"event_id" validate:"required"`
	Name       string `json:"name" form:"name" validate:"required"`
	TicketDate string `json:"ticket_date" form:"ticket_date" validate:"required"`
	Quantity   int    `json:"quantity" form:"quantity" validate:"required"`
	Price      int    `json:"price" form:"price" validate:"required"`
}

type UpdateTicketRequest struct {
	EventID    int    `json:"event_id" form:"event_id"`
	Name       string `json:"name" form:"name"`
	TicketDate string `json:"ticket_date" form:"ticket_date"`
	Quantity   int    `json:"quantity" form:"quantity"`
	Price      int    `json:"price" form:"price"`
}
