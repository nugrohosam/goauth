package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	validator "github.com/go-playground/validator/v10"
	helpers "github.com/nugrohosam/gosampleapi/helpers"
	auth "github.com/nugrohosam/gosampleapi/services/http/requests/v1"
	resources "github.com/nugrohosam/gosampleapi/services/http/resources/v1"
	usecases "github.com/nugrohosam/gosampleapi/usecases"
)

// KeyMiddlewareAuth ..
const KeyMiddlewareAuth = "auth-middleware"

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

		token, err := usecases.AuthBasic(authLogin.EmailOrUsername, authLogin.Password)
		if err != nil {
			c.JSON(http.StatusBadRequest, helpers.ResponseErr(err.Error()))
			return
		}

		dataResponse := resources.AuthSuccess{
			Token: token,
		}

		c.JSON(http.StatusOK, helpers.ResponseModelStruct(dataResponse))
	}
}

// AuthHandlerRegister is use
func AuthHandlerRegister() gin.HandlerFunc {
	return func(c *gin.Context) {
		var authRegister auth.AuthRegister
		c.BindJSON(&authRegister)

		validate := validator.New()
		if err := validate.Struct(authRegister); err != nil {
			validationErrors := err.(validator.ValidationErrors)
			fieldsErrors := helpers.TransformValidations(validationErrors)
			c.JSON(http.StatusBadRequest, helpers.ResponseErrValidation(fieldsErrors))
			return
		}

		if err := usecases.RegisterBasic(authRegister.Name, authRegister.Username, authRegister.Email, authRegister.Password); err != nil {
			c.JSON(http.StatusBadRequest, helpers.ResponseErr(err.Error()))
			return
		}

		c.JSON(http.StatusOK, nil)
	}
}
