package demo

import (
	"github.com/labstack/echo/v4"
	"github.com/www-printf/wepress-core/config"
	"github.com/www-printf/wepress-core/modules"
	"github.com/www-printf/wepress-core/modules/demo/handlers"
	"github.com/www-printf/wepress-core/modules/demo/usecases"
	"go.uber.org/dig"
)

var Module modules.ModuleInstance = &DemoModule{}

type DemoModule struct{}

func (m *DemoModule) RegisterRepositories(container *dig.Container) error {
	return nil
}

func (m *DemoModule) RegisterUseCases(container *dig.Container) error {
	_ = container.Provide(usecases.NewDemoUsecase)
	return nil
}

func (m *DemoModule) RegisterHandlers(g *echo.Group, container *dig.Container) error {
	return container.Invoke(func(
		appConf *config.AppConfig,
		demoUsecase usecases.DemoUsecase,
	) {
		handlers.NewDemoHandler(g, demoUsecase)
	})
}
