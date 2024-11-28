package middlewares

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"

	"github.com/www-printf/wepress-core/config"
	"github.com/www-printf/wepress-core/modules/auth/usecases"
)

type MiddlewareManager interface {
	Auth() echo.MiddlewareFunc
	Permission(allowedRoles ...string) echo.MiddlewareFunc
}

type middlewareManager struct {
	authUC  usecases.AuthUsecase
	appConf *config.AppConfig
}

func NewMiddlewareManager(authUC usecases.AuthUsecase, appConf *config.AppConfig) MiddlewareManager {
	return &middlewareManager{
		authUC:  authUC,
		appConf: appConf,
	}
}

func (m *middlewareManager) Auth() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			token, err := c.Cookie("token")
			if err != nil || token.Value == "" {
				authHeader := c.Request().Header.Get("Authorization")
				if !strings.HasPrefix(authHeader, "Bearer ") {
					return echo.NewHTTPError(http.StatusUnauthorized, "unauthorized")
				}
				token.Value = strings.TrimPrefix(authHeader, "Bearer ")
				if token.Value == "" {
					return echo.NewHTTPError(http.StatusUnauthorized, "unauthorized")
				}
			}
			claims, erro := m.authUC.ValidateToken(c.Request().Context(), token.Value)
			if erro != nil {
				return echo.NewHTTPError(http.StatusUnauthorized, "unauthorized")
			}
			c.Set("uid", claims["uid"])
			c.Set("role", claims["role"])
			return next(c)
		}
	}
}

func (m *middlewareManager) Permission(allowedRoles ...string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if len(allowedRoles) == 0 {
				allowedRoles = m.appConf.AllowedRoles
			}

			role, ok := c.Get("role").(string)
			if !ok {
				return echo.NewHTTPError(http.StatusUnauthorized, "unauthorized")
			}

			for _, allowedRole := range allowedRoles {
				if role == allowedRole {
					return next(c)
				}
			}

			return echo.NewHTTPError(http.StatusForbidden, "forbidden")
		}
	}
}
