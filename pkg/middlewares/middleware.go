package middlewares

import (
	"github.com/labstack/echo/v4"

	"github.com/www-printf/wepress-core/config"
)

type MiddlewareManager struct {
	appConf *config.AppConfig
}

func NewMiddlewareManager(appConf *config.AppConfig) *MiddlewareManager {
	return &MiddlewareManager{
		appConf: appConf,
	}
}

func (m MiddlewareManager) Auth() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			return next(c)
		}
	}
}
