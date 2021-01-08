package http

import (
	sentrygin "github.com/getsentry/sentry-go/gin"

	"github.com/gin-gonic/gin"
	"github.com/nugrohosam/gosampleapi/services/http/controllers"
	"github.com/nugrohosam/gosampleapi/services/http/exceptions"
	"github.com/spf13/viper"

	"github.com/nugrohosam/gosampleapi/services/http/middlewares"
)

// Routes ...
var Routes *gin.Engine

// Serve using for listen to specific port
func Serve() error {
	Prepare()

	port := viper.GetString("app.port")
	if err := Routes.Run(":" + port); err != nil {
		return err
	}

	return nil
}

// Prepare ...
func Prepare() {

	Routes = gin.New()
	Routes.Use(exceptions.Recovery500())
	Routes.Static("/assets", "./assets")
	Routes.Use(sentrygin.New(sentrygin.Options{
		Repanic: true,
	}))

	// test-sentry
	Routes.GET("/test-sentry", func(ctx *gin.Context) {
		panic("make panic test")
	})

	// v1
	v1 := Routes.Group("/v1")

	// v1/auth
	auth := v1.Group("/auth")
	{
		auth.POST("/login", controllers.AuthHandlerLogin())
		auth.POST("/register", controllers.AuthHandlerRegister())
	}

	userRole := v1.Group("/user-role")
	userRole.Use(middlewares.AuthJwt())
	{
		userRole.GET("/", controllers.UserRoleHandlerIndex())
		userRole.POST("/", controllers.UserRoleHandlerCreate())
		userRole.PUT("/:id", controllers.UserRoleHandlerUpdate())
		userRole.DELETE("/:id", controllers.UserRoleHandlerDelete())
	}

	rolePermission := v1.Group("/role-permission")
	rolePermission.Use(middlewares.AuthJwt())
	{
		rolePermission.GET("/", controllers.RolePermissionHandlerIndex())
		rolePermission.POST("/", controllers.RolePermissionHandlerCreate())
		rolePermission.PUT("/:id", controllers.RolePermissionHandlerUpdate())
		rolePermission.DELETE("/:id", controllers.RolePermissionHandlerDelete())
	}

	permission := v1.Group("/permission")
	permission.Use(middlewares.AuthJwt())
	{
		permission.GET("/", controllers.PermissionHandlerIndex()).Use(middlewares.CanAccessWith(
			[]string{
				viper.GetString("permission.permission.retrieve"),
			},
		))
	}
}
