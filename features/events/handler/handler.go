package handler

import (
	"capstone-mikti/features/events"
	"capstone-mikti/helper"
	"capstone-mikti/helper/jwt"
	"net/http"
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

// func (eh *EventHandler) GetEvent() echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		result, err := eh.service.GetEvent()
// 		if err != nil {
// 			c.Logger().Error("Handler : Get Event Error : ", err.Error())
// 			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Get Event Error", nil))
// 		}

// 		var responses []EventResponse
// 		for _, event := range result {
// 			response := EventResponse{
// 				CategoryFK:    event.CategoryFK,
// 				Title:         event.Title,
// 				StartDate:     event.StartDate,
// 				EndDate:       event.EndDate,
// 				Description:   event.Description,
// 				StartingPrice: event.StartingPrice,
// 			}
// 			responses = append(responses, response)
// 		}

// 		return c.JSON(http.StatusOK, helper.FormatResponse("Success Get Events", responses))
// 	}
// }
