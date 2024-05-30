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

// GetAll implements wishlists.WishlistServiceInterface.
func (ws *WishlistService) GetAll() ([]wishlists.WishlistInfo, error) {
	result, err := ws.data.GetAll()

	if err != nil {
		logrus.Error("Service : GetAll Error : ", err.Error())
		return nil, errors.New("ERROR GetAll")
	}

	return result, nil
}
