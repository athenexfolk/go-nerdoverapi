package category

import (
	"nerdoverapi/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllCategoriesHandler(c *gin.Context) {
	categoryList, err := GetAllCategories(c.Request.Context())

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, categoryList)
}

func GetCategoryByIDHandler(c *gin.Context) {
	id := c.Param("id")
	category, err := GetCategoryByID(c.Request.Context(), id)
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

	if err := util.ValidateSlug(newCategory.Slug); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdCategory, err := CreateCategory(c.Request.Context(), newCategory)
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

	updatedCategory, err := UpdateCategory(c.Request.Context(), id, category)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedCategory)
}

func DeleteCategoryHandler(c *gin.Context) {
	id := c.Param("id")

	deletedCategory, err := DeleteCategory(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, deletedCategory)
}
