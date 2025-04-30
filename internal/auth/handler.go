package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoginWithGoogleHandler(c *gin.Context) {
	var body struct {
		IDToken string `json:"idToken" binding:"required"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing idToken"})
		return
	}

	jwt, err := VerifyGoogleToken(c.Request.Context(), body.IDToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"accessToken": jwt})
}
