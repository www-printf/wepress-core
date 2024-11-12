package middlewares

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"

	"github.com/www-printf/wepress-core/modules/auth/usecases"
)

func Auth(authUC usecases.AuthUsecase) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			token, err := c.Cookie("token")
			if err != nil || token.Value == "" {
				authHeader := c.Request().Header.Get("Authorization")
				if !strings.HasPrefix(authHeader, "Bearer ") {
					return echo.NewHTTPError(http.StatusUnauthorized)
				}
				token.Value = strings.TrimPrefix(authHeader, "Bearer ")
				if token.Value == "" {
					return echo.NewHTTPError(http.StatusUnauthorized)
				}
			}
			claims, erro := authUC.ValidateToken(c.Request().Context(), token.Value)
			if erro != nil {
				return echo.NewHTTPError(http.StatusUnauthorized)
			}
			c.Set("uid", claims["uid"])

			return next(c)
		}
	}
}
