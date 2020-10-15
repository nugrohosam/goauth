package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	validator "github.com/go-playground/validator/v10"
	user "github.com/nugrohosam/gosampleapi/services/http/requests/v1"
	usecases "github.com/nugrohosam/gosampleapi/usecases"
)

// UserHandlerDetail is use
func UserHandlerDetail() gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO
	}
}

// UserHandlerIndex is use
func UserHandlerIndex() gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO
	}
}

// UserHandlerStore is use
func UserHandlerStore() gin.HandlerFunc {
	return func(c *gin.Context) {
		validate := validator.New()
		user := user.UserStoreDto{}
		if err := validate.Struct(&user); err != nil {
			validationErrors := err.(validator.ValidationErrors)
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErrors})
			return
		}

		usecases.CreateUser(user.Name)
	}
}
