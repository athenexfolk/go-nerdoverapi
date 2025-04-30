package lesson

import (
	"nerdoverapi/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllCategoriesHandler(c *gin.Context) {
	lessonList, err := GetAllLessons(c.Request.Context())

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, lessonList)
}

func GetLessonByIDHandler(c *gin.Context) {
	id := c.Param("id")
	category, err := GetLessonByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, category)
}

func CreateLessonHandler(c *gin.Context) {
	var dto CreateLessonDto

	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := util.ValidateSlug(dto.Slug); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := util.ValidateSlug(dto.CategorySlug); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdLesson, err := CreateLesson(c.Request.Context(), dto)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdLesson)
}

func UpdateLessonHandler(c *gin.Context) {
	id := c.Param("id")

	var dto UpdateLessonDto

	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedLesson, err := UpdateLesson(c.Request.Context(), id, dto)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedLesson)
}

func UpdateContentHandler(c *gin.Context) {
	id := c.Param("id")

	var dto UpdateContentDto

	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedLesson, err := UpdateContent(c.Request.Context(), id, dto)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedLesson)
}

func DeleteCategoryHandler(c *gin.Context) {
	id := c.Param("id")

	deletedCategory, err := DeleteLesson(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, deletedCategory)
}
