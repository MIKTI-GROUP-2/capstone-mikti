package events

import (
	"mime/multipart"
	"time"

	"github.com/labstack/echo/v4"
)

type TitleEvent struct {
	EventTitle string `json:"event_title"`
}

type PublicID struct {
	PublicID string `json:"public_id"`
}
type AllEvent struct {
	ID            uint   `json:"id"`
	CategoryID    int    `json:"category_id"`
	CategoryName  string `json:"category_name"`
	EventTitle    string `json:"event_title"`
	StartDate     string `json:"start_date"`
	EndDate       string `json:"end_date"`
	StartingPrice int    `json:"starting_price"`
	City          string `json:"city"`
}

type Event struct {
	ID                   uint           `json:"id"`
	CategoryID           int            `json:"category_id"`
	CategoryName         string         `json:"category_name"`
	EventTitle           string         `json:"event_title"`
	StartDate            string         `json:"start_date"`
	EndDate              string         `json:"end_date"`
	ParseStartDate       time.Time      `json:"-"`
	ParseEndDate         time.Time      `json:"-"`
	City                 string         `json:"city"`
	StartingPrice        int            `json:"starting_price"`
	Description          string         `json:"description"`
	Highlight            string         `json:"highlight"`
	ImportantInformation string         `json:"important_information"`
	Address              string         `json:"address"`
	ImageFile            multipart.File `json:"image_file"`
	ImageUrl             string         `json:"image_url"`
	PublicID             string         `json:"public_id"`
}

type EventHandlerInterface interface {
	CreateEvent() echo.HandlerFunc
	GetAll() echo.HandlerFunc
	GetDetail() echo.HandlerFunc
	UpdateEvent() echo.HandlerFunc
	DeleteEvent() echo.HandlerFunc
}

type EventServiceInterface interface {
	CreateEvent(newData Event) (*Event, error)
	GetAll(title string, category string, times string, city string, price int, sort string) ([]AllEvent, error)
	GetDetail(id int) ([]Event, error)
	UpdateEvent(id int, newData Event) (*Event, error)
	DeleteEvent(id int) (bool, error)
}

type EventDataInterface interface {
	CreateEvent(newData Event) (*Event, error)
	GetByTitle(username string) ([]TitleEvent, error)
	GetAll(title string, category string, times string, city string, price int, sort string) ([]AllEvent, error)
	GetDetail(id int) ([]Event, error)
	UpdateEvent(id int, newData Event) (*Event, error)
	DeleteEvent(id int) (bool, error)
	GetPublicID(id int) (string, error)
}
