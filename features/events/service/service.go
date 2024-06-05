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

	newData.ImageUrl = secureURL
	newData.PublicID = publicId

	//waktu
	layout := "2006-01-02"

	parseStartDate, _ := time.Parse(layout, newData.StartDate)
	parseEndDate, _ := time.Parse(layout, newData.EndDate)

	newData.ParseStartDate = parseStartDate
	newData.ParseEndDate = parseEndDate

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

	//tiome
	layout := "2006-01-02"

	parseStartDate, _ := time.Parse(layout, newData.StartDate)
	parseEndDate, _ := time.Parse(layout, newData.EndDate)

	newData.ParseStartDate = parseStartDate
	newData.ParseEndDate = parseEndDate

	//image
	getPublicID, _ := e.data.GetPublicID(id)

	if newData.ImageFile != nil {
		_, err := e.cloudinary.DeleteImageHelper(getPublicID)
		if err != nil {
			logrus.Error("Error delete image: ", err)
			return nil, errors.New("Error delete Image")
		}

		secureURL, publicId, err := e.cloudinary.UploadImageHelper(newData.ImageFile)
		if err != nil {
			logrus.Error("Error uploading image: ", err)
			return nil, errors.New("Error Upload Image")
		}

		newData.ImageUrl = secureURL
		newData.PublicID = publicId
	}

	result, err := e.data.UpdateEvent(id, newData)
	if err != nil {
		logrus.Error("Service : Error Create : ", err.Error())
		return nil, errors.New("ERROR Create")
	}

	return result, nil
}

func (e *EventService) DeleteEvent(id int) (bool, error) {

	publicID, err := e.data.GetPublicID(id)

	if err != nil {
		logrus.Error("Service : Error Get Public Id : ", err.Error())
		return false, errors.New("ERROR Get Public Id")
	}

	_, err = e.cloudinary.DeleteImageHelper(publicID)

	if err != nil {
		logrus.Error("Service : Error Delete Image : ", err.Error())
		return false, errors.New("ERROR  Delete Image")
	}

	result, err := e.data.DeleteEvent(id)

	if err != nil {
		logrus.Error("Service : Error Delete : ", err.Error())
		return false, errors.New("ERROR Delete")
	}

	return result, nil

}
