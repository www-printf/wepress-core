package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/www-printf/wepress-core/modules/auth/dto"
	"github.com/www-printf/wepress-core/modules/auth/usecases"
	"github.com/www-printf/wepress-core/pkg/middlewares"
	"github.com/www-printf/wepress-core/pkg/wrapper"
)

type AuthHandler struct {
	authUC usecases.AuthUsecase
}

func NewAuthHandler(g *echo.Group, authUC usecases.AuthUsecase, authMiddleware *middlewares.MiddlewareManager) {
	h := &AuthHandler{
		authUC: authUC,
	}

	api := g.Group("auth")
	api.POST("/login", wrapper.Wrap(h.Login)).Name = "auth:login"
	api.POST("/verify", wrapper.Wrap(h.Verify)).Name = "auth:verify"
	api.POST("/forgot-password", wrapper.Wrap(h.ForgotPassword)).Name = "auth:forgot-password"

	api.Use(authMiddleware.Auth())
	api.GET("/me", nil).Name = "get:me"
}

// @Summary Post Login
// @Description Post Login
// @Tags auth
// @Accept json
// @Produce json
// @Param request body dto.LoginRequestBody true "Login Request Body"
// @Success      200  {object}  wrapper.SuccessResponse{data=dto.AuthResponse}
// @Security     Bearer
// @Router       /auth/login [post]
func (h *AuthHandler) Login(c echo.Context) wrapper.Response {
	req := &dto.LoginRequestBody{}
	if err := c.Bind(req); err != nil {
		return wrapper.Response{Error: err, Status: http.StatusBadRequest}
	}

	auth, err := h.authUC.UserLogin(c.Request().Context(), req)
	if err != nil {
		return wrapper.Response{Error: err, Status: http.StatusUnauthorized}
	}

	return wrapper.Response{Data: auth, Status: http.StatusOK}
}

func (h *AuthHandler) Verify(c echo.Context) wrapper.Response {
	return wrapper.Response{Data: nil, Status: http.StatusOK}
}

func (h *AuthHandler) ForgotPassword(c echo.Context) wrapper.Response {
	return wrapper.Response{Data: nil, Status: http.StatusOK}
}
