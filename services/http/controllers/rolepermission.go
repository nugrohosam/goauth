package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	validator "github.com/go-playground/validator/v10"
	copier "github.com/jinzhu/copier"
	helpers "github.com/nugrohosam/gosampleapi/helpers"
	rolePermission "github.com/nugrohosam/gosampleapi/services/http/requests/v1"
	rolePermissionPermission "github.com/nugrohosam/gosampleapi/services/http/requests/v1"
	resource "github.com/nugrohosam/gosampleapi/services/http/resources/v1"
	usecases "github.com/nugrohosam/gosampleapi/usecases"
)

// RolePermissionHandlerIndex ..
func RolePermissionHandlerIndex() gin.HandlerFunc {
	return func(c *gin.Context) {
		var queryParams rolePermissionPermission.ListQuery
		c.BindQuery(&queryParams)

		rolePermissionPermissions, total, err := usecases.GetRolePermission(queryParams.Search, queryParams.PerPage, queryParams.Page, queryParams.OrderBy)

		if err != nil {
			c.JSON(http.StatusInternalServerError, helpers.ResponseErr(err.Error()))
			return
		}

		if cap(rolePermissionPermissions) > 0 {
			var listRolePermissionsResource = resource.RolePermissionListItems{}
			copier.Copy(&listRolePermissionsResource, &rolePermissionPermissions)
			if queryParams.Paginate {
				resourceData := helpers.BuildPaginate(queryParams.PerPage, queryParams.Page, total, &rolePermissionPermissions, &listRolePermissionsResource)
				c.JSON(http.StatusOK, helpers.ResponseModelStruct(resourceData))
			} else {
				c.JSON(http.StatusOK, helpers.ResponseMany(listRolePermissionsResource))
			}
		} else {
			c.JSON(http.StatusOK, helpers.ResponseMany(nil))
		}
	}
}

// RolePermissionHandlerShow ..
func RolePermissionHandlerShow() gin.HandlerFunc {
	return func(c *gin.Context) {
		ID := c.Param("id")

		if len(ID) < 1 {
			c.JSON(http.StatusBadRequest, helpers.ResponseErr("params id should filled"))
		} else {
			rolePermissionPermission := usecases.ShowRolePermission(ID)
			var rolePermissionPermissionItem = resource.RolePermissionDetail{}
			copier.Copy(&rolePermissionPermissionItem, &rolePermissionPermission)
			if rolePermissionPermission.ID > 0 {
				c.JSON(http.StatusOK, helpers.ResponseOne(rolePermissionPermissionItem))
			} else {
				c.JSON(http.StatusOK, helpers.ResponseOne(nil))
			}
		}

	}
}

// RolePermissionHandlerCreate ..
func RolePermissionHandlerCreate() gin.HandlerFunc {
	return func(c *gin.Context) {
		var rolePermission rolePermission.CreateRolePermission
		c.ShouldBindJSON(&rolePermission)

		validate := helpers.NewValidation()
		if err := validate.Struct(rolePermission); err != nil {
			validationErrors := err.(validator.ValidationErrors)
			fieldsErrors := helpers.TransformValidations(validationErrors)
			c.JSON(http.StatusBadRequest, helpers.ResponseErrValidation(fieldsErrors))
			return
		}

		if err := usecases.CreateRolePermission(rolePermission.RoleID.(int), rolePermission.PermisisonID.(int)); err != nil {
			c.JSON(http.StatusBadRequest, helpers.ResponseErr(err.Error()))
			return
		}
	}
}

// RolePermissionHandlerDelete ..
func RolePermissionHandlerDelete() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, helpers.ResponseModelStruct(nil))
	}
}
