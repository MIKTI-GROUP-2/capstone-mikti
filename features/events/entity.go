package events

import (
	"capstone-mikti/helper/customtime"

	"github.com/labstack/echo/v4"
)

type Event struct {
	ID                   uint                  `json:"id"`
	CategoryFK           int                   `json:"category_id"`
	Title                string                `json:"event_title"`
	StartDate            customtime.CustomTime `json:"start_date"`
	EndDate              customtime.CustomTime `json:"end_date"`
	City                 string                `json:"city"`
	StartingPrice        int                   `json:"starting_price"`
	Description          string                `json:"description"`
	Highlight            string                `json:"highlight"`
	ImportantInformation string                `json:"important_information"`
	Address              string                `json:"address"`
	Image                string                `json:"image_url"`
}

type EventHandlerInterface interface {
	CreateEvent() echo.HandlerFunc
	// GetEvent() echo.HandlerFunc
}

type EventServiceInterface interface {
	CreateEvent(newData Event) (*Event, error)
	// GetEvent() ([]Event, error)
}

type EventDataInterface interface {
	CreateEvent(newData Event) (*Event, error)
	GetByTitle(username string) (*Event, error)
	// GetEvent() ([]Event, error)
}
