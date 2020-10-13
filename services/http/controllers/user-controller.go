package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	user "github.com/nugrohosam/gosampleapi/services/http/requests/user"
	usecases "github.com/nugrohosam/gosampleapi/usecases"
)

// UserHandlerDetail is use
func UserHandlerDetail() gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO
	}
}

// UserHandlerStore is use
func UserHandlerStore() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user user.UserStoreDto
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		usecases.CreateUser(user.Name)
	}
}
