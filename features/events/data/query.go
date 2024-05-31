package data

import (
	events "capstone-mikti/features/events"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type EventData struct {
	db *gorm.DB
}

func New(db *gorm.DB) *EventData {
	return &EventData{
		db: db,
	}
}

func (ed *EventData) GetByTitle(title string) ([]events.Event, error) {
	var dbData = []events.Event{}

	var qry = ed.db.Where("event_title = ?", title).First(&dbData)

	if err := qry.Error; err != nil {
		logrus.Error("DATA : Error Get By Title : ", err.Error())
		return nil, err
	}

	return dbData, nil
}

func (ed *EventData) CreateEvent(newData events.Event) (*events.Event, error) {
	var dbData = new(Event)

	dbData.CategoryFK = newData.CategoryFK
	dbData.Title = newData.Title
	dbData.City = newData.City
	dbData.Address = newData.Address
	dbData.StartingPrice = newData.StartingPrice
	dbData.StartDate = newData.ParseStartDate
	dbData.EndDate = newData.ParseEndDate
	dbData.Description = newData.Description
	dbData.Highlight = newData.Highlight
	dbData.ImportantInformation = newData.ImportantInformation
	dbData.Image = newData.Image

	if err := ed.db.Create(dbData).Error; err != nil {
		logrus.Error("DATA : Create Error : ", err.Error())
		return nil, err
	}

	return &newData, nil
}

func (e *EventData) GetAll() ([]events.AllEvent, error) {
	var listEvent = []events.AllEvent{}

	err := e.db.Table("events").Find(&listEvent).Error

	if err != nil {
		logrus.Error("Data : Get All Error : ", err.Error())
		return listEvent, err
	}

	return listEvent, nil
}

func (e *EventData) GetDetail(id int) ([]events.Event, error) {
	var event = []events.Event{}

	var query = e.db.Where("id = ? ", id).First(&event)

	if err := query.Error; err != nil {
		logrus.Error("DATA : Error Get By ID : ", err.Error())
		return nil, err
	}

	return event, nil

}
