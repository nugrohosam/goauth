package controllers

import (
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

		permissions, total, err := usecases.GetPermission(queryParams.Search, queryParams.PerPage, queryParams.Page, queryParams.OrderBy)

		if err != nil {
			c.JSON(http.StatusInternalServerError, helpers.ResponseErr(err.Error()))
			return
		}

		if cap(permissions) > 0 {
			var listPermissionsResource = resource.PermissionListItems{}
			copier.Copy(&listPermissionsResource, &permissions)
			if queryParams.Paginate {
				resourceData := helpers.BuildPaginate(queryParams.PerPage, queryParams.Page, total, &permissions, &listPermissionsResource)
				c.JSON(http.StatusOK, helpers.ResponseModelStruct(resourceData))
			} else {
				c.JSON(http.StatusOK, helpers.ResponseMany(listPermissionsResource))
			}
		} else {
			c.JSON(http.StatusOK, helpers.ResponseMany(nil))
		}
	}
}

// PermissionHandlerShow ..
func PermissionHandlerShow() gin.HandlerFunc {
	return func(c *gin.Context) {
		ID := c.Param("id")

		if len(ID) < 1 {
			c.JSON(http.StatusBadRequest, helpers.ResponseErr("params id should filled"))
		} else {
			IDInt, _ := strconv.Atoi(ID)
			permission := usecases.ShowPermission(IDInt)
			var permissionItem = resource.PermissionDetail{}
			copier.Copy(&permissionItem, &permission)
			if permission.ID > 0 {
				c.JSON(http.StatusOK, helpers.ResponseOne(permissionItem))
			} else {
				c.JSON(http.StatusOK, helpers.ResponseOne(nil))
			}
		}

	}
}

// PermissionHandlerCreate ..
func PermissionHandlerCreate() gin.HandlerFunc {
	return func(c *gin.Context) {
		var permission permission.CreateRole
		c.BindJSON(&permission)

		validate := helpers.NewValidation()
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

		validate := helpers.NewValidation()
		if err := validate.Struct(permission); err != nil {
			validationErrors := err.(validator.ValidationErrors)
			fieldsErrors := helpers.TransformValidations(validationErrors)
			c.JSON(http.StatusBadRequest, helpers.ResponseErrValidation(fieldsErrors))
			return
		}

		roleID := c.Param("id")
		roleIDInt, _ := strconv.Atoi(roleID)

		if err := usecases.UpdateRole(roleIDInt, permission.Name); err != nil {
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
