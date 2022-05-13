package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func MessageGet() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, map[string]string{
			"message": "Found me",
		})
	}
}
