package category

type Category struct {
	ID   string `gorm:"primaryKey" json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
}
