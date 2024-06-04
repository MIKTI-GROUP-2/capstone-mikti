package tickets

import (
	"time"

	"github.com/labstack/echo/v4"
)

// Entity for Create, Update
type Ticket struct {
	ID         uint      `json:"id"`
	EventID    int       `json:"event_id"`
	Name       string    `json:"name"`
	TicketDate time.Time `json:"ticket_date"`
	Quantity   int       `json:"quantity"`
	Price      int       `json:"price"`
}

// Entity for GetAll, GetByID
type TicketInfo struct {
	ID         uint      `json:"id"`
	EventID    int       `json:"event_id"`
	EventTitle string    `json:"event_title"`
	Name       string    `json:"name"`
	TicketDate time.Time `json:"ticket_date"`
	Quantity   int       `json:"quantity"`
	Price      int       `json:"price"`
}

// Controller
type TicketHandlerInterface interface {
	GetAll() echo.HandlerFunc
}

// Service
type TicketServiceInterface interface {
	GetAll() ([]TicketInfo, error)
}

// Repository
type TicketDataInterface interface {
	GetAll() ([]TicketInfo, error)
}
