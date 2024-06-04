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

		// Get Entity
		newWishlist := wishlists.Wishlist{
			UserID:  user_id,
			EventID: request.EventID,
		}

		// Call Service
		create, err := wh.service.Create(user_id, newWishlist)

		if err != nil {
			c.Logger().Error("Handler : Create Error : ", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Create process failed", nil))
		}

		// Get Response
		response := WishlistResponse{
			ID:      create.ID,
			UserID:  create.UserID,
			EventID: create.EventID,
		}

		return c.JSON(http.StatusCreated, helper.FormatResponse("Create process success", response))
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

// GetByID
func (wh *WishlistHandler) GetByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		// Extract user.id from JWT
		token, err := wh.jwt.ExtractToken(c)

		if err != nil {
			c.Logger().Error("Handler : Extract Token Error : ", err)
			return c.JSON(http.StatusUnauthorized, helper.FormatResponse("Unauthorized", nil))
		}

		user_id := int(token.ID)

		// Extract wishlist.id from path parameter
		wishlist_id, _ := strconv.Atoi(c.Param("id"))

		// Call Service
		getById, err := wh.service.GetByID(user_id, wishlist_id)

		if err != nil {
			c.Logger().Error("Handler : GetByID Error : ", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("GetByID process failed", nil))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse("GetByID process success", getById))
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
		err = wh.service.Delete(user_id, event_id)

		if err != nil {
			c.Logger().Error("Handler : Delete Error : ", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Delete process failed", nil))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse("Delete process success", nil))
	}
}
