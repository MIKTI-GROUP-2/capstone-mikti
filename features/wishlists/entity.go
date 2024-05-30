package wishlists

import "github.com/labstack/echo/v4"

type Wishlist struct {
	ID      uint `json:"id"`
	UserID  int  `json:"user_id"`
	EventID int  `json:"event_id"`
}

type WishlistInfo struct {
	ID         uint   `json:"id"`
	UserName   string `json:"username"`
	EventTitle string `json:"event_title"`
}

// Controller
type WishlistHandlerInterface interface {
	GetAll() echo.HandlerFunc
}

// Service
type WishlistServiceInterface interface {
	GetAll() ([]WishlistInfo, error)
}

// Repository
type WishlistDataInterface interface {
	GetAll() ([]WishlistInfo, error)
}
