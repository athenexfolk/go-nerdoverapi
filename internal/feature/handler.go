package feature

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func UploadImageHandler(c *gin.Context) {
	file, header, err := c.Request.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read image: " + err.Error()})
		return
	}
	defer file.Close()

	if isImage := strings.HasPrefix(header.Header.Get("Content-Type"), "image/"); !isImage {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File is not an image"})
		return
	}

	publicURL, err := UploadImage(c.Request.Context(), file, header.Filename, header.Header.Get("Content-Type"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Upload failed: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"url": publicURL})
}

func GetAllImagesHandler(c *gin.Context) {
	imageUrls, err := GetAllImages(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Get images failed: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, imageUrls)
}

func ExportLessonHandler(c *gin.Context) {
	zipData, err := ExportLesson(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Export file failed: " + err.Error()})
		return
	}

	filename := fmt.Sprintf("export_%d.zip", time.Now().UTC().Unix())

	c.Header("Content-Type", "application/zip")
	c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename))
	c.Header("Content-Length", fmt.Sprintf("%d", len(zipData)))

	c.Writer.Write(zipData)
}
