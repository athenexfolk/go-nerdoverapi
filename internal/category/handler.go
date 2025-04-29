package category

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAllCategoriesHandler(c *gin.Context) {
	categoryList := GetAllCategories()
	c.JSON(http.StatusOK, categoryList)
}

func GetCategoryByIDHandler(c *gin.Context) {
	id := c.Param("id")
	category, err := GetCategoryByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, category)
}

func CreateCategoryHandler(c *gin.Context) {
	var newCategory Category

	if err := c.ShouldBindJSON(&newCategory); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdCategory, err := CreateCategory(newCategory)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdCategory)
}

func UpdateCategoryHandler(c *gin.Context) {
	id := c.Param("id")

	var category Category

	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedCategory, err := UpdateCategory(id, category)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedCategory)
}

func DeleteCategoryHandler(c *gin.Context) {
	id := c.Param("id")

	deletedCategory, err := DeleteCategory(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, deletedCategory)
}
