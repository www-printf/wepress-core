package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/www-printf/wepress-core/modules/auth/dto"
	"github.com/www-printf/wepress-core/modules/auth/usecases"
	"github.com/www-printf/wepress-core/pkg/wrapper"
)

type AuthHandlerImpl struct {
	AuthUC usecases.AuthUsecase
}

func NewAuthHandler(g *echo.Group, authUC usecases.AuthUsecase) {
	h := &AuthHandlerImpl{
		AuthUC: authUC,
	}

	api := g.Group("auth")
	api.POST("/login", wrapper.Wrap(h.Login)).Name = "auth:login"
	api.POST("/verify", wrapper.Wrap(h.Verify)).Name = "auth:verify"
	api.POST("/forgot-password", wrapper.Wrap(h.ForgotPassword)).Name = "auth:forgot-password"
}

// @Summary Post Login
// @Description Post Login
// @Tags auth
// @Accept json
// @Produce json
// @Success      200  {object}  wrapper.SuccessResponse{data=domains.AuthResponse}
// @Security     Bearer
// @Router       /auth/login [post]
func (h *AuthHandlerImpl) Login(c echo.Context) wrapper.Response {
	auth := dto.AuthResponse{
		Token: "token",
		Type:  "Bearer",
	}
	return wrapper.Response{Data: auth, Status: http.StatusOK}
}

func (h *AuthHandlerImpl) Verify(c echo.Context) wrapper.Response {
	return wrapper.Response{Data: nil, Status: http.StatusOK}
}

func (h *AuthHandlerImpl) ForgotPassword(c echo.Context) wrapper.Response {
	return wrapper.Response{Data: nil, Status: http.StatusOK}
}
