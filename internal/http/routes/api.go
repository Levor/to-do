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
			authGroup := v1Group.Group("/auth")
			{
				authGroup.GET("/login", authHandler.Login)
				authGroup.GET("/refresh", authHandler.Refresh)
				authGroup.GET("/logout", authHandler.Logout)
				authGroup.GET("/signup", authHandler.Signup)
			}

		}
	}
	return r

}
