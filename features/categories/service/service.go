package service

import (
	"capstone-mikti/features/categories"
	"errors"

	"github.com/sirupsen/logrus"
)

type CategoryService struct {
	data categories.CategoryDataInterface
}

func New(d categories.CategoryDataInterface) *CategoryService {
	return &CategoryService{
		data: d,
	}
}

func (c *CategoryService) CreateCategory(newData categories.Category) (*categories.Category, error) {
	_, err := c.data.GetByCategoryName(newData.CategoryName)
	if err == nil {
		logrus.Error("Service : Category already existed")
		return nil, errors.New("ERROR category already created")
	}
	result, err := c.data.CreateCategory(newData)
	if err != nil {
		logrus.Error("Service : Error Register : ", err.Error())
		return nil, errors.New("ERROR Register")
	}

	return result, nil

}
func (c *CategoryService) UpdateCategoryName(id int, newData categories.UpdateCategory) (bool, error) {
	res, err := c.data.UpdateName(id, newData)

	if err != nil {
		logrus.Error("Service : Update Category Name error : ", err.Error())
		return false, errors.New("ERROR Category Update")
	}
	return res, nil
}
func (c *CategoryService) GetCategoryDetail(id int) (*categories.Category, error) {
	res, err := c.data.GetByID(id)

	if err != nil {
		logrus.Error("Service : Get Profile Error : ", err.Error())
		return nil, errors.New("ERROR Get Profile")
	}

	return res, nil
}
func (c *CategoryService) GetCategories() ([]categories.Category, error) {
	res, err := c.data.GetCategories()

	if err != nil {
		logrus.Error("Service : Get All Category Error ", err.Error())
		return nil, errors.New("ERROR Get All Category")
	}
	return res, nil
}
