package wishlists

import (
	"github.com/labstack/echo/v4"
)

// Entity for CheckEvent
type Event struct {
	ID uint `json:"event_id"`
}

// Entity for Create
type Wishlist struct {
	ID      uint `json:"id"`
	UserID  int  `json:"user_id"`
	EventID int  `json:"event_id"`
}

// Entity for GetAll, GetByEventID
type WishlistInfo struct {
	ID                   uint   `json:"id"`
	EventID              int    `json:"event_id"`
	CategoryID           int    `json:"category_id"`
	CategoryName         string `json:"category_name"`
	EventTitle           string `json:"event_title"`
	StartDate            string `json:"start_date"`
	EndDate              string `json:"end_date"`
	City                 string `json:"city"`
	StartingPrice        int    `json:"starting_price"`
	Description          string `json:"description"`
	Highlight            string `json:"highlight"`
	ImportantInformation string `json:"important_information"`
	Address              string `json:"address"`
	ImageUrl             string `json:"image_url"`
}

// Controller
type WishlistHandlerInterface interface {
	Create() echo.HandlerFunc
	GetAll() echo.HandlerFunc
	GetByEventID() echo.HandlerFunc
	Delete() echo.HandlerFunc
}

// Service
type WishlistServiceInterface interface {
	Create(user_id int, new_data Wishlist) (*Wishlist, error)
	GetAll(user_id int) ([]WishlistInfo, error)
	GetByEventID(user_id int, event_id int) ([]WishlistInfo, error)
	Delete(user_id int, event_id int) (bool, error)
}

// Repository
type WishlistDataInterface interface {
	CheckEvent(event_id int) ([]Event, error)
	CheckUnique(user_id int, event_id int) ([]Wishlist, error)
	Create(user_id int, new_data Wishlist) (*Wishlist, error)
	GetAll(user_id int) ([]WishlistInfo, error)
	GetByEventID(user_id int, event_id int) ([]WishlistInfo, error)
	Delete(user_id int, event_id int) (bool, error)
}
