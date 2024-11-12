package usecases

import (
	"context"

	"crypto/ed25519"
	"encoding/base64"

	jwtLib "github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"

	"github.com/www-printf/wepress-core/modules/auth/dto"
	"github.com/www-printf/wepress-core/modules/auth/repository"
	"github.com/www-printf/wepress-core/pkg/constants"
	"github.com/www-printf/wepress-core/pkg/errors"
	"github.com/www-printf/wepress-core/pkg/jwt"
	"github.com/www-printf/wepress-core/pkg/key"
)

type AuthUsecase interface {
	UserLogin(ctx context.Context, req *dto.LoginRequestBody) (*dto.AuthResponseBody, *errors.HTTPError)
	ValidateToken(ctx context.Context, token string) (jwtLib.MapClaims, *errors.HTTPError)
	GetMe(ctx context.Context, uid string) (*dto.UserResponseBody, *errors.HTTPError)
}

type authUsecase struct {
	authRepo    repository.AuthRepository
	tokenManger jwt.TokenManager
}

func NewAuthUsecase(authRepo repository.AuthRepository, tokenMng jwt.TokenManager) AuthUsecase {
	return &authUsecase{
		authRepo:    authRepo,
		tokenManger: tokenMng,
	}
}

func (u *authUsecase) UserLogin(
	ctx context.Context, req *dto.LoginRequestBody) (*dto.AuthResponseBody, *errors.HTTPError) {
	user, err := u.authRepo.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return nil, constants.HTTPUnauthorized
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return nil, constants.HTTPUnauthorized
	}

	if user.PrivKey == "" {
		keyPair, err := key.GenerateKeyPair()
		if err != nil {
			return nil, constants.HTTPInternal
		}
		err = u.authRepo.InsertKeyPair(ctx, user, keyPair)
		if err != nil {
			return nil, constants.HTTPInternal
		}
	}

	privKeyBytes, err := base64.StdEncoding.DecodeString(user.PrivKey)
	if err != nil {
		return nil, constants.HTTPInternal
	}
	claims := jwtLib.MapClaims{
		"uid": user.ID.String(),
	}
	token, err := u.tokenManger.Generate(claims, ed25519.PrivateKey(privKeyBytes))
	if err != nil {
		return nil, constants.HTTPInternal
	}

	return &dto.AuthResponseBody{
		Token: token,
		Type:  "Bearer",
	}, nil
}

func (u *authUsecase) ValidateToken(
	ctx context.Context, token string) (jwtLib.MapClaims, *errors.HTTPError) {
	mapClaims, err := u.tokenManger.GetClaims(token)
	if err != nil {
		return nil, constants.HTTPUnauthorized
	}

	uid := mapClaims["uid"].(string)
	if uid == "" {
		return nil, constants.HTTPUnauthorized
	}

	user, err := u.authRepo.GetUserByID(ctx, uid)
	if err != nil {
		return nil, constants.HTTPNotFound
	}

	pubKeyBytes, err := base64.StdEncoding.DecodeString(user.PubKey)
	if err != nil {
		return nil, constants.HTTPInternal
	}

	claims, err := u.tokenManger.Validate(token, ed25519.PublicKey(pubKeyBytes))
	if err != nil {
		return nil, constants.HTTPUnauthorized
	}

	return claims, nil
}

func (u *authUsecase) GetMe(ctx context.Context, uid string) (*dto.UserResponseBody, *errors.HTTPError) {
	user, err := u.authRepo.GetUserByID(ctx, uid)
	if err != nil {
		return nil, constants.HTTPNotFound
	}

	return &dto.UserResponseBody{
		ID:        user.ID.String(),
		Email:     user.Email,
		Fullname:  user.FullName,
		CreatedAt: user.CreatedAt.String(),
		PubKey:    user.PubKey,
	}, nil
}
