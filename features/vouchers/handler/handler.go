package handler

import (
	"capstone-mikti/features/vouchers"
	"capstone-mikti/helper"
	"capstone-mikti/helper/jwt"
	"errors"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type VoucherHandler struct {
	service vouchers.VoucherServiceInterface
	jwt     jwt.JWTInterface
}

func NewHandler(s vouchers.VoucherServiceInterface, j jwt.JWTInterface) *VoucherHandler {
	return &VoucherHandler{
		service: s,
		jwt:     j,
	}
}

// CreateVoucher handles the creation of a new voucher.
func (v *VoucherHandler) CreateVoucher() echo.HandlerFunc {
	return func(c echo.Context) error {
		if err := v.jwt.ValidateRole(c); !err {
			c.Logger().Info("Handler : Unauthorized Access : ", errors.New("you have no permission to access this feature"))
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Restricted Access", nil))
		}

		var input = new(InputRequest)
		if err := c.Bind(input); err != nil {
			logrus.Error(err.Error())
			c.Logger().Info("Handler : Bind Input Error : ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Invalid Voucher Input", nil))
		}

		isValid, errors := helper.ValidateJSON(input)
		if !isValid {
			return c.JSON(http.StatusBadRequest, helper.FormatResponseValidation("Invalid Format Request", errors))
		}

		var serviceInput = new(vouchers.Voucher)
		serviceInput.Code = input.Code
		serviceInput.Name = input.Name
		serviceInput.Quantity = input.Quantity
		serviceInput.Price = input.Price
		serviceInput.ExpiryDate = input.ExpiredDate
		serviceInput.EventID = uint(input.EventID)

		result, err := v.service.CreateVoucher(*serviceInput)
		if err != nil {
			c.Logger().Info("Handler : CreateVoucher Failed : ", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Create Voucher Process Failed", nil))
		}

		var response = new(InputResponse)
		response.Code = result.Code
		response.Name = result.Name
		response.Quantity = result.Quantity
		response.Price = result.Price
		response.ExpiryDate = result.ExpiryDate
		response.EventID = uint(result.EventID)
		response.Status = result.Status

		return c.JSON(http.StatusCreated, helper.FormatResponse("Success Create Voucher", response))
	}
}

// GetVoucherByID handles retrieving a voucher by its ID.
func (v *VoucherHandler) GetVoucher() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		voucherID, err := strconv.Atoi(id)
		if err != nil {
			c.Logger().Info("Handler : Invalid ID : ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Invalid Voucher ID", nil))
		}

		result, err := v.service.GetVoucher(voucherID)
		if err != nil {
			c.Logger().Info("Handler : GetVoucherByID Failed : ", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Get Voucher Process Failed", nil))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse("Success Get Voucher", result))
	}
}

// GetVouchers handles retrieving a voucher by its ID.
func (v *VoucherHandler) GetVouchers() echo.HandlerFunc {
	return func(c echo.Context) error {
		result, err := v.service.GetVouchers()
		if err != nil {
			c.Logger().Info("Handler : Get All Voucher Failed : ", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Get All Voucher Process Failed", nil))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse("Success Get Voucher", result))
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

		return c.JSON(http.StatusOK, helper.FormatResponse("Success Get Voucher", result))
	}
}

// UpdateVoucher handles updating an existing voucher.
func (v *VoucherHandler) UpdateVoucher() echo.HandlerFunc {
	return func(c echo.Context) error {
		if err := v.jwt.ValidateRole(c); !err {
			c.Logger().Info("Handler : Unauthorized Access : ", errors.New("you have no permission to access this feature"))
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Restricted Access", nil))
		}

		id := c.Param("id")
		voucherID, err := strconv.Atoi(id)
		if err != nil {
			c.Logger().Info("Handler : Invalid ID : ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Invalid Voucher ID", nil))
		}

		var input = new(UpdateRequest)
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
		serviceUpdate.Name = input.Name
		serviceUpdate.Quantity = input.Quantity
		serviceUpdate.Price = input.Price
		serviceUpdate.ExpiryDate = input.ExpiredDate
		serviceUpdate.EventID = uint(input.EventID)

		res, err := v.service.UpdateVoucher(voucherID, *serviceUpdate)
		if err != nil {
			c.Logger().Info("Handler : UpdateVoucher Failed : ", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Update Voucher Process Failed", nil))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse("Success Update Voucher", res))
	}
}

// ActivateVoucher handles deleting an existing voucher.
func (v *VoucherHandler) ActivateVoucher() echo.HandlerFunc {
	return func(c echo.Context) error {
		if err := v.jwt.ValidateRole(c); !err {
			c.Logger().Info("Handler : Unauthorized Access : ", errors.New("you have no permission to access this feature"))
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Restricted Access", nil))
		}

		id := c.Param("id")
		voucherID, err := strconv.Atoi(id)
		if err != nil {
			c.Logger().Info("Handler : Invalid ID : ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Invalid Voucher ID", nil))
		}

		res, err := v.service.ActivateVoucher(voucherID)
		if err != nil {
			c.Logger().Info("Handler : ActivateVoucher Failed : ", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Activate Voucher Process Failed", nil))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse("Success Activate Voucher", res))
	}
}

// DeactivateVoucher handles deleting an existing voucher.
func (v *VoucherHandler) DeactivateVoucher() echo.HandlerFunc {
	return func(c echo.Context) error {
		if err := v.jwt.ValidateRole(c); !err {
			c.Logger().Info("Handler : Unauthorized Access : ", errors.New("you have no permission to access this feature"))
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Restricted Access", nil))
		}

		id := c.Param("id")
		voucherID, err := strconv.Atoi(id)
		if err != nil {
			c.Logger().Info("Handler : Invalid ID : ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Invalid Voucher ID", nil))
		}

		res, err := v.service.DeactivateVoucher(voucherID)
		if err != nil {
			c.Logger().Info("Handler : DeactivateVoucher Failed : ", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Deactivate Voucher Process Failed", nil))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse("Success Deactivate Voucher", res))
	}
}
