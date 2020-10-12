package http

import (
	"github.com/gin-gonic/gin"
	"github.com/nugrohosam/gosampleapi/services/http/controllers"
)

// Serve using for listen to specific port
func Serve() error {
	routes := gin.Default()

	user := routes.Group("/users")
	{
		user.POST("/", controllers.UserHandlerStore())
		user.POST("/detail", controllers.UserHandlerDetail())
	}

	if err := routes.Run(); err != nil {
		return err
	}

	return nil
}
