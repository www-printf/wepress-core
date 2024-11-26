package document

import (
	"github.com/labstack/echo/v4"
	"go.uber.org/dig"

	"github.com/www-printf/wepress-core/modules"
	"github.com/www-printf/wepress-core/modules/document/handlers"
	"github.com/www-printf/wepress-core/modules/document/repository"
	"github.com/www-printf/wepress-core/modules/document/usecases"
	"github.com/www-printf/wepress-core/pkg/middlewares"
	"github.com/www-printf/wepress-core/pkg/s3"
)

var Module modules.ModuleInstance = &DocumentModule{}

type DocumentModule struct{}

func (m *DocumentModule) RegisterRepositories(container *dig.Container) error {
	_ = container.Provide(repository.NewDocumentRepository)
	return nil
}

func (m *DocumentModule) RegisterUseCases(container *dig.Container) error {
	_ = container.Provide(s3.NewS3Client)
	_ = container.Provide(usecases.NewDocumentUsecase)
	return nil
}

func (m *DocumentModule) RegisterHandlers(g *echo.Group, container *dig.Container) error {
	return container.Invoke(func(
		documentUsecase usecases.DocumentUsecase,
		middlewareMngr middlewares.MiddlewareManager,
	) {
		handlers.NewDocumentHandler(g, documentUsecase, middlewareMngr)
	})
}
