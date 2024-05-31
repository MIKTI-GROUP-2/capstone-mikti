package handler

import (
	events "capstone-mikti/features/events"
	"capstone-mikti/helper"
	"capstone-mikti/helper/jwt"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type EventHandler struct {
	service events.EventServiceInterface
	jwt     jwt.JWTInterface
}

func NewHandler(s events.EventServiceInterface, j jwt.JWTInterface) *EventHandler {
	return &EventHandler{
		service: s,
		jwt:     j,
	}
}

func (e *EventHandler) CreateEvent() echo.HandlerFunc {
	return func(c echo.Context) error {

		if isValid := e.jwt.ValidateRole(c); !isValid {
			return c.JSON(http.StatusUnauthorized, helper.FormatResponse("Only Admin can access this endpoint", nil))
		}

		var input = new(EventInput)
		if err := c.Bind(&input); err != nil {
			c.Logger().Info("Handler : Bind Input Error : ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Invalid Event Input", nil))
		}

		isValid, errors := helper.ValidateJSON(input)
		if !isValid {
			return c.JSON(http.StatusBadRequest, helper.FormatResponseValidation("Invalid Format Request", errors))
		}

		var serviceInput = new(events.Event)

		serviceInput.CategoryFK = input.CategoryFK
		serviceInput.Title = input.Title
		serviceInput.City = input.City
		serviceInput.Address = input.Address
		serviceInput.StartingPrice = input.StartingPrice
		serviceInput.StartDate = input.StartDate
		serviceInput.EndDate = input.EndDate
		serviceInput.Description = input.Description
		serviceInput.Highlight = input.Highlight
		serviceInput.ImportantInformation = input.ImportantInformation
		serviceInput.Image = input.Image

		result, err := e.service.CreateEvent(*serviceInput)

		if err != nil {
			if strings.Contains(err.Error(), "Title already registered by another event") {
				return c.JSON(http.StatusBadRequest, helper.FormatResponse("Title Already Registered", nil))
			}
			c.Logger().Info("Handler : Create Failed : ", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Title Already Registered", nil))
		}

		var response = new(EventResponse)
		response.CategoryFK = result.CategoryFK
		response.Title = result.Title
		response.StartDate = result.StartDate
		response.EndDate = result.EndDate
		response.Description = result.Description
		response.StartingPrice = result.StartingPrice

		return c.JSON(http.StatusCreated, helper.FormatResponse("Succesfuly created event", response))
	}
}

func (e *EventHandler) GetAll() echo.HandlerFunc {
	return func(c echo.Context) error {
		getAll, err := e.service.GetAll()

		if err != nil {
			c.Logger().Info("Handler : Get All Error : ", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Get All proses Failed", nil))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse("Get All Proses Succes", getAll))
	}
}

func (e *EventHandler) GetDetail() echo.HandlerFunc {
	return func(c echo.Context) error {

		id := c.Param("id")

		eventID, err := strconv.Atoi(id)

		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Invalid Event ID", nil))
		}

		get, err := e.service.GetDetail(eventID)

		if err != nil {
			c.Logger().Info("Handler : Get Event Error : ", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Get Event proses Failed", nil))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse("Get Event Proses Succes", get))

	}
}

func (e *EventHandler) UpdateEvent() echo.HandlerFunc {
	return func(c echo.Context) error {

		if isValid := e.jwt.ValidateRole(c); !isValid {
			return c.JSON(http.StatusUnauthorized, helper.FormatResponse("Only Admin can access this endpoint", nil))
		}

		id := c.Param("id")
		eventID, err := strconv.Atoi(id)

		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Invalid Event ID", nil))
		}

		var input = new(UpdateEvent)
		if err := c.Bind(&input); err != nil {
			c.Logger().Info("Handler : Bind Input Error : ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Invalid Event Input", nil))
		}

		var serviceUpdate = new(events.UpdateEvent)
		serviceUpdate.CategoryFK = input.CategoryFK
		serviceUpdate.Title = input.Title
		serviceUpdate.City = input.City
		serviceUpdate.Address = input.Address
		serviceUpdate.StartingPrice = input.StartingPrice
		serviceUpdate.StartDate = input.StartDate
		serviceUpdate.EndDate = input.EndDate
		serviceUpdate.Description = input.Description
		serviceUpdate.Highlight = input.Highlight
		serviceUpdate.ImportantInformation = input.ImportantInformation
		serviceUpdate.Image = input.Image
		serviceUpdate.PublicID = input.PublicID

		result, err := e.service.UpdateEvent(int(eventID), *serviceUpdate)
		if err != nil {
			c.Logger().Info("Handler : Update Event Error : ", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Update Event Error", nil))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse("Success Update Event", result))

	}
}

func (e *EventHandler) DeleteEvent() echo.HandlerFunc {
	return func(c echo.Context) error {
		if isValid := e.jwt.ValidateRole(c); !isValid {
			return c.JSON(http.StatusUnauthorized, helper.FormatResponse("Only Admin can access this endpoint", nil))
		}

		id := c.Param("id")
		eventID, err := strconv.Atoi(id)

		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Invalid Event ID", nil))
		}

		result, err := e.service.DeleteEvent(eventID)

		if err != nil {
			c.Logger().Info("Handler : Delete Event Error : ", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Delete Event Error", nil))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse("Success Delete Event", result))
	}
}
