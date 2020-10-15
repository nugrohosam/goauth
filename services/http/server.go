package http

import (
	"github.com/gin-gonic/gin"
	"github.com/nugrohosam/gosampleapi/services/http/controllers"
	"github.com/nugrohosam/gosampleapi/services/http/middlewares"
)

// Serve using for listen to specific port
func Serve() error {
	routes := gin.New()
	routes.Use(gin.Logger())
	routes.Use(gin.Recovery())

	// v1
	v1 := routes.Group("/v1")
	auth := v1.Group("/auth")
	{
		auth.POST("/login", controllers.AuthHandlerLogin())
	}

	home := v1.Group("/users")
	home.Use(middlewares.AuthJwt())
	{
		home.GET("/", controllers.UserHandlerIndex())
	}

	if err := routes.Run(); err != nil {
		return err
	}

	return nil
}
