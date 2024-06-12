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
	// Call Data CheckEvent
	checkEvent, err := ws.data.CheckEvent(new_data.EventID)

	if err != nil {
		logrus.Error("Service : CheckEvent Error : ", err.Error())
		return nil, errors.New("ERROR CheckEvent")
	}

	if len(checkEvent) == 0 {
		logrus.Warn("Service : CheckEvent Warning")
		return nil, errors.New("WARNING Event Does Not Exists")
	}

	// Call Data CheckUnique
	checkUnique, err := ws.data.CheckUnique(user_id, new_data.EventID)

	if err != nil {
		logrus.Error("Service : CheckUnique Error : ", err.Error())
		return nil, errors.New("ERROR CheckUnique")
	}

	if len(checkUnique) > 0 {
		logrus.Warn("Service : Warning CheckUnique")
		return nil, errors.New("WARNING Event Already Added In The User's Wishlist")
	}

	// Call Data Create
	create, err := ws.data.Create(user_id, new_data)

	if err != nil {
		logrus.Error("Service : Create Error : ", err.Error())
		return nil, errors.New("ERROR Create")
	}

	return create, nil
}

// GetAll
func (ws *WishlistService) GetAll(user_id int) ([]wishlists.WishlistInfo, error) {
	// Call Data GetAll
	getAll, err := ws.data.GetAll(user_id)

	// Error Handling
	if err != nil {
		logrus.Error("Service : GetAll Error : ", err.Error())
		return nil, errors.New("ERROR GetAll")
	}

	return getAll, nil
}

// GetByEventID
func (ws *WishlistService) GetByEventID(user_id int, event_id int) ([]wishlists.WishlistInfo, error) {
	// Call Data GetByEventID
	getByEventId, err := ws.data.GetByEventID(user_id, event_id)

	// Error Handling
	if err != nil {
		logrus.Error("Service : GetByEventID Error : ", err.Error())
		return nil, errors.New("ERROR GetByEventID")
	}

	return getByEventId, nil
}

// Delete
func (ws *WishlistService) Delete(user_id int, event_id int) (bool, error) {
	// Call Data Delete
	delete, err := ws.data.Delete(user_id, event_id)

	// Error Handling
	if err != nil {
		logrus.Error("Service : Delete Error : ", err.Error())
		return false, errors.New("ERROR Delete")
	}

	return delete, nil
}
