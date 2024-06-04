package service

import (
	"capstone-mikti/features/tickets"
	"errors"

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

// GetAll
func (ts *TicketService) GetAll() ([]tickets.TicketInfo, error) {
	// Get Data
	getAll, err := ts.data.GetAll()

	// Error Handling
	if err != nil {
		logrus.Error("Service : GetAll Error : ", err.Error())
		return nil, errors.New("ERROR GetAll")
	}

	return getAll, nil
}
