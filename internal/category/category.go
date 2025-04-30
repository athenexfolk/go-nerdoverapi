package category

type Category struct {
	Name string `json:"name" binding:"required"`
	Slug string `json:"slug" binding:"required"`
}
