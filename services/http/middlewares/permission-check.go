package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	helpers "github.com/nugrohosam/gosampleapi/helpers"
	usecases "github.com/nugrohosam/gosampleapi/usecases"
	"github.com/spf13/viper"
)

// CanAccessWith using for ..
func CanAccessWith(permissionAccess []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := helpers.GetSessionDataString(c.Request, c.Writer, viper.GetString("session.userID"))
		if isPermited := usecases.CheckPermissionUser(userID, permissionAccess); !isPermited {
			c.JSON(http.StatusNotAcceptable, helpers.ResponseErr("Cannot Access With Your Permission"))
			c.Abort()
			return
		}

		c.Next()
	}
}
