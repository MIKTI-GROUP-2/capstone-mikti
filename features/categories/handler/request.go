package handler

type CreateCategoryRequest struct {
	Name string `json:"category_name" form:"category_name" validate:"required"`
}

type UpdateCategoryRequest struct {
	Name string `json:"update_category_name" form:"update_category_name" validate:"required"`
}
