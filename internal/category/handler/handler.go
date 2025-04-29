package handler

import (
	"github.com/gin-gonic/gin"
	"nerdoverapi/internal/category/model"
	"nerdoverapi/internal/category/service"
	"net/http"
)

func GetAllCategoriesHandler(c *gin.Context) {
	categoryList, err := service.GetAllCategories(c.Request.Context())

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, categoryList)
}

func GetCategoryByIDHandler(c *gin.Context) {
	id := c.Param("id")
	category, err := service.GetCategoryByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, category)
}

func CreateCategoryHandler(c *gin.Context) {
	var newCategory model.Category

	if err := c.ShouldBindJSON(&newCategory); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdCategory, err := service.CreateCategory(c.Request.Context(), newCategory)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdCategory)
}

func UpdateCategoryHandler(c *gin.Context) {
	id := c.Param("id")

	var category model.Category

	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedCategory, err := service.UpdateCategory(c.Request.Context(), id, category)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedCategory)
}

func DeleteCategoryHandler(c *gin.Context) {
	id := c.Param("id")

	deletedCategory, err := service.DeleteCategory(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, deletedCategory)
}
