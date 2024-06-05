package categories

import "github.com/labstack/echo/v4"

type Category struct {
	ID           uint   `json:"id"`
	CategoryName string `json:"category_name"`
}
type UpdateCategory struct {
	NewCategoryName string `json:"new_category_name"`
}

type CategoryDataInterface interface {
	CreateCategory(newData Category) (*Category, error)
	GetByID(id int) (*Category, error)
	GetCategories() ([]Category, error)
	GetByCategoryName(Name string) (*Category, error)
	UpdateName(id int, newData UpdateCategory) (bool, error)
}

type CategoryServiceInterface interface {
	CreateCategory(newData Category) (*Category, error)
	UpdateCategoryName(id int, newData UpdateCategory) (bool, error)
	GetCategoryDetail(id int) (*Category, error)
	GetCategories() ([]Category, error)
}
type CategoryHandlerInterface interface {
	CreateCategory() echo.HandlerFunc
	UpdateCategory() echo.HandlerFunc
	GetCategories() echo.HandlerFunc
	GetCategory() echo.HandlerFunc
}
