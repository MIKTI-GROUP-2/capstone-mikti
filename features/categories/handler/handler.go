package handler

import (
	"capstone-mikti/features/categories"
	"capstone-mikti/helper"
	"capstone-mikti/helper/jwt"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type CategoryHandler struct {
	service categories.CategoryServiceInterface
	jwt     jwt.JWTInterface
}

func NewHandler(s categories.CategoryServiceInterface, j jwt.JWTInterface) *CategoryHandler {
	return &CategoryHandler{
		service: s,
		jwt:     j,
	}
}
func (ch *CategoryHandler) CreateCategory() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input = new(CreateCategoryRequest)

		if err := ch.jwt.ValidateRole(c); !err {
			c.Logger().Info("Handler : Unauthorized Access : ", errors.New("you have no permission to access this feature"))
			return c.JSON(http.StatusUnauthorized, helper.FormatResponse("Restricted Access", nil))
		}

		if err_bind := c.Bind(&input); err_bind != nil {
			c.Logger().Info("Handler : Bind Input Error : ", err_bind.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Invalid User Input", nil))
		}

		isValid, err_validate := helper.ValidateJSON(input)
		if !isValid {
			return c.JSON(http.StatusBadRequest, helper.FormatResponseValidation("Invalid Format Request", err_validate))
		}

		var serviceInput = new(categories.Category)
		serviceInput.CategoryName = input.Name

		result, error_service := ch.service.CreateCategory(*serviceInput)

		if error_service != nil {
			if strings.Contains(error_service.Error(), "ERROR category already created") {
				return c.JSON(http.StatusBadRequest, helper.FormatResponse("Category Name already Existed", nil))
			}
			c.Logger().Info("Handler : Create Category Failed : ", error_service.Error())
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Create Category Process Failed", nil))
		}
		var response = new(CategoryResponse)
		response.CategoryName = result.CategoryName

		return c.JSON(http.StatusCreated, helper.FormatResponse("Success Creating New Category", response))
	}
}

func (ch *CategoryHandler) UpdateCategory() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))

		var input = new(UpdateCategoryRequest)
		if err := ch.jwt.ValidateRole(c); !err {
			c.Logger().Info("Handler : Unauthorized Access : ", errors.New("you have no permission to access this feature"))
			return c.JSON(http.StatusUnauthorized, helper.FormatResponse("Restricted Access", nil))
		}

		if err_bind := c.Bind(&input); err_bind != nil {
			c.Logger().Info("Handler : Bind Input Error : ", err_bind.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Invalid User Input", nil))
		}

		isValid, err_validate := helper.ValidateJSON(input)
		if !isValid {
			return c.JSON(http.StatusBadRequest, helper.FormatResponseValidation("Invalid Format Request", err_validate))
		}
		var serviceUpdate = new(categories.UpdateCategory)
		serviceUpdate.NewCategoryName = input.Name
		res_update, err_update := ch.service.UpdateCategoryName(id, *serviceUpdate)

		if err_update != nil {
			c.Logger().Info("Handler : Update Category Error : ", err_update.Error())
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Update Category Error", nil))
		}
		return c.JSON(http.StatusOK, helper.FormatResponse("Success Update Category Name", res_update))
	}
}
func (ch *CategoryHandler) GetCategories() echo.HandlerFunc {
	return func(c echo.Context) error {
		res, err_get := ch.service.GetCategories()
		if err_get != nil {
			c.Logger().Info("Handler : Get Categories Error : ", err_get.Error())
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Get Categories Error", nil))
		}
		return c.JSON(http.StatusOK, helper.FormatResponse("All Categories Alvailable", res))
	}
}
func (ch *CategoryHandler) GetCategory() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		res, err_get := ch.service.GetCategoryDetail(id)
		if err_get != nil {
			c.Logger().Info("Handler : Get Categories Error : ", err_get.Error())
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Get Categories Error", nil))
		}
		response := new(CategoryResponse)
		response.CategoryName = res.CategoryName
		return c.JSON(http.StatusOK, helper.FormatResponse("This is the Detail of category", response))
	}
}
