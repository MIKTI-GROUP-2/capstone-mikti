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
func (ws *WishlistService) Create(new_data wishlists.Wishlist) (*wishlists.Wishlist, error) {
	create, err := ws.data.Create(new_data)

	if err != nil {
		logrus.Error("Service : Create Error : ", err.Error())
		return nil, errors.New("ERROR Create")
	}

	return create, nil
}

// GetAll
func (ws *WishlistService) GetAll() ([]wishlists.WishlistInfo, error) {
	result, err := ws.data.GetAll()

	if err != nil {
		logrus.Error("Service : GetAll Error : ", err.Error())
		return nil, errors.New("ERROR GetAll")
	}

	return result, nil
}

// GetByID
func (ws *WishlistService) GetByID(id int) ([]wishlists.WishlistInfo, error) {
	result, err := ws.data.GetByID(id)

	if err != nil {
		logrus.Error("Service : GetByID Error : ", err.Error())
		return nil, errors.New("ERROR GetByID")
	}

	return result, nil
}

// Delete
func (ws *WishlistService) Delete(event_id int) error {
	err := ws.data.Delete(event_id)

	if err != nil {
		logrus.Error("Service : Delete Error : ", err.Error())
		return errors.New("ERROR Delete")
	}

	return nil
}
