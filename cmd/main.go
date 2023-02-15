package main

import (
	"github.com/gin-gonic/gin"
	"github.com/levor/to-do/internal/config"
	"github.com/levor/to-do/internal/di"
	"log"
)

func main()  {
	c := di.Container

	err := c.Invoke(func(
		api *gin.Engine,
		config *config.Config,
	) {
		err := api.Run(":" + config.ServerPort)
		if err != nil {
			log.Panic(err)
		}
	})
	if err != nil {
		log.Panic(err)
	}
}
