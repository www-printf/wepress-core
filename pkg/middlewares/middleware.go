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
			token := ""
			for _, cookie := range c.Cookies() {
				if cookie.Name == "auth" {
					token = cookie.Value
					break
				}
			}
			if token == "" {
				authHeader := c.Request().Header.Get("Authorization")
				if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
					return constants.ErrUnauthorized
				}
				token = strings.TrimPrefix(authHeader, "Bearer ")
				if token == "" {
					return constants.ErrUnauthorized
				}
			}

			log.Info().Str("token", token).Msg("token")

			claims, err := authUC.ValidateToken(c.Request().Context(), token)
			if err != nil {
				return err
			}
			c.Set("uid", claims["uid"])

			return next(c)
		}
	}
}
