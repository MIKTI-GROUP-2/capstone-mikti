package handler

import (
	"capstone-mikti/features/wishlists"
	"capstone-mikti/helper"
	"capstone-mikti/helper/jwt"
	"strconv"

	"net/http"

	"github.com/labstack/echo/v4"
)

// Controller

type WishlistHandler struct {
	service wishlists.WishlistServiceInterface
	jwt     jwt.JWTInterface
}

func NewHandler(s wishlists.WishlistServiceInterface, j jwt.JWTInterface) *WishlistHandler {
	return &WishlistHandler{
		service: s,
		jwt:     j,
	}
}

// Create
func (wh *WishlistHandler) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		// Extract user.id from JWT
		token, err := wh.jwt.ExtractToken(c)

		if err != nil {
			c.Logger().Error("Handler : Extract Token Error : ", err)
			return c.JSON(http.StatusUnauthorized, helper.FormatResponse("Unauthorized", nil))
		}

		user_id := int(token.ID)

		// Get Request
		var request WishlistRequest

		if err := c.Bind(&request); err != nil {
			c.Logger().Error("Handler : Create Bind Error : ", err)
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Invalid request body", nil))
		}

		// New Data
		new_data := wishlists.Wishlist{
			UserID:  user_id,
			EventID: request.EventID,
		}

		// Call Service
		create, err := wh.service.Create(user_id, new_data)

		if err != nil {
			c.Logger().Error("Handler : Create Error : ", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Create process failed", nil))
		}

		// Get Response
		response := WishlistResponse{
			UserID:  create.UserID,
			EventID: create.EventID,
		}

		return c.JSON(http.StatusOK, helper.FormatResponse("Create process success", response))
	}
}

// GetAll
func (wh *WishlistHandler) GetAll() echo.HandlerFunc {
	return func(c echo.Context) error {
		// Extract user.id from JWT
		token, err := wh.jwt.ExtractToken(c)

		if err != nil {
			c.Logger().Error("Handler : Extract Token Error : ", err)
			return c.JSON(http.StatusUnauthorized, helper.FormatResponse("Unauthorized", nil))
		}

		user_id := int(token.ID)

		// Call Service
		getAll, err := wh.service.GetAll(user_id)

		if err != nil {
			c.Logger().Error("Handler : GetAll Error : ", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("GetAll process failed", nil))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse("GetAll process success", getAll))
	}
}

// GetByEventID
func (wh *WishlistHandler) GetByEventID() echo.HandlerFunc {
	return func(c echo.Context) error {
		// Extract user.id from JWT
		token, err := wh.jwt.ExtractToken(c)

		if err != nil {
			c.Logger().Error("Handler : Extract Token Error : ", err)
			return c.JSON(http.StatusUnauthorized, helper.FormatResponse("Unauthorized", nil))
		}

		user_id := int(token.ID)

		// Extract wishlist.id from path parameter
		event_id, _ := strconv.Atoi(c.Param("event_id"))

		// Call Service
		getByEventId, err := wh.service.GetByEventID(user_id, event_id)

		if err != nil {
			c.Logger().Error("Handler : GetByEventID Error : ", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("GetByEventID process failed", nil))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse("GetByEventID process success", getByEventId))
	}
}

// Delete
func (wh *WishlistHandler) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		// Extract user.id from JWT
		token, err := wh.jwt.ExtractToken(c)

		if err != nil {
			c.Logger().Error("Handler : Extract Token Error : ", err)
			return c.JSON(http.StatusUnauthorized, helper.FormatResponse("Unauthorized", nil))
		}

		user_id := int(token.ID)

		// Extract event.id from path parameter
		event_id, _ := strconv.Atoi(c.Param("event_id"))

		// Call Service
		delete, err := wh.service.Delete(user_id, event_id)

		if err != nil {
			c.Logger().Error("Handler : Delete Error : ", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Delete process failed", nil))
		}

		if !delete {
			return c.JSON(http.StatusNotFound, helper.FormatResponse("No wishlist found with the given ID", nil))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse("Delete process success", nil))
	}
}
