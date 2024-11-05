package usecases

import (
	"context"

	"golang.org/x/crypto/bcrypt"

	jwtLib "github.com/golang-jwt/jwt/v5"

	"github.com/www-printf/wepress-core/modules/auth/dto"
	"github.com/www-printf/wepress-core/modules/auth/repository"
	"github.com/www-printf/wepress-core/pkg/constants"
	"github.com/www-printf/wepress-core/pkg/jwt"
)

type AuthUsecase interface {
	UserLogin(ctx context.Context, req *dto.LoginRequestBody) (*dto.AuthResponse, error)
}

type AuthUsecaseImpl struct {
	authRepo   repository.AuthRepository
	jwtManager jwt.JWTManager
}

func NewAuthUsecase(authRepo repository.AuthRepository, jwtMng jwt.JWTManager) AuthUsecase {
	return &AuthUsecaseImpl{
		authRepo:   authRepo,
		jwtManager: jwtMng,
	}
}

func (u *AuthUsecaseImpl) UserLogin(
	ctx context.Context, req *dto.LoginRequestBody) (*dto.AuthResponse, error) {
	user, err := u.authRepo.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return nil, constants.ErrNotFound
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return nil, constants.ErrUnauthorized
	}

	claims := jwtLib.MapClaims{
		"uid": user.ID.String(),
	}
	token, err := u.jwtManager.GenerateToken(claims)
	if err != nil {
		return nil, constants.ErrInternal
	}

	return &dto.AuthResponse{
		Token: token,
		Type:  "Bearer",
	}, nil

}
