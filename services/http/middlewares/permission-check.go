package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	helpers "github.com/nugrohosam/gosampleapi/helpers"
	usecases "github.com/nugrohosam/gosampleapi/usecases"
)

// CanAccessWith using for ..
func CanAccessWith(permissionAccess []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		accessAll := c.GetBool(AdminCanAccessAllKey)
		if accessAll {
			c.Next()
			return
		}

		auth := helpers.GetAuth()
		if isPermited := usecases.CheckPermissionUser(auth.ID, permissionAccess); !isPermited {
			c.JSON(http.StatusNotAcceptable, helpers.ResponseErr("Cannot Access With Your Permission"))
			c.Abort()
			return
		}

		c.Next()
	}
}
