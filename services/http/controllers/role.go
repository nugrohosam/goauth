package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	validator "github.com/go-playground/validator/v10"
	copier "github.com/jinzhu/copier"
	helpers "github.com/nugrohosam/gosampleapi/helpers"
	permission "github.com/nugrohosam/gosampleapi/services/http/requests/v1"
	role "github.com/nugrohosam/gosampleapi/services/http/requests/v1"
	resource "github.com/nugrohosam/gosampleapi/services/http/resources/v1"
	usecases "github.com/nugrohosam/gosampleapi/usecases"
)

// RoleHandlerIndex ..
func RoleHandlerIndex() gin.HandlerFunc {
	return func(c *gin.Context) {
		var queryParams permission.ListQuery
		c.BindQuery(&queryParams)

		roles, total, err := usecases.GetRole(queryParams.Search, queryParams.PerPage, queryParams.Page, queryParams.OrderBy)

		if err != nil {
			c.JSON(http.StatusInternalServerError, helpers.ResponseErr(err.Error()))
			return
		}

		if cap(roles) > 0 {
			var listRolesResource = resource.RoleListItems{}
			copier.Copy(&listRolesResource, &roles)
			if queryParams.Paginate {
				resourceData := helpers.BuildPaginate(queryParams.PerPage, queryParams.Page, total, &roles, &listRolesResource)
				c.JSON(http.StatusOK, helpers.ResponseModelStruct(resourceData))
			} else {
				c.JSON(http.StatusOK, helpers.ResponseMany(listRolesResource))
			}
		} else {
			c.JSON(http.StatusOK, helpers.ResponseMany(nil))
		}
	}
}

// RoleHandlerShow ..
func RoleHandlerShow() gin.HandlerFunc {
	return func(c *gin.Context) {
		ID := c.Param("id")

		if len(ID) < 1 {
			c.JSON(http.StatusBadRequest, helpers.ResponseErr("params id should filled"))
		} else {
			IDInt, _ := strconv.Atoi(ID)
			permission := usecases.ShowRole(IDInt)
			var permissionItem = resource.RoleDetail{}
			copier.Copy(&permissionItem, &permission)
			if permission.ID > 0 {
				c.JSON(http.StatusOK, helpers.ResponseOne(permissionItem))
			} else {
				c.JSON(http.StatusOK, helpers.ResponseOne(nil))
			}
		}
	}
}

// RoleHandlerCreate ..
func RoleHandlerCreate() gin.HandlerFunc {
	return func(c *gin.Context) {
		var role role.CreateRole
		c.BindJSON(&role)

		validate := helpers.NewValidation()
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
		var role role.UpdateRole
		c.BindJSON(&role)

		validate := helpers.NewValidation()
		if err := validate.Struct(role); err != nil {
			validationErrors := err.(validator.ValidationErrors)
			fieldsErrors := helpers.TransformValidations(validationErrors)
			c.JSON(http.StatusBadRequest, helpers.ResponseErrValidation(fieldsErrors))
			return
		}

		roleID := c.Param("id")
		roleIDInt, _ := strconv.Atoi(roleID)

		if err := usecases.UpdateRole(roleIDInt, role.Name); err != nil {
			c.JSON(http.StatusBadRequest, helpers.ResponseErr(err.Error()))
			return
		}
	}
}

// RoleHandlerDelete ..
func RoleHandlerDelete() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, helpers.ResponseModelStruct(nil))
	}
}
