package dtos

type CreateCategoryDto struct {
	Name string `json:"name" binding:"required"`
	Slug string `json:"slug" binding:"required"`
}
