package app

import (
	"admin/internal/middleware"
	"admin/pkg/jwt"
	"admin/pkg/response"
	"time"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
)

func (a *App) initRoutes() {
	JWTMaker := jwt.NewJWTMaker()
	JWTAuthMiddleware := middleware.NewJWTAuthMiddleware(JWTMaker, a.repos.User, a.rdb)
	AuthMiddleware := JWTAuthMiddleware.AuthMiddleware()
	// 404
	a.router.NoRoute(func(c *gin.Context) {
		response.FailWithCode(c, response.NOT_FOUND)
	})
	// a.router.GET("/", func(c *gin.Context) {
	// 	c.Redirect(response.MOVED_PERMANENTLY, "https://"+constant.SITE)
	// })
	v1 := a.router.Group("/v1")
	{
		// 版本信息
		v1.GET("/version", a.handlers.App.Version)
		v1.GET("/setup", a.handlers.App.Setup)
		// 用户相关路由
		v1.POST("/signup", a.handlers.User.Signup)
		v1.POST("/signin", a.handlers.User.Signin)
		v1.GET("/email/isexists/:email", a.handlers.User.IsEmailExists)
		v1.POST("/email/code", a.handlers.User.SendCode)
		v1.GET("/google/:access_token", a.handlers.User.GoogleSignin)

		// 用户相关路由
		users := v1.Group("/user")
		users.Use(AuthMiddleware)
		{
			users.GET("/info", a.handlers.User.GetUser)
			users.GET("/routes", a.handlers.User.GetRouters)
			users.GET("/signout", a.handlers.User.Signout)
		}

		lookup := v1.Group("/lookup")
		lookup.Use(AuthMiddleware)
		{
			lookup.GET("/groups", a.handlers.Lookup.Groups)
			lookup.GET("/group/:group_value", a.handlers.Lookup.Group)
			lookup.PATCH("/status/:id/:status", a.handlers.Lookup.Status)
			lookup.PUT("/sort", a.handlers.Lookup.Sort)
			lookup.POST("", a.handlers.Lookup.POST)
			lookup.DELETE("/:id", a.handlers.Lookup.DELETE)
			lookup.PUT("/:id", a.handlers.Lookup.PUT)
			lookup.GET("/:id", a.handlers.Lookup.GET)
		}
		permissions := v1.Group("/permissions")
		permissions.Use(AuthMiddleware)
		{
			permissions.PATCH("/status/:id/:status", a.handlers.Permissions.Status)
			permissions.GET("/parent/:type", a.handlers.Permissions.ParentType)
			permissions.DELETE("/:id", a.handlers.Permissions.DELETE)
			permissions.PUT("/:id", a.handlers.Permissions.PUT)
			permissions.GET("/:id", a.handlers.Permissions.GET)
			permissions.GET("", a.handlers.Permissions.LIST)
			permissions.POST("", a.handlers.Permissions.POST)
		}

	}

	if gin.IsDebugging() {
		pprof.Register(v1)
	}
}

func (a *App) initMiddleware() {
	// 全局中间件
	a.router.Use(gin.Logger())
	a.router.Use(gin.Recovery())
	config := middleware.RateLimitConfig{
		MaxRequests:     100,       // 每秒100个请求
		MaxQueryParams:  10,        // 最多10个查询参数
		MaxParamLength:  500,       // 参数最大长度
		CleanupInterval: time.Hour, // 每小时清理一次
	}
	middleware, _ := middleware.NewRateLimitMiddleware(config)
	a.router.Use(middleware)
	// 其他全局中间件...
}
