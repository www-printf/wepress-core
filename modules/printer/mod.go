package printer

import (
	"github.com/labstack/echo/v4"
	"go.uber.org/dig"

	"github.com/www-printf/wepress-core/modules"
	"github.com/www-printf/wepress-core/modules/printer/handlers"
	"github.com/www-printf/wepress-core/modules/printer/repository"
	"github.com/www-printf/wepress-core/modules/printer/usecases"
	"github.com/www-printf/wepress-core/pkg/clusters"
	"github.com/www-printf/wepress-core/pkg/middlewares"
)

var Module modules.ModuleInstance = &PrinterModule{}

type PrinterModule struct{}

func (m *PrinterModule) RegisterRepositories(container *dig.Container) error {
	_ = container.Provide(repository.NewPrinterRepository)
	return nil
}

func (m *PrinterModule) RegisterUseCases(container *dig.Container) error {
	_ = container.Provide(clusters.NewClusterManager)
	_ = container.Provide(usecases.NewPrinterUsecase)
	return nil
}

func (m *PrinterModule) RegisterHandlers(g *echo.Group, container *dig.Container) error {
	return container.Invoke(func(
		printerUsecase usecases.PrinterUsecase,
		middlewareMngr middlewares.MiddlewareManager,
	) {
		handlers.NewPrinterHandler(g, printerUsecase, middlewareMngr)
	})
}
