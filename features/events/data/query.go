package data

import (
	events "capstone-mikti/features/events"
	"errors"
	"time"

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

func (ed *EventData) GetByTitle(title string) ([]events.TitleEvent, error) {
	var dbData = []events.TitleEvent{}

	var qry = ed.db.Table("events").
		Where("event_title = ?", title).
		Where("deleted_at is null").
		Find(&dbData)

	if err := qry.Error; err != nil {
		logrus.Error("DATA : Error Get By Title : ", err.Error())
		return nil, err
	}

	return dbData, nil
}

func (ed *EventData) CreateEvent(newData events.Event) (*events.Event, error) {
	var dbData = new(Event)

	dbData.CategoryID = uint(newData.CategoryID)
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

func (e *EventData) GetAll(title string, category string, times string, city string, price int, sort string) ([]events.AllEvent, error) {
	var listEvent = []events.AllEvent{}

	var query = e.db.Table("events").
		Select("events.*, categories.id as category_id, categories.category_name").
		Joins("left join categories on categories.id = events.category_id").
		Where("events.deleted_at is null")

	layout := "2006-01-02"

	if title != "" {
		query.Where("events.event_title LIKE ?", "%"+title+"%")
	}

	if category != "" {
		query.Where("categories.category_name LIKE ?", "%"+category+"%")
	}

	if city != "" {
		query.Where("events.city LIKE ?", "%"+city+"%")
	}

	if price != 0 {
		priceRange := 0.2
		minPrice := float64(price) * (1 - priceRange)
		maxPrice := float64(price) * (1 + priceRange)
		query = query.Where("events.starting_price BETWEEN ? AND ?", minPrice, maxPrice)
	}

	if times != "" {
		parseTimes, _ := time.Parse(layout, times)
		query.Where("events.start_date <= ? AND events.end_date >= ?", parseTimes, parseTimes)
	}

	switch sort {
	case "terbaru":
		query.Order("events.created_at DESC")
	case "termahal":
		query.Order("events.starting_price DESC")
	case "termurah":
		query.Order("events.starting_price ASC")
	default:
		query.Order("events.created_at DESC")
	}

	if err := query.Scan(&listEvent).Error; err != nil {
		logrus.Error("Data : Get All Error : ", err.Error())
		return nil, err
	}

	return listEvent, nil
}

func (e *EventData) GetDetail(id int) ([]events.Event, error) {
	var event = []events.Event{}

	err := e.db.Table("events").
		Select("events.*, categories.category_name").
		Joins("left JOIN categories ON categories.id = events.category_id").
		Where("events.id = ?", id).
		Where("events.deleted_at is null").
		First(&event).Error

	if err != nil {
		logrus.Error("DATA : Error Get By ID : ", err.Error())
		return nil, err
	}

	return event, nil

}

func (e *EventData) UpdateEvent(id int, newData events.Event) (bool, error) {
	var qry = e.db.Table("events").
		Where("id = ?", id).
		Where("deleted_at is null").
		Updates(Event{
			CategoryID:           uint(newData.CategoryID),
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
		return false, err
	}

	if dataCount := qry.RowsAffected; dataCount < 1 {
		logrus.Error("DATA : No Row Affected")
		return false, nil
	}

	newData.StartDate = newData.ParseStartDate.Format("2006-01-02")
	newData.EndDate = newData.ParseEndDate.Format("2006-01-02")

	return true, nil
}

func (ed *EventData) GetPublicID(id int) (string, error) {
	var dbData events.PublicID

	var qry = ed.db.Table("events").Select("public_id").Where("id = ?", id).Scan(&dbData)

	if err := qry.Error; err != nil {
		logrus.Error("DATA : Error Get By ID : ", err.Error())
		return "", err
	}

	return dbData.PublicID, nil
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
