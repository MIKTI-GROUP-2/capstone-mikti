package events

import (
	"time"

	"github.com/labstack/echo/v4"
)

type AllEvent struct {
	ID             uint      `json:"id"`
	CategoryID     int       `json:"category_id"`
	CategoryName   string    `json:"category_name"`
	EventTitle     string    `json:"event_title"`
	StartDate      string    `json:"start_date"`
	EndDate        string    `json:"end_date"`
	ParseStartDate time.Time `json:"-"`
	ParseEndDate   time.Time `json:"-"`
	StartingPrice  int       `json:"starting_price"`
	City           string    `json:"city"`
	EventName      string    `json:"event_name"`
}

type Event struct {
	ID                   uint      `json:"id"`
	CategoryID           int       `json:"category_id"`
	CategoryName         string    `json:"category_name"`
	EventTitle           string    `json:"event_title"`
	StartDate            string    `json:"start_date"`
	EndDate              string    `json:"end_date"`
	ParseStartDate       time.Time `json:"-"`
	ParseEndDate         time.Time `json:"-"`
	City                 string    `json:"city"`
	StartingPrice        int       `json:"starting_price"`
	Description          string    `json:"description"`
	Highlight            string    `json:"highlight"`
	ImportantInformation string    `json:"important_information"`
	Address              string    `json:"address"`
	ImageUrl             string    `json:"image_url"`
	PublicID             string    `json:"public_id"`
}

type EventHandlerInterface interface {
	CreateEvent() echo.HandlerFunc
	GetAll() echo.HandlerFunc
	GetDetail() echo.HandlerFunc
	UpdateEvent() echo.HandlerFunc
	DeleteEvent() echo.HandlerFunc
	// FilterEvent() echo.HandlerFunc
}

type EventServiceInterface interface {
	CreateEvent(newData Event) (*Event, error)
	GetAll() ([]AllEvent, error)
	GetDetail(id int) ([]Event, error)
	UpdateEvent(id int, newData Event) (*Event, error)
	DeleteEvent(id int) (bool, error)
	// FilterEvent(data string) ([]AllEvent, error)
}

type EventDataInterface interface {
	CreateEvent(newData Event) (*Event, error)
	GetByTitle(username string) ([]Event, error)
	GetAll() ([]AllEvent, error)
	GetDetail(id int) ([]Event, error)
	UpdateEvent(id int, newData Event) (*Event, error)
	DeleteEvent(id int) (bool, error)
	// FilterEvent(data string) ([]AllEvent, error)
}
