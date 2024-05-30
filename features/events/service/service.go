package service

import (
	"capstone-mikti/features/events"
	"capstone-mikti/helper/jwt"
	"errors"

	"github.com/sirupsen/logrus"
)

type EventService struct {
	data events.EventDataInterface
	jwt  jwt.JWTInterface
}

func New(d events.EventDataInterface, j jwt.JWTInterface) *EventService {
	return &EventService{
		data: d,
		jwt:  j,
	}
}

func (e *EventService) CreateEvent(newData events.Event) (*events.Event, error) {
	_, err := e.data.GetByTitle(newData.Title)

	if err == nil {
		logrus.Error("Service : Titke already registered")
		return nil, errors.New("ERROR Title already registered by another user")
	}

	result, err := e.data.CreateEvent(newData)
	if err != nil {
		logrus.Error("Service : Error Create : ", err.Error())
		return nil, errors.New("ERROR Create")
	}

	return result, nil
}

// func (e *EventService) GetEvent() (*[]events.Event, error) {

// 	res, err := e.data.GetEvent()

// 	if err != nil {
// 		logrus.Error("Service : Get Event Error : ", err.Error())
// 		return nil, errors.New("ERROR Get Event")
// 	}

// 	return &res, nil

// }
