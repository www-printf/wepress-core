package auth_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"handlers"

	"github.com/www-printf/wepress-core/modules/auth/dto"
	"github.com/www-printf/wepress-core/modules/auth/usecases/mocks"
	"github.com/www-printf/wepress-core/pkg/wrapper"
)

func TestAuthHandler_Login(t *testing.T) {
	e := echo.New()
	mockAuthUC := new(mocks.AuthUsecase)
	handler := &handlers.AuthHandler{AuthUC: mockAuthUC}

	reqBody := &dto.LoginRequestBody{
		Email:    "test@example.com",
		Password: "password",
	}
	expectedResp := &dto.AuthResponseBody{
		Token: "mock-token",
	}

	mockAuthUC.On("UserLogin", mock.Anything, reqBody).Return(expectedResp, nil)

	reqBytes, _ := json.Marshal(reqBody)
	req := httptest.NewRequest(http.MethodPost, "/auth/login", bytes.NewReader(reqBytes))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)

	resp := handler.Login(ctx)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, wrapper.Response{Data: expectedResp, Status: http.StatusOK}, resp)
	mockAuthUC.AssertExpectations(t)
}

func TestAuthHandler_Validate(t *testing.T) {
	e := echo.New()
	mockAuthUC := new(mocks.AuthUsecase)
	handler := &handlers.AuthHandler{AuthUC: mockAuthUC}

	reqBody := &dto.VerifyTokenRequestBody{
		Token: "valid-token",
	}
	mockAuthUC.On("ValidateToken", mock.Anything, reqBody.Token).Return(nil, nil)

	reqBytes, _ := json.Marshal(reqBody)
	req := httptest.NewRequest(http.MethodPost, "/auth/verify", bytes.NewReader(reqBytes))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)

	resp := handler.Validate(ctx)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, wrapper.Response{Data: nil, Status: http.StatusOK}, resp)
	mockAuthUC.AssertExpectations(t)
}

func TestAuthHandler_Profile(t *testing.T) {
	e := echo.New()
	mockAuthUC := new(mocks.AuthUsecase)
	handler := &handlers.AuthHandler{AuthUC: mockAuthUC}

	expectedUser := &dto.UserResponseBody{
		ID:    "user-id",
		Email: "test@example.com",
	}
	mockAuthUC.On("GetMe", mock.Anything, "user-id").Return(expectedUser, nil)

	req := httptest.NewRequest(http.MethodGet, "/auth/me", nil)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	ctx.Set("uid", "user-id")

	resp := handler.Profile(ctx)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, wrapper.Response{Data: expectedUser, Status: http.StatusOK}, resp)
	mockAuthUC.AssertExpectations(t)
}

func TestAuthHandler_RequestOauth(t *testing.T) {
	e := echo.New()
	mockAuthUC := new(mocks.AuthUsecase)
	handler := &handlers.AuthHandler{AuthUC: mockAuthUC}

	provider := "google"
	expectedResp := &dto.OauthResponseBody{
		URL: "https://oauth.example.com/google",
	}
	mockAuthUC.On("InitiateOAuth", mock.Anything, provider).Return(expectedResp, nil)

	req := httptest.NewRequest(http.MethodGet, "/oauth/google", nil)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	ctx.SetParamNames("provider")
	ctx.SetParamValues(provider)

	resp := handler.RequestOauth(ctx)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, wrapper.Response{Data: expectedResp, Status: http.StatusOK}, resp)
	mockAuthUC.AssertExpectations(t)
}

func TestAuthHandler_HandleCallBack(t *testing.T) {
	e := echo.New()
	mockAuthUC := new(mocks.AuthUsecase)
	handler := &handlers.AuthHandler{AuthUC: mockAuthUC}

	reqBody := &dto.OauthCallbackRequestBody{
		Code:  "auth-code",
		State: "state",
	}
	expectedResp := &dto.AuthResponseBody{
		Token: "mock-token",
	}
	mockAuthUC.On("HandleOAuthCallback", mock.Anything, reqBody).Return(expectedResp, nil)

	reqBytes, _ := json.Marshal(reqBody)
	req := httptest.NewRequest(http.MethodPost, "/oauth/callback", bytes.NewReader(reqBytes))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)

	resp := handler.HandleCallBack(ctx)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, wrapper.Response{Data: expectedResp, Status: http.StatusOK}, resp)
	mockAuthUC.AssertExpectations(t)
}
