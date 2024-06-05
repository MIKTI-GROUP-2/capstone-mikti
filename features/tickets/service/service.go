package service

import (
	"capstone-mikti/features/tickets"
	"errors"
	"time"

	"github.com/sirupsen/logrus"
)

// Service

type TicketService struct {
	data tickets.TicketDataInterface
}

func New(d tickets.TicketDataInterface) *TicketService {
	return &TicketService{
		data: d,
	}
}

// Create
func (ts *TicketService) Create(new_data tickets.Ticket) (*tickets.Ticket, error) {
	// Parse Ticket Date
	layout := "2006-01-02"

	parse_ticketDate, err := time.Parse(layout, new_data.TicketDate)

	if err != nil {
		logrus.Error("Service : Parse Ticket Date Error : ", err.Error())
		return nil, errors.New("ERROR Parse Ticket Date")
	}

	new_data.ParseTicketDate = parse_ticketDate

	// Get Data
	create, err := ts.data.Create(new_data)

	if err != nil {
		logrus.Error("Service : Create Error : ", err.Error())
		return nil, errors.New("ERROR Create")
	}

	return create, nil
}

// GetAll
func (ts *TicketService) GetAll() ([]tickets.TicketInfo, error) {
	// Get Data
	getAll, err := ts.data.GetAll()

	if err != nil {
		logrus.Error("Service : GetAll Error : ", err.Error())
		return nil, errors.New("ERROR GetAll")
	}

	return getAll, nil
}

// GetByID
func (ts *TicketService) GetByID(id int) ([]tickets.TicketInfo, error) {
	// Get Data
	getAll, err := ts.data.GetByID(id)

	if err != nil {
		logrus.Error("Service : GetAll Error : ", err.Error())
		return nil, errors.New("ERROR GetAll")
	}

	return getAll, nil
}
