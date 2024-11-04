package usecases

import (
	"context"

	"github.com/www-printf/wepress-core/modules/auth/dto"
	"github.com/www-printf/wepress-core/modules/auth/repository"
)

type AuthUsecase interface {
	UserLogin(ctx context.Context, req dto.LoginRequestBody) (*dto.AuthResponse, error)
}

type AuthUsecaseImpl struct {
	UserRepo repository.UserRepository
}

func NewAuthUsecase(userRepo repository.UserRepository) AuthUsecase {
	return &AuthUsecaseImpl{
		UserRepo: userRepo,
	}
}

func (u *AuthUsecaseImpl) UserLogin(
	ctx context.Context,
	req dto.LoginRequestBody,
) (*dto.AuthResponse, error) {
	msg := dto.AuthResponse{
		Token: "token",
		Type:  "Bearer",
	}
	return &msg, nil
}
