package lesson

type Lesson struct {
	Title        string  `json:"title" binding:"required"`
	Slug         string  `json:"slug" binding:"required"`
	Cover        *string `json:"cover,omitempty"`
	Content      *string `json:"content"`
	ContentPath  string  `json:"contentPath"`
	CategorySlug string  `json:"categorySlug" binding:"required"`
	CategoryName string  `json:"categoryName" binding:"required"`
}

type CreateLessonDto struct {
	Title        string  `json:"title" binding:"required"`
	Slug         string  `json:"slug" binding:"required"`
	CategorySlug string  `json:"categorySlug" binding:"required"`
	CategoryName string  `json:"categoryName" binding:"required"`
	Cover        *string `json:"cover,omitempty"`
}

type UpdateLessonDto struct {
	Title *string `json:"title,omitempty"`
	Cover *string `json:"cover,omitempty"`
}

type UpdateContentDto struct {
	Content string `json:"content" binding:"required"`
}
