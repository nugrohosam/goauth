package http

import (
	"github.com/gin-gonic/gin"
	"github.com/nugrohosam/gosampleapi/services/http/controllers"
	"github.com/nugrohosam/gosampleapi/services/http/middlewares"
)

// Serve using for listen to specific port
func Serve() *gin.Engine {
	routes := gin.New()
	routes.Use(gin.Logger())
	routes.Use(gin.Recovery())

	// v1
	v1 := routes.Group("/v1")

	// v1/auth
	auth := v1.Group("/auth")
	{
		auth.POST("/login", controllers.AuthHandlerLogin())
		auth.POST("/register", controllers.AuthHandlerRegister())
	}

	// v1/hello-with-middleware
	helloWithMiddleware := v1.Group("/hello-with-middleware")
	helloWithMiddleware.Use(middlewares.AuthJwt())
	{
		helloWithMiddleware.GET("/", controllers.HelloWorldHandler())
	}

	helloWithoutMiddleware := v1.Group("/hello-without-middleware")
	{
		helloWithoutMiddleware.GET("/", controllers.HelloWorldHandler())
	}

	if err := routes.Run(); err != nil {
		panic("Cannot Serve")
	}

	return routes
}
