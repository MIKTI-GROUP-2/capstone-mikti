package data

import (
	"capstone-mikti/features/tickets"
	"errors"
	"time"

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
func (td *TicketData) CheckEvent(event_id int, ticket_date time.Time) (bool, error) {
	// Get Entity
	event := []tickets.Event{}

	// Query to Validate Event ID is Stored in Table Event
	err := td.db.Table("events").
		Where("events.id = ?", event_id).
		Where("events.deleted_at is null").
		Scan(&event).Error

	if err != nil {
		logrus.Error("Data : CheckEvent Error : ", err.Error())
		return false, err
	}

	if len(event) == 0 {
		logrus.Warn("Data : CheckEvent Warning")
		return false, errors.New("WARNING Event ID is Not Found")
	}
	// Validate Ticket Date is in Range of Event Start Date & End Date
	start_date := event[0].StartDate
	end_date := event[0].EndDate

	isCheck := validateTicketDate(start_date, end_date, ticket_date)

	if !isCheck {
		logrus.Warn("Data: CheckEvent Warning")
		return false, errors.New("WARNING Ticket Date is Out of Event Date Range")
	}

	return true, nil
}

func validateTicketDate(start_date, end_date, ticket_date time.Time) bool {
	if ticket_date.Before(start_date) || ticket_date.After(end_date) {
		return false
	}

	return true
}

// CheckTicketDate
func (td *TicketData) CheckTicketDate(event_id int, ticket_date time.Time) (bool, error) {
	// Query
	var count int64

	err := td.db.Table("tickets").
		Where("tickets.event_id = ?", event_id).
		Where("tickets.ticket_date = ?", ticket_date).
		Where("tickets.deleted_at is null").
		Count(&count).Error

	if err != nil {
		logrus.Error("Data : CheckTicketDate Error : ", err.Error())
		return false, err
	}

	if count > 0 {
		return false, errors.New("WARNING Duplicate Ticket Date on the same Event ID")
	}

	return true, nil
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

	if err != nil {
		logrus.Error("Data : Create Error : ", err.Error())
		return nil, err
	}

	// Parse Ticket Date
	new_data.TicketDate = new_data.ParseTicketDate.Format("2006-01-02")

	return &new_data, nil
}

// GetAll
func (td *TicketData) GetAll() ([]tickets.TicketInfo, error) {
	// Get Entity
	ticket := []tickets.TicketInfo{}

	// Query
	err := td.db.Table("tickets").
		Select("tickets.*, events.event_title").
		Joins("JOIN events on events.id = tickets.event_id").
		Where("tickets.deleted_at is null").
		Scan(&ticket).Error

	if err != nil {
		logrus.Error("Data : GetAll Error : ", err.Error())
		return nil, err
	}

	return ticket, nil
}

// GetByID
func (td *TicketData) GetByID(id int) ([]tickets.TicketInfo, error) {
	// Get Entity
	ticket := []tickets.TicketInfo{}

	// Query
	err := td.db.Table("tickets").
		Select("tickets.*, events.event_title").
		Joins("JOIN events on events.id = tickets.event_id").
		Where("tickets.id = ?", id).
		Where("tickets.deleted_at is null").
		Scan(&ticket).Error

	if err != nil {
		logrus.Error("Data : GetByID Error : ", err.Error())
		return nil, err
	}

	return ticket, nil
}

// GetByEventID
func (td *TicketData) GetByEventID(event_id int) ([]tickets.TicketInfo, error) {
	// Get Entity
	ticket := []tickets.TicketInfo{}

	// Query
	err := td.db.Table("tickets").
		Select("tickets.*, events.event_title").
		Joins("JOIN events on events.id = tickets.event_id").
		Where("tickets.event_id = ?", event_id).
		Where("tickets.deleted_at is null").
		Scan(&ticket).Error

	if err != nil {
		logrus.Error("Data : GetByEventID Error : ", err.Error())
		return nil, err
	}

	return ticket, nil
}

// Update
func (td *TicketData) Update(id int, new_data tickets.Ticket) (bool, error) {
	// Query
	err := td.db.Table("tickets").
		Where("id = ?", id).
		Where("tickets.deleted_at is null").
		Updates(Ticket{
			EventID:    new_data.EventID,
			Name:       new_data.Name,
			TicketDate: new_data.ParseTicketDate,
			Quantity:   new_data.Quantity,
			Price:      new_data.Price,
		})

	if err.Error != nil {
		logrus.Error("Data : Update Error : ", err.Error)
		return false, err.Error
	}

	if err.RowsAffected == 0 {
		logrus.Warn("Data : Update Warning")
		return false, errors.New("WARNING No Rows Affected")
	}

	// Parse Ticket Date
	new_data.TicketDate = new_data.ParseTicketDate.Format("2006-01-02")

	return true, nil
}

func (td *TicketData) UpdateQtyAndPrice(id int, new_data tickets.Ticket) (bool, error) {
	// Query
	err := td.db.Table("tickets").
		Where("id = ?", id).
		Where("tickets.deleted_at is null").
		Updates(Ticket{
			Quantity: new_data.Quantity,
			Price:    new_data.Price,
		})

	if err.Error != nil {
		logrus.Error("Data : Update Error : ", err.Error)
		return false, err.Error
	}

	if err.RowsAffected == 0 {
		logrus.Warn("Data : Update Warning")
		return false, errors.New("WARNING No Rows Affected")
	}

	return true, nil
}

// Delete
func (td *TicketData) Delete(id int) (bool, error) {
	// Get Model
	ticket := new(Ticket)

	// Query
	err := td.db.Where("id = ?", id).Delete(&ticket)

	if err.Error != nil {
		logrus.Error("Data : Delete Error : ", err.Error)
		return false, err.Error
	}

	if err.RowsAffected == 0 {
		logrus.Warn("Data : Delete Warning")
		return false, errors.New("WARNING No Rows Affected")
	}

	return true, nil
}
