package service

import (
	"capstone-mikti/features/wishlists"
	"errors"

	"github.com/sirupsen/logrus"
)

// Service

type WishlistService struct {
	data wishlists.WishlistDataInterface
}

func New(d wishlists.WishlistDataInterface) *WishlistService {
	return &WishlistService{
		data: d,
	}
}

// Create
func (ws *WishlistService) Create(user_id int, new_data wishlists.Wishlist) (*wishlists.Wishlist, error) {
	// Get Data CheckEvent
	checkEvent, err := ws.data.CheckEvent(new_data.EventID)

	if err != nil {
		logrus.Error("Service : CheckEvent Error : ", err.Error())
		return nil, errors.New("ERROR CheckEvent")
	}

	if len(checkEvent) == 0 {
		logrus.Error("Service : Event does not exists")
		return nil, errors.New("ERROR Event does not exists")
	}

	// Get Data CheckUnique
	checkUnique, err := ws.data.CheckUnique(user_id, new_data.EventID)

	if err != nil {
		logrus.Error("Service : CheckUnique Error : ", err.Error())
		return nil, errors.New("ERROR CheckUnique")
	}

	if len(checkUnique) > 0 {
		logrus.Error("Service : Event already exists in the user's wishlist")
		return nil, errors.New("EROR Event already exists in the user's wishlist")
	}

	// Get Data Create
	create, err := ws.data.Create(user_id, new_data)

	if err != nil {
		logrus.Error("Service : Create Error : ", err.Error())
		return nil, errors.New("ERROR Create")
	}

	return create, nil
}

// GetAll
func (ws *WishlistService) GetAll(user_id int) ([]wishlists.WishlistInfo, error) {
	// Get Data
	getAll, err := ws.data.GetAll(user_id)

	// Error Handling
	if err != nil {
		logrus.Error("Service : GetAll Error : ", err.Error())
		return nil, errors.New("ERROR GetAll")
	}

	return getAll, nil
}

// GetByID
func (ws *WishlistService) GetByID(user_id, id int) ([]wishlists.WishlistInfo, error) {
	// Get Data
	getById, err := ws.data.GetByID(user_id, id)

	// Error Handling
	if err != nil {
		logrus.Error("Service : GetByID Error : ", err.Error())
		return nil, errors.New("ERROR GetByID")
	}

	return getById, nil
}

// Delete
func (ws *WishlistService) Delete(user_id int, event_id int) error {
	// Get Data
	err := ws.data.Delete(user_id, event_id)

	// Error Handling
	if err != nil {
		logrus.Error("Service : Delete Error : ", err.Error())
		return errors.New("ERROR Delete")
	}

	return nil
}
