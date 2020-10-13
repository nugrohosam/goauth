package http

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nugrohosam/gosampleapi/services/http/controllers"
	"github.com/nugrohosam/gosampleapi/services/http/middlewares"
)

// Serve using for listen to specific port
func Serve() error {
	routes := gin.New()

	routes.Use(gin.Logger())

	routes.Use(gin.CustomRecovery(func(c *gin.Context, recovered interface{}) {
		if err, ok := recovered.(string); ok {
			c.String(http.StatusInternalServerError, fmt.Sprintf("error: %s", err))
		}
		c.AbortWithStatus(http.StatusInternalServerError)
	}))

	authorize := routes.Group("/")

	authorize.Use(middlewares.AuthJwt())

	user := authorize.Group("/users")
	{
		user.POST("/", controllers.UserHandlerStore())
		user.POST("/detail", controllers.UserHandlerDetail())
	}

	if err := routes.Run(); err != nil {
		return err
	}

	return nil
}
