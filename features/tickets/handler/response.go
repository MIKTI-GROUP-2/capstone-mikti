package handler

type TicketResponse struct {
	EventID    int    `json:"event_id"`
	Name       string `json:"name"`
	TicketDate string `json:"ticket_date"`
	Quantity   int    `json:"quantity"`
	Price      int    `json:"price"`
}
