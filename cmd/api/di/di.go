package di

import (
	"github.com/labstack/echo/v4"

	"go.uber.org/dig"

	"github.com/www-printf/wepress-core/config"
	"github.com/www-printf/wepress-core/infrastructure/datastore"
	"github.com/www-printf/wepress-core/modules"
	"github.com/www-printf/wepress-core/modules/auth"
	"github.com/www-printf/wepress-core/modules/demo"
	"github.com/www-printf/wepress-core/modules/document"
	"github.com/www-printf/wepress-core/pkg/middlewares"
)

func BuildDIContainer(
	conf *config.AppConfig,
) *dig.Container {
	container := dig.New()
	_ = container.Provide(func() *config.AppConfig {
		return conf
	})

	_ = container.Provide(func() string {
		return conf.DatabaseDSN
	})
	_ = container.Provide(datastore.ProvideDatabase)

	return container
}

func RegisterModules(e *echo.Group, container *dig.Container) error {
	var err error
	mapModules := map[string]modules.ModuleInstance{
		"demo":     demo.Module,
		"auth":     auth.Module,
		"document": document.Module,
	}

	gRoot := e.Group("/")
	for _, m := range mapModules {
		err = m.RegisterRepositories(container)
		if err != nil {
			return err
		}

		err = m.RegisterUseCases(container)
		if err != nil {
			return err
		}
	}

	err = container.Provide(middlewares.NewMiddlewareManager)

	for _, m := range mapModules {
		err = m.RegisterHandlers(gRoot, container)
		if err != nil {
			return err
		}
	}

	return err
}
