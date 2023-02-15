package di

import (
	"github.com/levor/to-do/internal/config"
	"github.com/levor/to-do/internal/db"
	"github.com/levor/to-do/internal/http/handlers"
	"github.com/levor/to-do/internal/http/routes"
	"github.com/levor/to-do/internal/repositories"
	"go.uber.org/dig"
	"log"
)

var Container *dig.Container

func init() {
	Container = dig.New()
	providers := []interface{}{
		config.Read,
		routes.API,
		db.NewConnection,
	}

	providers = append(providers, repositories.RepositoryProvider()...)
	providers = append(providers, handlers.HandlerProvider()...)

	for _, provider := range providers {
		err := Container.Provide(provider)
		if err != nil {
			log.Fatal(err)
		}
	}
}
