package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HelloWorldHandler is use
func HelloWorldHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "Hello World"})
	}
}
