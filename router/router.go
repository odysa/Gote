package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/handler"
	"server/router/middlewares"
)

func Load(g *gin.Engine, wares ...gin.HandlerFunc) *gin.Engine {
	g.Use(gin.Recovery())
	g.Use(gin.Logger())
	g.Use(wares...)
	g.Use(middlewares.Options)
	g.Use(middlewares.Secure)

	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "Invalid Address")
	})

	// health check
	hc := g.Group("/hc")
	{
		hc.GET("/health", handler.HealthCheck)
	}

	return g
}
