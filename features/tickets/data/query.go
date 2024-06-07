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

// CheckEvent
func (td *TicketData) CheckEvent(event_id int) ([]tickets.Event, error) {
	// Get Entity
	var event = []tickets.Event{}

	// Query
	err := td.db.Table("events").
		Where("events.id = ?", event_id).
		Find(&event).Error

	// Error Handling
	if err != nil {
		logrus.Error("DATA : CheckEvent Error : ", err.Error())
		return nil, err
	}

	return event, nil
}

// Create
func (td *TicketData) Create(new_data tickets.Ticket) (*tickets.Ticket, error) {
	// Get Model
	ticket := new(Ticket)

	ticket.EventID = new_data.EventID
	ticket.Name = new_data.Name
	ticket.TicketDate = new_data.ParseTicketDate
	ticket.Quantity = new_data.Quantity
	ticket.Price = new_data.Price

	// Query
	err := td.db.Create(ticket).Error

	// Error Handling
	if err != nil {
		logrus.Error("DATA : Create Error : ", err.Error())
		return nil, err
	}

	// Parse Ticket Date
	new_data.TicketDate = new_data.ParseTicketDate.Format("2006-01-02")

	return &new_data, nil
}

// GetAll
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

// GetByID
func (td *TicketData) GetByID(id int) ([]tickets.TicketInfo, error) {
	// Get Entity
	var ticket = []tickets.TicketInfo{}

	// Query
	err := td.db.Table("tickets").
		Select("tickets.id, tickets.event_id, events.event_title, tickets.name, tickets.ticket_date, tickets.quantity, tickets.price").
		Joins("JOIN events on events.id = tickets.event_id").
		Where("tickets.id = ?", id).
		Find(&ticket).Error

	// Error Handling
	if err != nil {
		logrus.Error("DATA : GetAll Error : ", err.Error())
		return nil, err
	}

	return ticket, nil
}

// Update
func (td *TicketData) Update(id int, new_data tickets.Ticket) (bool, error) {
	// Query
	err := td.db.Table("tickets").
		Where("id = ?", id).
		Updates(Ticket{
			EventID:    new_data.EventID,
			Name:       new_data.Name,
			TicketDate: new_data.ParseTicketDate,
			Quantity:   new_data.Quantity,
			Price:      new_data.Price,
		})

	// Error Handling
	if err := err.Error; err != nil {
		logrus.Error("DATA : Error Update Ticket : ", err.Error())
		return false, err
	}

	if err.RowsAffected == 0 {
		logrus.Error("DATA : Record Not Found")
		return false, nil
	}

	new_data.TicketDate = new_data.ParseTicketDate.Format("2006-01-02")

	return true, nil
}
