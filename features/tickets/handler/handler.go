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
		var request CreateTicketRequest

		if err := c.Bind(&request); err != nil {
			c.Logger().Error("Handler : Create Bind Error : ", err)
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Invalid request body", nil))
		}

		// New Data
		new_data := tickets.Ticket{
			EventID:    request.EventID,
			Name:       request.Name,
			TicketDate: request.TicketDate,
			Quantity:   request.Quantity,
			Price:      request.Price,
		}

		// Call Service
		create, err := th.service.Create(new_data)

		if err != nil {
			c.Logger().Error("Handler : Create Error : ", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Create process failed", nil))
		}

		// Get Response
		response := TicketResponse{
			EventID:    create.EventID,
			Name:       create.Name,
			TicketDate: create.TicketDate,
			Quantity:   create.Quantity,
			Price:      create.Price,
		}

		return c.JSON(http.StatusOK, helper.FormatResponse("Create process success", response))
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
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Invalid ID", nil))
		}

		// Call Service
		getById, err := th.service.GetByID(id)

		if err != nil {
			c.Logger().Error("Handler : GetByID Error : ", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("GetByID process failed", nil))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse("GetByID process success", getById))
	}
}

// Update
func (th *TicketHandler) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		// Validate Admin
		is_admin := th.jwt.ValidateRole(c)

		if !is_admin {
			return c.JSON(http.StatusUnauthorized, helper.FormatResponse("Only admin can access this endpoint", nil))
		}

		// Extract ticket.id from path parameter
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Invalid ID", nil))
		}

		// Get Request
		var request UpdateTicketRequest

		if err := c.Bind(&request); err != nil {
			c.Logger().Error("Handler : Create Bind Error : ", err)
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Invalid request body", nil))
		}

		// New Data
		new_data := tickets.Ticket{
			EventID:    request.EventID,
			Name:       request.Name,
			TicketDate: request.TicketDate,
			Quantity:   request.Quantity,
			Price:      request.Price,
		}

		// Call Service
		update, err := th.service.Update(id, new_data)

		if err != nil {
			c.Logger().Error("Handler : Update Error : ", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Update process failed", nil))
		}

		if !update {
			return c.JSON(http.StatusNotFound, helper.FormatResponse("No ticket found with the given ID", nil))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse("Update process success", update))
	}
}

// Delete
func (th *TicketHandler) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		// Validate Admin
		is_admin := th.jwt.ValidateRole(c)

		if !is_admin {
			return c.JSON(http.StatusUnauthorized, helper.FormatResponse("Only admin can access this endpoint", nil))
		}

		// Extract ticket.id from path parameter
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Invalid ID", nil))
		}

		// Call Service
		update, err := th.service.Delete(id)

		if err != nil {
			c.Logger().Error("Handler : Update Error : ", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Delete process failed", nil))
		}

		if !update {
			return c.JSON(http.StatusNotFound, helper.FormatResponse("No ticket found with the given ID", nil))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse("Delete process success", update))
	}
}
