package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	validator "github.com/go-playground/validator/v10"
	copier "github.com/jinzhu/copier"
	helpers "github.com/nugrohosam/gosampleapi/helpers"
	permission "github.com/nugrohosam/gosampleapi/services/http/requests/v1"
	role "github.com/nugrohosam/gosampleapi/services/http/requests/v1"
	resource "github.com/nugrohosam/gosampleapi/services/http/resources/v1"
	usecases "github.com/nugrohosam/gosampleapi/usecases"
)

// UserRoleHandlerIndex ..
func UserRoleHandlerIndex() gin.HandlerFunc {
	return func(c *gin.Context) {
		var queryParams permission.ListQuery
		c.BindQuery(&queryParams)

		userRoles, total, err := usecases.GetUserRole(queryParams.Search, queryParams.PerPage, queryParams.Page, queryParams.OrderBy)

		if err != nil {
			c.JSON(http.StatusInternalServerError, helpers.ResponseErr(err.Error()))
			return
		}
		if cap(userRoles) > 0 {
			var listUserRolesResource = resource.UserRoleListItems{}
			copier.Copy(&listUserRolesResource, &userRoles)
			if queryParams.Paginate {
				resourceData := helpers.BuildPaginate(queryParams.PerPage, queryParams.Page, total, &userRoles, &listUserRolesResource)
				c.JSON(http.StatusOK, helpers.ResponseModelStruct(resourceData))
			} else {
				c.JSON(http.StatusOK, helpers.ResponseMany(listUserRolesResource))
			}
		} else {
			c.JSON(http.StatusOK, helpers.ResponseMany(nil))
		}
	}
}

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
