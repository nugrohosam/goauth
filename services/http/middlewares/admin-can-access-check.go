package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	helpers "github.com/nugrohosam/gosampleapi/helpers"
	usecases "github.com/nugrohosam/gosampleapi/usecases"
	"github.com/spf13/viper"
)

// AdminCanAccessAllKey ..
const AdminCanAccessAllKey = "admin-can-access-all"

// AdminCanAccessAll using for ..
func AdminCanAccessAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := helpers.GetAuth()
		adminRole := viper.GetString("role.admin")
		roles := []string{adminRole}
		isExists, err := usecases.IsHaveAnyRole(auth.ID, roles)
		if err != nil {
			c.JSON(http.StatusNotAcceptable, helpers.ResponseErr(err.Error()))
			c.Abort()
			return
		}

		c.Set(AdminCanAccessAllKey, isExists)
		c.Next()
	}
}

// AdminCannotAccessAll using for ..
func AdminCannotAccessAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := helpers.GetAuth()
		adminRole := viper.GetString("role.admin")
		roles := []string{adminRole}
		isExists, err := usecases.IsHaveAnyRole(auth.ID, roles)
		if err != nil {
			c.JSON(http.StatusNotAcceptable, helpers.ResponseErr(err.Error()))
			c.Abort()
			return
		}

		c.Set(AdminCanAccessAllKey, !isExists)
		c.Next()
	}
}
