package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	validator "github.com/go-playground/validator/v10"
	copier "github.com/jinzhu/copier"
	helpers "github.com/nugrohosam/gosampleapi/helpers"
	role "github.com/nugrohosam/gosampleapi/services/http/requests/v1"
	rolePerission "github.com/nugrohosam/gosampleapi/services/http/requests/v1"
	resource "github.com/nugrohosam/gosampleapi/services/http/resources/v1"
	usecases "github.com/nugrohosam/gosampleapi/usecases"
)

// RolePermissionHandlerIndex ..
func RolePermissionHandlerIndex() gin.HandlerFunc {
	return func(c *gin.Context) {
		var queryParams rolePerission.ListQuery
		c.BindQuery(&queryParams)

		rolePerissions, total, err := usecases.GetRolePermission(queryParams.Search, queryParams.PerPage, queryParams.Page, queryParams.OrderBy)

		if err != nil {
			c.JSON(http.StatusInternalServerError, helpers.ResponseErr(err.Error()))
			return
		}

		if cap(rolePerissions) > 0 {
			var listRolePermissionsResource = resource.RolePermissionListItems{}
			copier.Copy(&listRolePermissionsResource, &rolePerissions)
			if queryParams.Paginate {
				resourceData := helpers.BuildPaginate(queryParams.PerPage, queryParams.Page, total, &rolePerissions, &listRolePermissionsResource)
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
			rolePerission := usecases.ShowRolePermission(ID)
			var rolePerissionItem = resource.RolePermissionDetail{}
			copier.Copy(&rolePerissionItem, &rolePerission)
			if rolePerission.ID > 0 {
				c.JSON(http.StatusOK, helpers.ResponseOne(rolePerissionItem))
			} else {
				c.JSON(http.StatusOK, helpers.ResponseOne(nil))
			}
		}

	}
}

// RolePermissionHandlerCreate ..
func RolePermissionHandlerCreate() gin.HandlerFunc {
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

// RolePermissionHandlerUpdate ..
func RolePermissionHandlerUpdate() gin.HandlerFunc {
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

// RolePermissionHandlerDelete ..
func RolePermissionHandlerDelete() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, helpers.ResponseModelStruct(nil))
	}
}
