package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	validator "github.com/go-playground/validator/v10"
	helpers "github.com/nugrohosam/gosampleapi/helpers"
	auth "github.com/nugrohosam/gosampleapi/services/http/requests/v1"
	usecases "github.com/nugrohosam/gosampleapi/usecases"
)

// AuthHandlerLogin is use
func AuthHandlerLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		var authLogin auth.AuthLogin
		c.BindJSON(&authLogin)

		validate := validator.New()
		if err := validate.Struct(authLogin); err != nil {
			validationErrors := err.(validator.ValidationErrors)
			fieldsErrors := helpers.TransformValidations(validationErrors)
			c.JSON(http.StatusBadRequest, helpers.ResponseErrValidation(fieldsErrors))
			return
		}

		usecases.AuthBasic(authLogin.EmailOrUsername, authLogin.Password)
		c.JSON(http.StatusOK, nil)
	}
}
