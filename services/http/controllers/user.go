package controllers

import (
	// "net/http"

	"net/http"

	"github.com/gin-gonic/gin"
	copier "github.com/jinzhu/copier"
	helpers "github.com/nugrohosam/gosampleapi/helpers"
	resource "github.com/nugrohosam/gosampleapi/services/http/resources/v1"
	usecases "github.com/nugrohosam/gosampleapi/usecases"
)

// UserHandlerShow ..
func UserHandlerShow() gin.HandlerFunc {
	return func(c *gin.Context) {

		auth := helpers.GetAuth()
		user := usecases.ShowUser(auth.ID)

		var userItem = resource.UserDetail{}
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
			var listUserRolesResource = resource.UserRoleItsOwnListItems{}
			copier.Copy(&listUserRolesResource, &userRoles)
			c.JSON(http.StatusOK, helpers.ResponseMany(listUserRolesResource))
		} else {
			c.JSON(http.StatusOK, helpers.ResponseMany(nil))
		}
	}
}

// UserPermissionItsOwnHandlerIndex ..
func UserPermissionItsOwnHandlerIndex() gin.HandlerFunc {
	return func(c *gin.Context) {

		auth := helpers.GetAuth()
		userPermissions := usecases.GetUserPermissionWithUserID(auth.ID)

		if cap(userPermissions) > 0 {
			var listUserRolesResource = resource.UserRoleItsOwnListItems{}
			copier.Copy(&listUserRolesResource, &userPermissions)
			c.JSON(http.StatusOK, helpers.ResponseMany(listUserRolesResource))
		} else {
			c.JSON(http.StatusOK, helpers.ResponseMany(nil))
		}
	}
}
