package wishlists

import "github.com/labstack/echo/v4"

// Entity for Create
type Wishlist struct {
	ID      uint `json:"id"`
	UserID  int  `json:"user_id"`
	EventID int  `json:"event_id"`
}

// Entity for GetAll, GetByID
type WishlistInfo struct {
	ID         uint   `json:"id"`
	UserID     int    `json:"user_id"`
	EventID    int    `json:"event_id"`
	EventTitle string `json:"event_title"`
}

// Controller
type WishlistHandlerInterface interface {
	Create() echo.HandlerFunc
	GetAll() echo.HandlerFunc
	GetByID() echo.HandlerFunc
	Delete() echo.HandlerFunc
}

// Service
type WishlistServiceInterface interface {
	Create(user_id int, new_data Wishlist) (*Wishlist, error)
	GetAll(user_id int) ([]WishlistInfo, error)
	GetByID(user_id, id int) ([]WishlistInfo, error)
	Delete(user_id int, event_id int) error
}

// Repository
type WishlistDataInterface interface {
	CheckUnique(user_id, event_id int) ([]Wishlist, error)
	Create(user_id int, new_data Wishlist) (*Wishlist, error)
	GetAll(user_id int) ([]WishlistInfo, error)
	GetByID(user_id, id int) ([]WishlistInfo, error)
	Delete(user_id int, event_id int) error
}
