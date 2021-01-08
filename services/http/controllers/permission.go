package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	validator "github.com/go-playground/validator/v10"
	copier "github.com/jinzhu/copier"
	helpers "github.com/nugrohosam/gosampleapi/helpers"
	permission "github.com/nugrohosam/gosampleapi/services/http/requests/v1"
	resource "github.com/nugrohosam/gosampleapi/services/http/resources/v1"
	usecases "github.com/nugrohosam/gosampleapi/usecases"
)

// PermissionHandlerIndex ..
func PermissionHandlerIndex() gin.HandlerFunc {
	return func(c *gin.Context) {
		var queryParams permission.ListQuery
		c.BindQuery(&queryParams)

		permissions, err := usecases.GetPermission(queryParams.Search, queryParams.PerPage, queryParams.Page, queryParams.OrderBy)

		if err != nil {
			c.JSON(http.StatusInternalServerError, helpers.ResponseErr(err.Error()))
			return
		}
		fmt.Println(permissions)

		if cap(permissions) > 0 {
			if queryParams.Paginate {
				var permissionItem = resource.PermissionListItems{}
				copier.Copy(&permissionItem, &permissions)
				perPage, _ := strconv.Atoi(queryParams.PerPage)
				resource := resource.PermissionPaginate{
					Items:   permissionItem,
					PerPage: perPage,
					Total:   cap(permissions),
				}

				c.JSON(http.StatusOK, helpers.ResponseModelStruct(resource))
			} else {
				c.JSON(http.StatusOK, helpers.ResponseModelStruct(permissions))
			}
		}
	}
}

// PermissionHandlerCreate ..
func PermissionHandlerCreate() gin.HandlerFunc {
	return func(c *gin.Context) {
		var permission permission.CreateRole
		c.BindJSON(&permission)

		validate := validator.New()
		if err := validate.Struct(permission); err != nil {
			validationErrors := err.(validator.ValidationErrors)
			fieldsErrors := helpers.TransformValidations(validationErrors)
			c.JSON(http.StatusBadRequest, helpers.ResponseErrValidation(fieldsErrors))
			return
		}

		if err := usecases.CreateRole(permission.Name); err != nil {
			c.JSON(http.StatusBadRequest, helpers.ResponseErr(err.Error()))
			return
		}
	}
}

// PermissionHandlerUpdate ..
func PermissionHandlerUpdate() gin.HandlerFunc {
	return func(c *gin.Context) {
		var permission permission.UpdateRole
		c.BindJSON(&permission)

		validate := validator.New()
		if err := validate.Struct(permission); err != nil {
			validationErrors := err.(validator.ValidationErrors)
			fieldsErrors := helpers.TransformValidations(validationErrors)
			c.JSON(http.StatusBadRequest, helpers.ResponseErrValidation(fieldsErrors))
			return
		}

		roleID := c.Param("id")
		if err := usecases.UpdateRole(roleID, permission.Name); err != nil {
			c.JSON(http.StatusBadRequest, helpers.ResponseErr(err.Error()))
			return
		}
	}
}

// PermissionHandlerDelete ..
func PermissionHandlerDelete() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, helpers.ResponseModelStruct(nil))
	}
}
