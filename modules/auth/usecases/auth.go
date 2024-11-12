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
	UserLogin(ctx context.Context, req *dto.LoginRequestBody) (*dto.AuthResponseBody, *errors.Error)
	ValidateToken(ctx context.Context, token string) (jwtLib.MapClaims, *errors.Error)
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
	ctx context.Context, req *dto.LoginRequestBody) (*dto.AuthResponseBody, *errors.Error) {
	user, err := u.authRepo.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return nil, &errors.Error{LogError: err, HTTPError: constants.ErrNotFound}
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return nil, &errors.Error{LogError: err, HTTPError: constants.ErrUnauthorized}
	}

	if user.PrivKey == "" {
		keyPair, err := key.GenerateKeyPair()
		if err != nil {
			return nil, &errors.Error{LogError: err, HTTPError: constants.ErrInternal}
		}
		err = u.authRepo.InsertKeyPair(ctx, user, keyPair)
		if err != nil {
			return nil, &errors.Error{LogError: err, HTTPError: constants.ErrInternal}
		}
	}

	privKeyBytes, err := base64.StdEncoding.DecodeString(user.PrivKey)
	if err != nil {
		return nil, &errors.Error{LogError: err, HTTPError: constants.ErrInternal}
	}
	claims := jwtLib.MapClaims{
		"uid": user.ID.String(),
	}
	token, err := u.tokenManger.Generate(claims, ed25519.PrivateKey(privKeyBytes))
	if err != nil {
		return nil, &errors.Error{LogError: err, HTTPError: constants.ErrInternal}
	}

	return &dto.AuthResponseBody{
		Token: token,
		Type:  "Bearer",
	}, nil
}

func (u *authUsecase) ValidateToken(
	ctx context.Context, token string) (jwtLib.MapClaims, *errors.Error) {
	mapClaims, err := u.tokenManger.GetClaims(token)
	if err != nil {
		return nil, &errors.Error{LogError: err, HTTPError: constants.ErrInvalid}
	}

	uid := mapClaims["uid"].(string)
	if uid == "" {
		return nil, &errors.Error{LogError: nil, HTTPError: constants.ErrInvalid}
	}

	user, err := u.authRepo.GetUserByID(ctx, uid)
	if err != nil {
		return nil, &errors.Error{LogError: err, HTTPError: constants.ErrNotFound}
	}

	pubKeyBytes, err := base64.StdEncoding.DecodeString(user.PubKey)
	if err != nil {
		return nil, &errors.Error{LogError: err, HTTPError: constants.ErrInternal}
	}

	claims, err := u.tokenManger.Validate(token, ed25519.PublicKey(pubKeyBytes))
	if err != nil {
		return nil, &errors.Error{LogError: err, HTTPError: constants.ErrUnauthorized}
	}

	return claims, nil
}
