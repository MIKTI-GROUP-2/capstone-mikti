package service

import (
	events "capstone-mikti/features/events"
	"capstone-mikti/helper/jwt"
	"capstone-mikti/utils/cloudinary"
	"errors"
	"time"

	"github.com/sirupsen/logrus"
)

type EventService struct {
	data       events.EventDataInterface
	jwt        jwt.JWTInterface
	cloudinary cloudinary.CloudinaryInterface
}

func New(d events.EventDataInterface, j jwt.JWTInterface, c cloudinary.CloudinaryInterface) *EventService {
	return &EventService{
		data:       d,
		jwt:        j,
		cloudinary: c,
	}
}

func (e *EventService) CreateEvent(newData events.Event) (*events.Event, error) {
	_, err := e.data.GetByTitle(newData.EventTitle)

	if err != nil {
		logrus.Error("Service : Titke already registered")
		return nil, errors.New("ERROR Title already registered by another user")
	}

	secureURL, publicId, err := e.cloudinary.UploadImageHelper(newData.ImageFile)
	if err != nil {
		logrus.Error("Error uploading image: ", err)
		return nil, errors.New("Error Upload Image")
	}

	//waktu
	layout := "2006-01-02"

	parseStartDate, _ := time.Parse(layout, newData.StartDate)
	parseEndDate, _ := time.Parse(layout, newData.EndDate)

	newData.ParseStartDate = parseStartDate
	newData.ParseEndDate = parseEndDate

	newData.ImageUrl = secureURL
	newData.PublicID = publicId

	result, err := e.data.CreateEvent(newData)
	if err != nil {
		logrus.Error("Service : Error Create : ", err.Error())
		return nil, errors.New("ERROR Create")
	}

	return result, nil
}

func (e *EventService) GetAll(title string, category string, times string, city string, price int, sort string) ([]events.AllEvent, error) {

	result, err := e.data.GetAll(title, category, times, city, price, sort)

	if err != nil {
		logrus.Error("Service : Get All Eerror : ", err.Error())
		return nil, errors.New("Error Get All")
	}

	return result, nil
}

func (e *EventService) GetDetail(id int) ([]events.Event, error) {
	result, err := e.data.GetDetail(id)

	if err != nil {
		logrus.Error("Service : Get All Eerror : ", err.Error())
		return nil, errors.New("Error Get All")
	}

	return result, nil
}

func (e *EventService) UpdateEvent(id int, newData events.Event) (*events.Event, error) {
	layout := "2006-01-02"

	parseStartDate, _ := time.Parse(layout, newData.StartDate)
	parseEndDate, _ := time.Parse(layout, newData.EndDate)

	newData.ParseStartDate = parseStartDate
	newData.ParseEndDate = parseEndDate

	result, err := e.data.UpdateEvent(id, newData)
	if err != nil {
		logrus.Error("Service : Error Create : ", err.Error())
		return nil, errors.New("ERROR Create")
	}

	return result, nil
}

func (e *EventService) DeleteEvent(id int) (bool, error) {
	result, err := e.data.DeleteEvent(id)

	if err != nil {
		logrus.Error("Service : Error Create : ", err.Error())
		return false, errors.New("ERROR Create")
	}

	return result, nil

}
