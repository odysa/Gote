package router

import (
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/odysa/Gote/controller"
	"github.com/odysa/Gote/router/middlewares"
	"log"
	"net/http"
)

func Load(g *gin.Engine, wares ...gin.HandlerFunc) *gin.Engine {
	g.Use(gin.Recovery())
	g.Use(gin.Logger())
	g.Use(middlewares.Validator())
	g.Use(wares...)
	g.Use(middlewares.Options())
	g.Use(middlewares.Secure())

	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "Invalid Address")
	})

	store, err := sessions.NewRedisStore(10, "tcp", "localhost:6379", "", []byte("secret"))

	// health check
	hc := g.Group("/hc")
	{
		hc.GET("/health", controller.HealthCheck)
	}
	if err != nil {
		log.Fatalln("unable to connect to redis store")
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
	return g
}
