package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/levor/to-do/internal/config"
	"github.com/levor/to-do/internal/http/handlers"
	"net/http"
)

func API(
	config *config.Config,
	authHandler *handlers.AuthHandler,
) *gin.Engine {
	r := gin.New()

	r.Use(
		gin.Recovery(),
		gin.Logger(),
	)

	apiGroup := r.Group("todo")
	privateGroup := apiGroup.Group("/private")
	{
		v1Group := privateGroup.Group("/v1")
		{
			v1Group.GET("/healthcheck", func(c *gin.Context) {
				c.JSON(http.StatusOK, gin.H{
					"Status":  http.StatusOK,
					"version": "1.0.0",
				})
			})
			v1Group.GET("/login", authHandler.Login)
			v1Group.GET("/refresh", authHandler.Refresh)
			v1Group.GET("/logout", authHandler.Logout)
		}
	}
	return r

}
