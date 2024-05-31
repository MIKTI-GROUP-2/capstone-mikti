package handler

import (
	"capstone-mikti/features/vouchers"
	"capstone-mikti/helper"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type VoucherHandler struct {
	service vouchers.VoucherServiceInterface
}

func NewHandler(s vouchers.VoucherServiceInterface) *VoucherHandler {
	return &VoucherHandler{
		service: s,
	}
}

// CreateVoucher handles the creation of a new voucher.
func (v *VoucherHandler) CreateVoucher() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input = new(CreateVoucherInput)
		if err := c.Bind(&input); err != nil {
			c.Logger().Info("Handler : Bind Input Error : ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Invalid Voucher Input", nil))
		}

		isValid, errors := helper.ValidateJSON(input)
		if !isValid {
			return c.JSON(http.StatusBadRequest, helper.FormatResponseValidation("Invalid Format Request", errors))
		}

		var serviceInput = new(vouchers.Voucher)
		serviceInput.Code = input.Code
		serviceInput.Discount = input.Discount
		serviceInput.ExpiryDate = input.ExpiryDate
		serviceInput.EventID = input.EventID
		serviceInput.Status = input.Status

		result, err := v.service.CreateVoucher(*serviceInput)
		if err != nil {
			c.Logger().Info("Handler : CreateVoucher Failed : ", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Create Voucher Process Failed", nil))
		}

		var response = new(VoucherResponse)
		response.ID = result.ID
		response.Code = result.Code
		response.Discount = result.Discount
		response.ExpiryDate = result.ExpiryDate
		response.EventID = result.EventID
		response.Status = result.Status

		return c.JSON(http.StatusCreated, helper.FormatResponse("Success Create Voucher", response))
	}
}

// GetVoucherByID handles retrieving a voucher by its ID.
func (v *VoucherHandler) GetVoucherByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		voucherID, err := strconv.Atoi(id)
		if err != nil {
			c.Logger().Info("Handler : Invalid ID : ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Invalid Voucher ID", nil))
		}

		result, err := v.service.GetVoucherByID(voucherID)
		if err != nil {
			c.Logger().Info("Handler : GetVoucherByID Failed : ", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Get Voucher Process Failed", nil))
		}

		var response = new(VoucherResponse)
		response.ID = result.ID
		response.Code = result.Code
		response.Discount = result.Discount
		response.ExpiryDate = result.ExpiryDate
		response.EventID = result.EventID
		response.Status = result.Status

		return c.JSON(http.StatusOK, helper.FormatResponse("Success Get Voucher", response))
	}
}

// GetVoucherByCode handles retrieving a voucher by its code.
func (v *VoucherHandler) GetVoucherByCode() echo.HandlerFunc {
	return func(c echo.Context) error {
		code := c.Param("code")

		result, err := v.service.GetVoucherByCode(code)
		if err != nil {
			c.Logger().Info("Handler : GetVoucherByCode Failed : ", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Get Voucher Process Failed", nil))
		}

		var response = new(VoucherResponse)
		response.ID = result.ID
		response.Code = result.Code
		response.Discount = result.Discount
		response.ExpiryDate = result.ExpiryDate
		response.EventID = result.EventID
		response.Status = result.Status

		return c.JSON(http.StatusOK, helper.FormatResponse("Success Get Voucher", response))
	}
}

// UpdateVoucher handles updating an existing voucher.
func (v *VoucherHandler) UpdateVoucher() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		voucherID, err := strconv.Atoi(id)
		if err != nil {
			c.Logger().Info("Handler : Invalid ID : ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Invalid Voucher ID", nil))
		}

		var input = new(UpdateVoucherInput)
		if err := c.Bind(&input); err != nil {
			c.Logger().Info("Handler : Bind Input Error : ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Invalid Voucher Input", nil))
		}

		isValid, errors := helper.ValidateJSON(input)
		if !isValid {
			return c.JSON(http.StatusBadRequest, helper.FormatResponseValidation("Invalid Format Request", errors))
		}

		var serviceUpdate = new(vouchers.UpdateVoucher)
		serviceUpdate.Code = input.Code
		serviceUpdate.Discount = input.Discount
		serviceUpdate.ExpiryDate = input.ExpiryDate
		serviceUpdate.EventID = input.EventID
		serviceUpdate.Status = input.Status

		success, err := v.service.UpdateVoucher(voucherID, *serviceUpdate)
		if err != nil {
			c.Logger().Info("Handler : UpdateVoucher Failed : ", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Update Voucher Process Failed", nil))
		}

		if !success {
			c.Logger().Info("Handler : UpdateVoucher No Rows Affected")
			return c.JSON(http.StatusNotFound, helper.FormatResponse("Voucher Not Found", nil))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse("Success Update Voucher", nil))
	}
}

// DeleteVoucher handles deleting an existing voucher.
func (v *VoucherHandler) DeleteVoucher() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		voucherID, err := strconv.Atoi(id)
		if err != nil {
			c.Logger().Info("Handler : Invalid ID : ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Invalid Voucher ID", nil))
		}

		err = v.service.DeleteVoucher(voucherID)
		if err != nil {
			c.Logger().Info("Handler : DeleteVoucher Failed : ", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Delete Voucher Process Failed", nil))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse("Success Delete Voucher", nil))
	}
}

// VoucherResponse represents the response structure for voucher operations.
type VoucherResponse struct {
	ID         int       json:"id"
	Code       string    json:"code"
	Discount   float64   json:"discount"
	ExpiryDate time.Time json:"expiry_date"
	EventID    int       json:"event_id"
	Status     bool      json:"status"
}

// CreateVoucherInput represents the input structure for creating a voucher.
type CreateVoucherInput struct {
	Code       string    json:"code" validate:"required"
	Discount   float64   json:"discount" validate:"required"
	ExpiryDate time.Time json:"expiry_date" validate:"required"
	EventID    int       json:"event_id" validate:"required"
	Status     bool      json:"status" validate:"required"
}

// UpdateVoucherInput represents the input structure for updating a voucher.
type UpdateVoucherInput struct {
	Code       string    json:"code" validate:"required"
	Discount   float64   json:"discount" validate:"required"
	ExpiryDate time.Time json:"expiry_date" validate:"required"
	EventID    int       json:"event_id" validate:"required"
	Status     bool      json:"status" validate:"required"
}