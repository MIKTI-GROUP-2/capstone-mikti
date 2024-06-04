package data

import (
	"capstone-mikti/features/tickets"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// Repository

type TicketData struct {
	db *gorm.DB
}

func New(db *gorm.DB) *TicketData {
	return &TicketData{
		db: db,
	}
}

// Create
func (td *TicketData) GetAll() ([]tickets.TicketInfo, error) {
	// Get Entity
	var ticket = []tickets.TicketInfo{}

	// Query
	err := td.db.Table("tickets").
		Select("tickets.id, tickets.event_id, events.event_title, tickets.name, tickets.ticket_date, tickets.quantity, tickets.price").
		Joins("JOIN events on events.id = tickets.event_id").
		Find(&ticket).Error

	// Error Handling
	if err != nil {
		logrus.Error("DATA : GetAll Error : ", err.Error())
		return nil, err
	}

	return ticket, nil
}
