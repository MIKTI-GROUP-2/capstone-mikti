package data

import (
	"time"

	dataEvent "capstone-mikti/features/events/data"

	"gorm.io/gorm"
)

type Ticket struct {
	*gorm.Model
	EventID    int             `gorm:"column:event_id;type:int"`
	Name       string          `gorm:"column:name;type:varchar(225)"`
	TicketDate time.Time       `gorm:"column:ticket_date;type:date"`
	Quantity   int             `gorm:"column:quantity;type:int"`
	Price      int             `gorm:"column:price;type:int"`
	Event      dataEvent.Event `gorm:"foreignKey:EventID;references:ID"`
}
