package handler

import (
	"capstone-mikti/features/tickets"
	"capstone-mikti/helper"
	"capstone-mikti/helper/jwt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// Controller

type TicketHandler struct {
	service tickets.TicketServiceInterface
	jwt     jwt.JWTInterface
}

func NewHandler(s tickets.TicketServiceInterface, j jwt.JWTInterface) *TicketHandler {
	return &TicketHandler{
		service: s,
		jwt:     j,
	}
}

// Create
func (th *TicketHandler) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		// Validate Admin
		is_admin := th.jwt.ValidateRole(c)

		if !is_admin {
			return c.JSON(http.StatusUnauthorized, helper.FormatResponse("Only admin can access this endpoint", nil))
		}

		// Get Request
		var request = new(CreateTicketRequest)

		if err := c.Bind(request); err != nil {
			c.Logger().Error("Handler : Create Bind Error : ", err)
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Invalid request body", nil))
		}

		// Get Entity
		newTicket := tickets.Ticket{
			EventID:    request.EventID,
			Name:       request.Name,
			TicketDate: request.TicketDate,
			Quantity:   request.Quantity,
			Price:      request.Price,
		}

		// Call Service
		create, err := th.service.Create(newTicket)

		if err != nil {
			c.Logger().Error("Handler : Create Error : ", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Create process failed", nil))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse("Create process success", create))
	}
}

// GetAll
func (th *TicketHandler) GetAll() echo.HandlerFunc {
	return func(c echo.Context) error {
		// Validate Admin
		is_admin := th.jwt.ValidateRole(c)

		if !is_admin {
			return c.JSON(http.StatusUnauthorized, helper.FormatResponse("Only admin can access this endpoint", nil))
		}

		// Call Service
		getAll, err := th.service.GetAll()

		if err != nil {
			c.Logger().Error("Handler : GetAll Error : ", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("GetAll process failed", nil))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse("GetAll process success", getAll))
	}
}

// GetByID
func (th *TicketHandler) GetByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		// Validate Admin
		is_admin := th.jwt.ValidateRole(c)

		if !is_admin {
			return c.JSON(http.StatusUnauthorized, helper.FormatResponse("Only admin can access this endpoint", nil))
		}

		// Extract ticket.id from path parameter
		ticket_id, _ := strconv.Atoi(c.Param("id"))

		// Call Service
		getById, err := th.service.GetByID(ticket_id)

		if err != nil {
			c.Logger().Error("Handler : GetByID Error : ", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("GetByID process failed", nil))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse("GetByID process success", getById))
	}
}