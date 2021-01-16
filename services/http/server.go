package http

import (
	"time"

	sentrygin "github.com/getsentry/sentry-go/gin"

	"github.com/cnjack/throttle"
	"github.com/gin-gonic/contrib/gzip"
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

	isDebug := viper.GetBool("debug")
	if !isDebug {
		Routes.Use(exceptions.Recovery500())
	}

	rateLimiterCount := viper.GetUint64("rate-limiter.count")
	rateLimiterTime := viper.GetInt("rate-limiter.time-in-minutes")
	Routes.Use(throttle.Policy(&throttle.Quota{
		Limit:  rateLimiterCount,
		Within: time.Duration(rateLimiterTime) * time.Minute,
	}))

	Routes.Static("/assets", "./assets")
	Routes.Static("/web", "./web")

	Routes.Use(sentrygin.New(sentrygin.Options{
		Repanic: true,
	}))

	Routes.Any("/test-throttle", func(c *gin.Context) {
		c.Writer.Write([]byte("hello world"))
	})

	// test-sentry
	Routes.GET("/test-sentry", func(ctx *gin.Context) {
		panic("make panic test")
	})

	// v1
	v1 := Routes.Group("/v1")

	// v1/auth
	auth := v1.Group("/auth")
	auth.Use(gzip.Gzip(gzip.DefaultCompression))
	{
		auth.POST("/login", controllers.AuthHandlerLogin())
		auth.POST("/register", controllers.AuthHandlerRegister())
	}

	user := v1.Group("/user")
	user.Use(gzip.Gzip(gzip.DefaultCompression), middlewares.AuthJwt())
	{
		user.GET("/profile", controllers.UserHandlerShow())
		user.GET("/permissions", controllers.UserPermissionItsOwnHandlerIndex())
		user.GET("/roles", controllers.UserRoleItsOwnHandlerIndex())
	}

	userRole := v1.Group("/user-role")
	userRole.Use(gzip.Gzip(gzip.DefaultCompression), middlewares.AuthJwt())
	{
		userRole.GET("/", controllers.UserRoleHandlerIndex())
		userRole.POST("/", controllers.UserRoleHandlerCreate())
		userRole.PUT("/:id", controllers.UserRoleHandlerUpdate())
		userRole.DELETE("/:id", controllers.UserRoleHandlerDelete())
	}

	rolePermission := v1.Group("/role-permission")
	rolePermission.Use(gzip.Gzip(gzip.DefaultCompression), middlewares.AuthJwt())
	{

		retrieveRolePermission := rolePermission.Use(middlewares.CanAccessWith(
			[]string{
				viper.GetString("permission.role.retrieve"),
			},
		))
		{
			retrieveRolePermission.GET("/", controllers.RolePermissionHandlerIndex())
			retrieveRolePermission.GET("/:id", controllers.RolePermissionHandlerShow())
		}

		rolePermission.POST("/", controllers.RolePermissionHandlerCreate()).Use(middlewares.CanAccessWith(
			[]string{
				viper.GetString("permission.role.create"),
			},
		))
	}

	role := v1.Group("/role")
	role.Use(gzip.Gzip(gzip.DefaultCompression), middlewares.AuthJwt())
	{
		retrieveRole := role.Use(middlewares.CanAccessWith(
			[]string{
				viper.GetString("permission.role.retrieve"),
			},
		))
		{
			retrieveRole.GET("/", controllers.RoleHandlerIndex())
			retrieveRole.GET("/:id", controllers.RoleHandlerShow())
		}
	}

	permission := v1.Group("/permission")
	permission.Use(gzip.Gzip(gzip.DefaultCompression), middlewares.AuthJwt())
	{
		retrievePermission := permission.Use(middlewares.CanAccessWith(
			[]string{
				viper.GetString("permission.permission.retrieve"),
			},
		))
		{
			retrievePermission.GET("/", controllers.PermissionHandlerIndex())
			retrievePermission.GET("/:id", controllers.PermissionHandlerShow())
		}
	}
}
