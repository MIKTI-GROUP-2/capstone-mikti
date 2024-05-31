package wishlists

import "github.com/labstack/echo/v4"

// Entity for Create
type Wishlist struct {
	ID      uint `json:"id"`
	UserID  int  `json:"user_id"`
	EventID int  `json:"event_id"`
}

// Entity for GetAll, GetByUserID
type WishlistInfo struct {
	ID         uint   `json:"id"`
	UserID     int    `json:"user_id"`
	UserName   string `json:"username"`
	EventID    int    `json:"event_id"`
	EventTitle string `json:"event_title"`
}

// Controller
type WishlistHandlerInterface interface {
	Create() echo.HandlerFunc
	GetAll() echo.HandlerFunc
	GetByUserID() echo.HandlerFunc
	Delete() echo.HandlerFunc
}

// Service
type WishlistServiceInterface interface {
	Create(new_data Wishlist) (*Wishlist, error)
	GetAll() ([]WishlistInfo, error)
	GetByUserID(user_id int) ([]WishlistInfo, error)
	Delete(id uint) error
}

// Repository
type WishlistDataInterface interface {
	Create(new_data Wishlist) (*Wishlist, error)
	GetAll() ([]WishlistInfo, error)
	GetByUserID(user_id int) ([]WishlistInfo, error)
	Delete(id uint) error
}
