package handler

import (
	"capstone-mikti/features/wishlists"
	"capstone-mikti/helper"
	"strconv"

	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
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

// GetByID
func (wh *WishlistHandler) GetByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.Logger().Error("Handler : GetByID Invalid id : ", err)
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Invalid id", nil))
		}

		result, err := wh.service.GetByID(id)

		if err != nil {
			c.Logger().Error("Handler : GetByID Error : ", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("GetByID process failed", nil))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse("GetByID process success", result))
	}
}

// Delete
func (wh *WishlistHandler) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		event_id, err := strconv.Atoi(c.Param("event_id"))

		if err != nil {
			c.Logger().Error("Handler : Delete Invalid id : ", err)
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Invalid event_id", nil))
		}

		err = wh.service.Delete(event_id)

		if err != nil {
			c.Logger().Error("Handler : Delete Error : ", err)
			if err == gorm.ErrRecordNotFound {
				return c.JSON(http.StatusNotFound, helper.FormatResponse("Wishlist not found", nil))
			}
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Delete process failed", nil))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse("Delete process success", nil))
	}
}
