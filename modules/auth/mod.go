package auth

import (
	"github.com/labstack/echo/v4"
	"go.uber.org/dig"

	"github.com/www-printf/wepress-core/config"
	"github.com/www-printf/wepress-core/modules"
	"github.com/www-printf/wepress-core/modules/auth/handlers"
	"github.com/www-printf/wepress-core/modules/auth/repository"
	"github.com/www-printf/wepress-core/modules/auth/usecases"
)

var Module modules.ModuleInstance = &AuthModule{}

type AuthModule struct{}

func (m *AuthModule) RegisterRepositories(container *dig.Container) error {
	_ = container.Provide(repository.NewUserRepository)
	return nil
}

func (m *AuthModule) RegisterUseCases(container *dig.Container) error {
	_ = container.Provide(usecases.NewAuthUsecase)
	return nil
}

func (m *AuthModule) RegisterHandlers(g *echo.Group, container *dig.Container) error {
	return container.Invoke(func(
		appConf *config.AppConfig,
		authUsecase usecases.AuthUsecase,
	) {
		handlers.NewAuthHandler(g, authUsecase)
	})
}
