package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	validator "github.com/go-playground/validator/v10"
	helpers "github.com/nugrohosam/gosampleapi/helpers"
	auth "github.com/nugrohosam/gosampleapi/services/http/requests/v1"
	resource "github.com/nugrohosam/gosampleapi/services/http/resources/v1"
	usecases "github.com/nugrohosam/gosampleapi/usecases"
)

func RoleHandlerCreate() gin.HandlerFunc {
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

		dataResponse := resource.AuthSuccess{
			Token: token,
		}

		c.JSON(http.StatusOK, helpers.ResponseModelStruct(dataResponse))
	}
}

func RoleHandlerUpdate() gin.HandlerFunc {
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

		dataResponse := resource.AuthSuccess{
			Token: token,
		}

		c.JSON(http.StatusOK, helpers.ResponseModelStruct(dataResponse))
	}
}

func RoleHandlerDelete() gin.HandlerFunc {
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

		dataResponse := resource.AuthSuccess{
			Token: token,
		}

		c.JSON(http.StatusOK, helpers.ResponseModelStruct(dataResponse))
	}
}
