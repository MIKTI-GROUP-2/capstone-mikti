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
	result, err := ws.data.Create(new_data)

	if err != nil {
		logrus.Error("Service : Create Error : ", err.Error())
		return nil, errors.New("ERROR Create")
	}

	return result, nil
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

// GetByUserID
func (ws *WishlistService) GetByUserID(user_id int) ([]wishlists.WishlistInfo, error) {
	result, err := ws.data.GetByUserID(user_id)

	if err != nil {
		logrus.Error("Service : GetByUserID Error : ", err.Error())
		return nil, errors.New("ERROR GetByUserID")
	}

	return result, nil
}
