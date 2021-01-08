package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	helpers "github.com/nugrohosam/gosampleapi/helpers"
	usecases "github.com/nugrohosam/gosampleapi/usecases"
	"github.com/spf13/viper"
)

// CanAccessWith using for ..
func CanAccessWith(s []string) gin.HandlerFunc {

	globalAccessOpen := viper.GetString("permission.global.open")
	globalAccessClose := viper.GetString("permission.global.close")

	return func(c *gin.Context) {
		userID := helpers.GetSessionDataString(c.Request, c.Writer, viper.GetString("session.userID"))
		if isExists, err := usecases.IsHaveAnyPermission(userID, []string{globalAccessOpen}); !isExists || err != nil {
			c.JSON(http.StatusNotAcceptable, helpers.ResponseErr("Cannot Access With Your Permission"))
			c.Abort()
			return
		} else if isExists, err := usecases.IsHaveAnyPermission(userID, []string{globalAccessClose}); isExists || err != nil {
			c.JSON(http.StatusNotAcceptable, helpers.ResponseErr("Cannot Access With Your Permission"))
			c.Abort()
			return
		} else if isExists, err := usecases.IsHaveAnyPermission(userID, s); !isExists || err != nil {
			c.JSON(http.StatusNotAcceptable, helpers.ResponseErr("Cannot Access With Your Permission"))
			c.Abort()
			return
		}

		c.Next()
	}
}
