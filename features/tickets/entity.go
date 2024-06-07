package tickets

import (
	"time"

	"github.com/labstack/echo/v4"
)

// Entity for CheckEvent
type Event struct {
	ID         uint   `json:"event_id"`
	EventTitle string `json:"event_title"`
}

// Entity for Create, Update
type Ticket struct {
	ID              uint      `json:"id"`
	EventID         int       `json:"event_id"`
	Name            string    `json:"name"`
	TicketDate      string    `json:"ticket_date"`
	ParseTicketDate time.Time `json:"-"`
	Quantity        int       `json:"quantity"`
	Price           int       `json:"price"`
}

// Entity for GetAll, GetByID
type TicketInfo struct {
	ID         uint   `json:"id"`
	EventID    int    `json:"event_id"`
	EventTitle string `json:"event_title"`
	Name       string `json:"name"`
	TicketDate string `json:"ticket_date"`
	Quantity   int    `json:"quantity"`
	Price      int    `json:"price"`
}

// Controller
type TicketHandlerInterface interface {
	Create() echo.HandlerFunc
	GetAll() echo.HandlerFunc
	GetByID() echo.HandlerFunc
	Update() echo.HandlerFunc
}

// Service
type TicketServiceInterface interface {
	Create(new_data Ticket) (*Ticket, error)
	GetAll() ([]TicketInfo, error)
	GetByID(id int) ([]TicketInfo, error)
	Update(id int, new_data Ticket) (bool, error)
}

// Repository
type TicketDataInterface interface {
	CheckEvent(event_id int) ([]Event, error)
	Create(new_data Ticket) (*Ticket, error)
	GetAll() ([]TicketInfo, error)
	GetByID(id int) ([]TicketInfo, error)
	Update(id int, new_data Ticket) (bool, error)
}
