package router

import (
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/odysa/Gote/controller"
	"github.com/odysa/Gote/router/middlewares"
	"log"
)

func Load(g *gin.Engine, wares ...gin.HandlerFunc) *gin.Engine {
	g.Use(gin.Recovery())
	g.Use(gin.Logger())
	g.Use(middlewares.Validator())
	g.Use(wares...)
	g.Use(middlewares.Options())
	g.Use(middlewares.Secure())

	g.NoRoute(controller.PageNotFound)

	store, err := sessions.NewRedisStore(10, "tcp", "localhost:6379", "", []byte("secret"))
	if err != nil {
		log.Fatalln("unable to connect to redis store")
	}

	// health check
	hc := g.Group("/hc")
	{
		hc.GET("/health", controller.HealthCheck)
	}

	loginGroup := g.Group("/login")
	loginGroup.Use(
		sessions.Sessions("adminSession", store))
	{
		loginGroup.POST("/admin", controller.AdminLogin)
	}

	adminGroup := g.Group("/admin")
	adminGroup.Use(
		sessions.Sessions("adminSession", store),
		middlewares.SessionAuth(),
	)
	{
		adminGroup.GET("/info", controller.AdminInfo)
		adminGroup.GET("/logout", controller.AdminLogout)
		adminGroup.POST("/change_password", controller.AdminChangePassword)
	}
	serviceGroup := g.Group("/service")
	serviceGroup.Use(
		sessions.Sessions("adminSession", store),
		middlewares.SessionAuth(),
	)
	{
		serviceGroup.GET("list", controller.ServiceList)
	}
	return g
}
