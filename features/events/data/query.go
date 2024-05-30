package data

import (
	"capstone-mikti/features/events"
	"capstone-mikti/helper/customtime"

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

func (ed *EventData) GetByTitle(title string) (*events.Event, error) {
	var dbData = new(Event)
	dbData.Title = title

	var qry = ed.db.Where("event_title = ?", dbData.Title).First(dbData)

	if err := qry.Error; err != nil {
		logrus.Error("DATA : Error Get By Title : ", err.Error())
		return nil, err
	}

	var result = new(events.Event)
	result.ID = dbData.ID
	result.CategoryFK = dbData.CategoryFK
	result.Title = dbData.Title
	result.City = dbData.City
	result.Address = dbData.Address
	result.StartDate = customtime.CustomTime{Time: dbData.StartDate}
	result.EndDate = customtime.CustomTime{Time: dbData.EndDate}
	result.StartingPrice = dbData.StartingPrice
	result.Description = dbData.Description
	result.Highlight = dbData.Highlight
	result.ImportantInformation = dbData.ImportantInformation
	result.Image = dbData.Image

	return result, nil
}

func (ed *EventData) CreateEvent(newData events.Event) (*events.Event, error) {
	var dbData = new(Event)

	dbData.CategoryFK = newData.CategoryFK
	dbData.Title = newData.Title
	dbData.City = newData.City
	dbData.Address = newData.Address
	dbData.StartingPrice = newData.StartingPrice
	dbData.StartDate = newData.StartDate.Time
	dbData.EndDate = newData.EndDate.Time
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

// func (ed *EventData) GetEvent() ([]events.Event, error) {
// 	var events []events.Event
// 	if err := ed.db.Find(&events).Error; err != nil {
// 		logrus.Error("DATA : Error Get All Events : ", err.Error())
// 		return nil, err
// 	}
// 	return events, nil
// }
