package handler

type BookingResponse struct {
	ID uint `json:"id"`

	FullName    string `json:"full_name"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
	BookingDate string `json:"booking_date"`
	IsBooked    bool   `json:"is_booked"`
	IsPaid      bool   `json:"is_paid"`
	Quantity    int    `json:"quantity"`
	Total       int    `json:"total"`
	Event       EventResponse
	Ticket      TicketResponse
	User        UserResponse
}

type TicketResponse struct {
	TicketID int    `json:"ticket_id"`
	Name     string `json:"name"`
	Price    int    `json:"price"`
}

type EventResponse struct {
	EventID    int    `json:"event_id"`
	EventTitle string `json:"event_title"`
}
type UserResponse struct {
	UserID   int    `json:"user_id"`
	Username string `json:"username"`
}

type AllBookingResponse struct {
	ID          uint   `json:"id"`
	UserID      int    `json:"user_id"`
	TicketID    int    `json:"ticket_id"`
	FullName    string `json:"full_name"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
	BookingDate string `json:"booking_date"`
	IsBooked    bool   `json:"is_booked"`
	IsPaid      bool   `json:"is_paid"`
	Quantity    int    `json:"quantity"`
}
