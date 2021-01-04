package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	validator "github.com/go-playground/validator/v10"
	helpers "github.com/nugrohosam/gosampleapi/helpers"
	role "github.com/nugrohosam/gosampleapi/services/http/requests/v1"
	usecases "github.com/nugrohosam/gosampleapi/usecases"
)

// UserRoleHandlerCreate ..
func UserRoleHandlerCreate() gin.HandlerFunc {
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

// UserRoleHandlerUpdate ..
func UserRoleHandlerUpdate() gin.HandlerFunc {
	return func(c *gin.Context) {
		var role role.UpdateRole
		c.BindJSON(&role)

		validate := validator.New()
		if err := validate.Struct(role); err != nil {
			validationErrors := err.(validator.ValidationErrors)
			fieldsErrors := helpers.TransformValidations(validationErrors)
			c.JSON(http.StatusBadRequest, helpers.ResponseErrValidation(fieldsErrors))
			return
		}

		roleID := c.Param("id")
		if err := usecases.UpdateRole(roleID, role.Name); err != nil {
			c.JSON(http.StatusBadRequest, helpers.ResponseErr(err.Error()))
			return
		}
	}
}

// UserRoleHandlerDelete ..
func UserRoleHandlerDelete() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, helpers.ResponseModelStruct(nil))
	}
}
