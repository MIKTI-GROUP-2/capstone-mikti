package data

import (
	"capstone-mikti/features/categories"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type CategoryData struct {
	db *gorm.DB
}

func New(db *gorm.DB) *CategoryData {
	return &CategoryData{
		db: db,
	}
}

func (cd *CategoryData) CreateCategory(newData categories.Category) (*categories.Category, error) {
	var dbCategory = new(Category)
	dbCategory.CategoryName = newData.CategoryName

	if err := cd.db.Create(dbCategory).Error; err != nil {
		logrus.Error("DATA : Create Category Error : ", err.Error())
		return nil, err
	}
	return &newData, nil
}
func (cd *CategoryData) GetByID(id int) (*categories.Category, error) {
	var listCategory categories.Category
	var qry = cd.db.Table("categories").Select("categories.*").
		Where("categories.id = ?", id).
		Scan(&listCategory)

	if err := qry.Error; err != nil {
		logrus.Error("DATA : Error Get By ID : ", err.Error())
		return &listCategory, err
	}

	return &listCategory, nil
}
func (cd *CategoryData) GetByCategoryName(Name string) (*categories.Category, error) {
	var dbData = new(Category)
	dbData.CategoryName = Name
	var qry = cd.db.Where("category_name = ?", dbData.CategoryName).First(dbData)

	if err := qry.Error; err != nil {
		logrus.Error("DATA : Error Get By Name : ", err.Error())
		return nil, err
	}
	var result = new(categories.Category)
	result.ID = dbData.ID
	result.CategoryName = dbData.CategoryName

	return result, nil
}
func (cd *CategoryData) UpdateName(id int, newData categories.UpdateCategory) (bool, error) {
	var qry = cd.db.Table("categories").Where("id = ?", id).Updates(Category{
		CategoryName: newData.NewCategoryName,
	})

	if err := qry.Error; err != nil {
		logrus.Error("DATA : Error Update Category : ", err.Error())
		return false, err
	}

	if dataCount := qry.RowsAffected; dataCount < 1 {
		logrus.Error("DATA : No Row Affected")
		return false, nil
	}

	return true, nil
}
func (cd *CategoryData) GetCategories() ([]categories.Category, error) {
	var categories []categories.Category

	err := cd.db.Find(&categories).Error

	if err != nil {
		logrus.Error("DATA : Error Get All Data : ", err.Error())
		return nil, err
	}
	return categories, nil
}
