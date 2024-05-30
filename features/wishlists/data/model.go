package data

import (
	dataEvent "capstone-mikti/features/events/data"
	"capstone-mikti/features/users/data"

	"gorm.io/gorm"
)

type Wishlist struct {
	*gorm.Model
	UserID  int             `gorm:"column:user_id;type:int"`
	EventID int             `gorm:"column:event_id;type:int"`
	User    data.User       `gorm:"foreignKey:UserID;references:ID"`
	Event   dataEvent.Event `gorm:"foreignKey:UserID;references:ID"`
}
