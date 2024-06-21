package handler

import (
	"capstone-mikti/features/payments"
	"capstone-mikti/helper"
	"capstone-mikti/helper/jwt"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type PaymentHandler struct {
	service payments.PaymentServiceInterface
	jwt     jwt.JWTInterface
}

func NewHandler(s payments.PaymentServiceInterface, j jwt.JWTInterface) *PaymentHandler {
	return &PaymentHandler{
		service: s,
		jwt:     j,
	}
}

func (p *PaymentHandler) CreatePayment() echo.HandlerFunc {
	return func(c echo.Context) error {
		token, err := p.jwt.ExtractToken(c)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, helper.FormatResponse("Fail, cant get ID from JWT", nil))
		}

		var input = new(InputRequest)
		if err := c.Bind(input); err != nil {
			logrus.Info("Handler : Bind Input Error : ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Fail to bind request", nil))
		}

		isValid, errors := helper.ValidateJSON(input)
		if !isValid {
			return c.JSON(http.StatusBadRequest, helper.FormatResponseValidation("Invalid Format Request", errors))
		}

		var serviceInput = new(payments.Payment)

		serviceInput.Price = input.Price
		serviceInput.Qty = input.Qty

		serviceInput.UserID = token.ID
		serviceInput.PaymentStatus = 5
		serviceInput.BookID = input.BookID
		serviceInput.VoucherCode = input.VoucherCode
		serviceInput.PaymentType = input.PaymentType
		serviceInput.TransactionDate = input.TransactionDate

		result, response, err := p.service.CreatePayment(*serviceInput)

		var responseData = new(InputResponse)
		responseData.BookID = result.BookID
		responseData.TransactionDate = result.TransactionDate
		responseData.Qty = result.Qty
		responseData.Price = result.Price
		responseData.GrandTotal = result.GrandTotal
		responseData.CallbackURL = response

		if err != nil {
			c.Logger().Info("Handler : Create Payment Failed : ", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Create Payment Process Failed", nil))
		}

		return c.JSON(http.StatusCreated, helper.FormatResponse("Success Create Payment", responseData))

	}
}

func (p *PaymentHandler) NotifPayment() echo.HandlerFunc {
	return func(c echo.Context) error {
		var notificationPayload map[string]interface{}

		err := json.NewDecoder(c.Request().Body).Decode(&notificationPayload)

		fmt.Println("Notification Payload:", notificationPayload)

		if err != nil {
			if strings.Contains(err.Error(), "Order ID Not Found") {
				return c.JSON(http.StatusBadRequest, helper.FormatResponse("Order ID Not Found", nil))
			}

			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Midtrans POST method error", nil))
		}

		var serviceUpdate = new(payments.UpdatePayment)

		res, err := p.service.UpdatePayment(notificationPayload, *serviceUpdate)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Midtrans cannot update the database", nil))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse("Success Update", res))
	}
}

func (p *PaymentHandler) GetAll() echo.HandlerFunc {
	return func(c echo.Context) error {
		token, err := p.jwt.ExtractToken(c)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, helper.FormatResponse("Fail, cant get ID from JWT", nil))
		}

		status := c.QueryParam("status")
		limit := c.QueryParam("limit")
		offset := c.QueryParam("offset")
		paymentDate := c.QueryParam("payment_date")

		limitInt, _ := strconv.Atoi(limit)
		offsetInt, _ := strconv.Atoi(offset)

		var timePublication int
		switch paymentDate {
		case "7Days":
			timePublication = 1
		case "30Days":
			timePublication = 2
		default:
			timePublication = -1
		}

		userID := token.ID

		if token.Role == "Admin" {
			result, err := p.service.GetAll(status, limitInt, offsetInt, timePublication)
			if err != nil {
				c.Logger().Info("Handler : Get All Payment Failed : ", err.Error())
				return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Get All Payment Process Failed", nil))
			}
			return c.JSON(http.StatusOK, helper.FormatResponse("Success Get Payment", result))
		}

		result, err := p.service.GetByUserID(userID, status, limitInt, offsetInt, timePublication)
		if err != nil {
			c.Logger().Info("Handler : Get All Payment Failed : ", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Get All Payment Process Failed", nil))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse("Success Get Payment", result))
	}
}
