package controllers

import (
	// "net/http"

	"net/http"

	"github.com/gin-gonic/gin"
	copier "github.com/jinzhu/copier"
	helpers "github.com/nugrohosam/gosampleapi/helpers"
	user "github.com/nugrohosam/gosampleapi/services/http/requests/v1"
	resource "github.com/nugrohosam/gosampleapi/services/http/resources/v1"
	usecases "github.com/nugrohosam/gosampleapi/usecases"
)

// UserHandlerIndex ..
func UserHandlerIndex() gin.HandlerFunc {
	return func(c *gin.Context) {
		var queryParams user.ListQuery
		c.BindQuery(&queryParams)

		users, total, err := usecases.GetUser(queryParams.Search, queryParams.PerPage, queryParams.Page, queryParams.OrderBy)

		if err != nil {
			c.JSON(http.StatusInternalServerError, helpers.ResponseErr(err.Error()))
			return
		}

		if cap(users) > 0 {
			var listUsersResource = resource.UserListItems{}
			copier.Copy(&listUsersResource, &users)
			if queryParams.Paginate {
				resourceData := helpers.BuildPaginate(queryParams.PerPage, queryParams.Page, total, &users, &listUsersResource)
				c.JSON(http.StatusOK, helpers.ResponseModelStruct(resourceData))
			} else {
				c.JSON(http.StatusOK, helpers.ResponseMany(listUsersResource))
			}
		} else {
			c.JSON(http.StatusOK, helpers.ResponseMany(nil))
		}
	}
}

// UserHandlerShow ..
func UserHandlerShow() gin.HandlerFunc {
	return func(c *gin.Context) {

		auth := helpers.GetAuth()
		user := usecases.ShowUser(auth.ID)

		userItem := resource.UserDetail{}
		copier.Copy(&userItem, &user)
		if user.ID > 0 {
			c.JSON(http.StatusOK, helpers.ResponseOne(userItem))
		} else {
			c.JSON(http.StatusOK, helpers.ResponseOne(nil))
		}
	}
}

// UserRoleItsOwnHandlerIndex ..
func UserRoleItsOwnHandlerIndex() gin.HandlerFunc {
	return func(c *gin.Context) {

		auth := helpers.GetAuth()
		userRoles := usecases.GetUserRoleWithUserID(auth.ID)

		if cap(userRoles) > 0 {
			listUserRolesResource := resource.UserRoleItsOwnListItems{}
			copier.Copy(&listUserRolesResource, &userRoles)
			c.JSON(http.StatusOK, helpers.ResponseMany(listUserRolesResource))
		} else {
			c.JSON(http.StatusOK, helpers.ResponseMany(nil))
		}
	}
}

// UserPermissionOwnedHandlerIndex ..
func UserPermissionOwnedHandlerIndex() gin.HandlerFunc {
	return func(c *gin.Context) {

		auth := helpers.GetAuth()
		userUsers := usecases.GetPermissionWithUserID(auth.ID)

		if cap(userUsers) > 0 {
			listUserRolesResource := resource.PermissionListItems{}
			copier.Copy(&listUserRolesResource, &userUsers)
			c.JSON(http.StatusOK, helpers.ResponseMany(listUserRolesResource))
		} else {
			c.JSON(http.StatusOK, helpers.ResponseMany(nil))
		}
	}
}
