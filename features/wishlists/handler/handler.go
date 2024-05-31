package handler

import (
	"capstone-mikti/features/wishlists"
	"capstone-mikti/helper"
	"strconv"

	"net/http"

	"github.com/labstack/echo/v4"
)

// Controller

type WishlistHandler struct {
	service wishlists.WishlistServiceInterface
}

func NewHandler(s wishlists.WishlistServiceInterface) *WishlistHandler {
	return &WishlistHandler{
		service: s,
	}
}

// Create
func (wh *WishlistHandler) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		var request WishlistRequest
		if err := c.Bind(&request); err != nil {
			c.Logger().Error("Handler : Create Bind Error : ", err)
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Invalid request body", nil))
		}

		newWishlist := wishlists.Wishlist{
			UserID:  request.UserID,
			EventID: request.EventID,
		}

		createdWishlist, err := wh.service.Create(newWishlist)

		if err != nil {
			c.Logger().Error("Handler : Create Error : ", err)
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Create process failed", nil))
		}

		response := WishlistResponse{
			ID:      createdWishlist.ID,
			UserID:  createdWishlist.UserID,
			EventID: createdWishlist.EventID,
		}

		return c.JSON(http.StatusCreated, helper.FormatResponse("Create process success", response))
	}
}

// GetAll
func (wh *WishlistHandler) GetAll() echo.HandlerFunc {
	return func(c echo.Context) error {
		result, err := wh.service.GetAll()

		if err != nil {
			c.Logger().Info("Handler : GetAll Error : ", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("GetAll process failed", nil))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse("GetAll process success", result))
	}
}

// GetByUserID
func (wh *WishlistHandler) GetByUserID() echo.HandlerFunc {
	return func(c echo.Context) error {
		user_id, _ := strconv.Atoi(c.Param("user_id"))

		result, err := wh.service.GetByUserID(user_id)

		if err != nil {
			c.Logger().Info("Handler : GetByUserID Error : ", err.Error())
			return c.JSON(http.StatusNotFound, helper.FormatResponse("GetByUserID process failed", nil))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse("GetByUserID process success", result))
	}
}
