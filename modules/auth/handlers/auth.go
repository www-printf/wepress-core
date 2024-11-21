package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/www-printf/wepress-core/modules/auth/dto"
	"github.com/www-printf/wepress-core/modules/auth/usecases"
	"github.com/www-printf/wepress-core/pkg/constants"
	"github.com/www-printf/wepress-core/pkg/middlewares"
	"github.com/www-printf/wepress-core/pkg/wrapper"
)

type AuthHandler struct {
	authUC usecases.AuthUsecase
}

func NewAuthHandler(g *echo.Group, authUC usecases.AuthUsecase, middlewareMngr middlewares.MiddlewareManager) {
	h := &AuthHandler{
		authUC: authUC,
	}

	api := g.Group("auth")
	api.POST("/login", wrapper.Wrap(h.Login)).Name = "auth:login"
	api.POST("/verify", wrapper.Wrap(h.Validate)).Name = "auth:verify"
	api.POST("/forgot-password", wrapper.Wrap(h.ForgotPassword)).Name = "auth:forgot-password"

	api.Use(middlewareMngr.Auth())
	api.GET("/me", wrapper.Wrap(h.Profile)).Name = "get:profile"
}

// @Summary Post Login
// @Description Post Login
// @Tags auth
// @Accept json
// @Produce json
// @Param request body dto.LoginRequestBody true "Login Request Body"
// @Success      200  {object}  wrapper.SuccessResponse{data=dto.AuthResponseBody}
// @Failure      400  {object}  wrapper.FailResponse
// @Failure      401  {object}  wrapper.FailResponse
// @Failure      500  {object}  wrapper.FailResponse
// @Header 200 {string} Set-Cookie "token=jwt-token; Path=/; Secure"
// @Security     Bearer
// @Router       /auth/login [post]
func (h *AuthHandler) Login(c echo.Context) wrapper.Response {
	req := &dto.LoginRequestBody{}
	if err := c.Bind(req); err != nil {
		return wrapper.Response{Error: constants.HTTPBadRequest}
	}

	auth, err := h.authUC.UserLogin(c.Request().Context(), req)
	if err != nil {
		return wrapper.Response{Error: err}
	}

	cookie := &http.Cookie{
		Name:   "token",
		Value:  auth.Token,
		Path:   "/",
		Secure: true,
	}
	c.SetCookie(cookie)

	return wrapper.Response{Data: auth, Status: http.StatusOK}
}

// @Summary Verify Token
// @Description Verify Token
// @Tags auth
// @Accept json
// @Produce json
// @Param request body dto.VerifyTokenRequestBody true "Token to verify"
// @Success      200  {object}  wrapper.SuccessResponse{data=nil}
// @Failure      400  {object}  wrapper.FailResponse
// @Failure      401  {object}  wrapper.FailResponse
// @Security     Bearer
// @Router       /auth/verify [post]
func (h *AuthHandler) Validate(c echo.Context) wrapper.Response {
	token := &dto.VerifyTokenRequestBody{}
	if err := c.Bind(token); err != nil {
		return wrapper.Response{Error: constants.HTTPBadRequest}
	}

	_, err := h.authUC.ValidateToken(c.Request().Context(), token.Token)
	if err != nil {
		return wrapper.Response{Error: constants.HTTPUnauthorized}
	}

	return wrapper.Response{Data: nil, Status: http.StatusOK}
}

// @Summary Forgot Password
// @Description Request To Reset Password
// @Tags auth
// @Accept json
// @Produce json
// @Param request body dto.ForgotPasswordRequestBody true "Forgot Password Request Body"
// @Success      200  {object}  wrapper.SuccessResponse{data=nil}
// @Router       /auth/forgot-password [post]
func (h *AuthHandler) ForgotPassword(c echo.Context) wrapper.Response {
	return wrapper.Response{Data: nil, Status: http.StatusOK}
}

// @Summary Get User Profile
// @Description Get User Profile
// @Tags auth
// @Accept json
// @Produce json
// @Success      200  {object}  wrapper.SuccessResponse{data=dto.UserResponseBody}
// @Failure      401  {object}  wrapper.FailResponse
// @Failure      404  {object}  wrapper.FailResponse
// @Security     Bearer
// @Router       /auth/me [get]
func (h *AuthHandler) Profile(c echo.Context) wrapper.Response {
	uid := c.Get("uid").(string)
	user, err := h.authUC.GetMe(c.Request().Context(), uid)
	if err != nil {
		return wrapper.Response{Error: err}
	}
	return wrapper.Response{Data: user, Status: http.StatusOK}
}
