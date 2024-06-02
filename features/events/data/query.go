package data

import (
	events "capstone-mikti/features/events"
	"errors"

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

	dbData.CategoryID = newData.CategoryID
	dbData.EventTitle = newData.EventTitle
	dbData.City = newData.City
	dbData.Address = newData.Address
	dbData.StartingPrice = newData.StartingPrice
	dbData.StartDate = newData.ParseStartDate
	dbData.EndDate = newData.ParseEndDate
	dbData.Description = newData.Description
	dbData.Highlight = newData.Highlight
	dbData.ImportantInformation = newData.ImportantInformation
	dbData.ImageUrl = newData.ImageUrl
	dbData.PublicID = newData.PublicID

	if err := ed.db.Create(dbData).Error; err != nil {
		logrus.Error("DATA : Create Error : ", err.Error())
		return nil, err
	}

	newData.StartDate = newData.ParseStartDate.Format("2006-01-02")
	newData.EndDate = newData.ParseEndDate.Format("2006-01-02")

	return &newData, nil
}

func (e *EventData) GetAll() ([]events.AllEvent, error) {
	var listEvent = []events.AllEvent{}

	err := e.db.Table("events").
		Select("events.*, categories.category_name").
		Joins("JOIN categories ON categories.id = events.category_id").
		Where("events.deleted_at is null").
		Scan(&listEvent).Error

	if err != nil {
		logrus.Error("Data : Get All Error : ", err.Error())
		return listEvent, err
	}

	return listEvent, nil
}

func (e *EventData) GetDetail(id int) ([]events.Event, error) {
	var event = []events.Event{}

	err := e.db.Table("events").
		Select("events.*, categories.category_name").
		Joins("JOIN categories ON categories.id = events.category_id").
		Where("events.id = ?", id).
		Where("events.deleted_at is null").
		First(&event).Error

	if err != nil {
		logrus.Error("DATA : Error Get By ID : ", err.Error())
		return nil, err
	}

	return event, nil

}

func (e *EventData) UpdateEvent(id int, newData events.Event) (*events.Event, error) {
	var qry = e.db.Table("events").
		Where("id = ?", id).
		Where("deleted_at is null").
		Updates(Event{
			CategoryID:           newData.CategoryID,
			EventTitle:           newData.EventTitle,
			City:                 newData.City,
			Address:              newData.Address,
			StartingPrice:        newData.StartingPrice,
			StartDate:            newData.ParseStartDate,
			EndDate:              newData.ParseEndDate,
			Description:          newData.Description,
			Highlight:            newData.Highlight,
			ImportantInformation: newData.ImportantInformation,
			ImageUrl:             newData.ImageUrl,
			PublicID:             newData.PublicID,
		})

	if err := qry.Error; err != nil {
		logrus.Error("DATA : Error Update Event : ", err.Error())
		return nil, err
	}

	if dataCount := qry.RowsAffected; dataCount < 1 {
		logrus.Error("DATA : No Row Affected")
		return nil, nil
	}

	newData.StartDate = newData.ParseStartDate.Format("2006-01-02")
	newData.EndDate = newData.ParseEndDate.Format("2006-01-02")

	return &newData, nil
}

func (e *EventData) DeleteEvent(id int) (bool, error) {

	var event Event
	var query = e.db.Where("id = ? ", id).Delete(&event)

	if err := query.Error; err != nil {
		logrus.Error("Data : Error Delete Event : ", err.Error())
		return false, nil
	}

	if query.RowsAffected == 0 {
		logrus.Warn("Data : No Event Deleted")
		return false, errors.New("no event deleted")
	}

	return true, nil

}
