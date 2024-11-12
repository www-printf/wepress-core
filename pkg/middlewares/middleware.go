package middlewares

import (
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"

	"github.com/www-printf/wepress-core/modules/auth/usecases"
	"github.com/www-printf/wepress-core/pkg/constants"
)

func Auth(authUC usecases.AuthUsecase) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			token, err := c.Cookie("token")
			if err != nil {
				return err
			}
			if token.Value == "" {
				authHeader := c.Request().Header.Get("Authorization")
				if !strings.HasPrefix(authHeader, "Bearer ") {
					return constants.ErrUnauthorized
				}
				token.Value = strings.TrimPrefix(authHeader, "Bearer ")
				if token.Value == "" {
					log.Error().Msg("token is empty")
					return constants.ErrUnauthorized
				}
			}

			claims, errs := authUC.ValidateToken(c.Request().Context(), token.Value)
			if errs != nil {
				return errs.HTTPError
			}
			c.Set("uid", claims["uid"])

			return next(c)
		}
	}
}
