package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	validator "github.com/go-playground/validator/v10"
	helpers "github.com/nugrohosam/gosampleapi/helpers"
	role "github.com/nugrohosam/gosampleapi/services/http/requests/v1"
	usecases "github.com/nugrohosam/gosampleapi/usecases"
)

// RoleHandlerCreate ..
func RoleHandlerCreate() gin.HandlerFunc {
	return func(c *gin.Context) {
		var role role.CreateRole
		c.BindJSON(&role)

		validate := validator.New()
		if err := validate.Struct(role); err != nil {
			validationErrors := err.(validator.ValidationErrors)
			fieldsErrors := helpers.TransformValidations(validationErrors)
			c.JSON(http.StatusBadRequest, helpers.ResponseErrValidation(fieldsErrors))
			return
		}

		if err := usecases.CreateRole(role.Name); err != nil {
			c.JSON(http.StatusBadRequest, helpers.ResponseErr(err.Error()))
			return
		}
	}
}

// RoleHandlerUpdate ..
func RoleHandlerUpdate() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, helpers.ResponseModelStruct(nil))
	}
}

// RoleHandlerDelete ..
func RoleHandlerDelete() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, helpers.ResponseModelStruct(nil))
	}
}
