package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	helpers "github.com/nugrohosam/gosampleapi/helpers"
	usecases "github.com/nugrohosam/gosampleapi/usecases"
)

// CanAccessBy using for ..
func CanAccessBy(s []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := helpers.GetAuth()
		if isExists, err := usecases.IsHaveAnyRole(auth.ID, s); !isExists || err != nil {
			c.JSON(http.StatusNotAcceptable, helpers.ResponseErr("Cannot Access With Your Role"))
			c.Abort()
			return
		}

		c.Next()
	}
}
