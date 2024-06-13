package handler

import (
	"capstone-mikti/features/bookings"
	"capstone-mikti/helper"
	"capstone-mikti/helper/jwt"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type BookingHandler struct {
	service bookings.BookingServiceInterface
	jwt     jwt.JWTInterface
}

func NewHandler(b bookings.BookingServiceInterface, j jwt.JWTInterface) *BookingHandler {
	return &BookingHandler{
		service: b,
		jwt:     j,
	}
}

func (bh *BookingHandler) CreateBooking() echo.HandlerFunc {
	return func(c echo.Context) error {

		var input = new(BookingRequest)
		if err := c.Bind(input); err != nil {
			c.Logger().Info("Handler : Bind Input Error : ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Invalid Event Input", nil))
		}

		isValid, errors := helper.ValidateJSON(input)
		if !isValid {
			return c.JSON(http.StatusBadRequest, helper.FormatResponseValidation("Invalid Format Request", errors))
		}

		token, err := bh.jwt.ExtractToken(c)
		if err != nil {
			c.Logger().Error("Handler : Extract Token Error : ", err)
			return c.JSON(http.StatusUnauthorized, helper.FormatResponse("Unauthorized", nil))
		}
		user_id := int(token.ID)

		var serviceInput = new(bookings.Booking)
		serviceInput.TicketID = input.TicketID
		serviceInput.UserID = user_id
		serviceInput.FullName = input.FullName
		serviceInput.Email = input.Email
		serviceInput.PhoneNumber = input.PhoneNumber
		serviceInput.Quantity = input.Quantity

		result, err := bh.service.CreateBooking(*serviceInput)
		if err != nil {
			if strings.Contains(err.Error(), "Booking Failed") {
				return c.JSON(http.StatusBadRequest, helper.FormatResponse("Booking Failed", nil))
			}
			c.Logger().Info("Handler : Create Failed : ", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Booking Failed", nil))
		}

		response := new(AllBookingResponse)
		response.ID = uint(result.ID)
		response.TicketID = result.TicketID
		response.UserID = result.UserID
		response.FullName = result.FullName
		response.Email = result.Email
		response.PhoneNumber = result.PhoneNumber
		response.BookingDate = result.BookingDate
		response.IsBooked = result.IsBooked
		response.IsPaid = result.IsPaid
		response.Quantity = input.Quantity

		return c.JSON(http.StatusCreated, helper.FormatResponse("Succesfuly Booking Ticket", response))
	}
}

func (bh *BookingHandler) GetAll() echo.HandlerFunc {
	return func(c echo.Context) error {
		sort := c.QueryParam("status")

		getAll, err := bh.service.GetAll(sort)

		if err != nil {
			c.Logger().Info("Handler : Get All Error : ", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Get All proses Failed", nil))
		}

		var response []AllBookingResponse
		for _, booking := range getAll {
			bookingResponse := AllBookingResponse{
				ID:          uint(booking.ID),
				TicketID:    booking.TicketID,
				UserID:      booking.UserID,
				FullName:    booking.FullName,
				PhoneNumber: booking.PhoneNumber,
				Email:       booking.Email,
				BookingDate: booking.BookingDate,
				Quantity:    booking.Quantity,
				IsBooked:    booking.IsBooked,
				IsPaid:      booking.IsPaid,
			}
			response = append(response, bookingResponse)
		}

		return c.JSON(http.StatusOK, helper.FormatResponse("Data Booking", response))
	}
}

func (bh *BookingHandler) DeleteBooking() echo.HandlerFunc {
	return func(c echo.Context) error {

		id := c.Param("id")
		bookingID, err := strconv.Atoi(id)

		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Invalid Booking ID", nil))
		}

		result, err := bh.service.DeleteBooking(bookingID)

		if err != nil {
			c.Logger().Info("Handler : Delete Booking Error : ", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Delete Booking Error", nil))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse("Success Delete Booking", result))
	}
}

func (bh *BookingHandler) GetDetail() echo.HandlerFunc {
	return func(c echo.Context) error {

		id := c.Param("id")
		bookingID, err := strconv.Atoi(id)

		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Invalid Booking ID", nil))
		}

		getDetail, err := bh.service.GetDetail(bookingID)

		if err != nil {
			c.Logger().Info("Handler : Get Detail Booking Error : ", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Get Detail Booking Error", nil))
		}

		var response []BookingResponse

		bookingResponse := BookingResponse{
			ID:          uint(getDetail.ID),
			FullName:    getDetail.FullName,
			PhoneNumber: getDetail.PhoneNumber,
			Email:       getDetail.Email,
			BookingDate: getDetail.BookingDate,
			Quantity:    getDetail.Quantity,
			Total:       getDetail.Total,
			IsBooked:    getDetail.IsBooked,
			IsPaid:      getDetail.IsPaid,
			Event: EventResponse{
				EventID:    getDetail.EventID,
				EventTitle: getDetail.EventTitle,
			},
			Ticket: TicketResponse{
				TicketID: getDetail.TicketID,
				Name:     getDetail.Name,
				Price:    getDetail.Price,
			},
			User: UserResponse{
				UserID:   getDetail.UserID,
				Username: getDetail.Username,
			},
		}
		response = append(response, bookingResponse)

		return c.JSON(http.StatusOK, helper.FormatResponse("Success Get Booking", response))
	}
}
