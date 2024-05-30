package handler

import (
	"capstone-mikti/features/wishlists"
	"capstone-mikti/helper"

	"net/http"

	"github.com/labstack/echo/v4"
)

type WishlistHandler struct {
	service wishlists.WishlistServiceInterface
}

func NewHandler(s wishlists.WishlistServiceInterface) *WishlistHandler {
	return &WishlistHandler{
		service: s,
	}
}

// GetAll implements wishlists.WishlistHandlerInterface.
func (wh *WishlistHandler) GetAll() echo.HandlerFunc {
	return func(c echo.Context) error {
		getAll, err := wh.service.GetAll()

		if err != nil {
			c.Logger().Info("Handler : GetAll Error : ", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Get All process failed", nil))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse("Get All process success", getAll))
	}
}
