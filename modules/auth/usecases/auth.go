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
	"github.com/www-printf/wepress-core/pkg/crypto"
	"github.com/www-printf/wepress-core/pkg/jwt"
)

type AuthUsecase interface {
	UserLogin(ctx context.Context, req *dto.LoginRequestBody) (*dto.AuthResponseBody, error)
	ValidateToken(ctx context.Context, token string) (jwtLib.MapClaims, error)
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
	ctx context.Context, req *dto.LoginRequestBody) (*dto.AuthResponseBody, error) {
	user, err := u.authRepo.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return nil, constants.ErrNotFound
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return nil, constants.ErrUnauthorized
	}

	if user.PrivKey == "" {
		keyPair, err := crypto.GenerateKeyPair()
		if err != nil {
			return nil, constants.ErrInternal
		}
		err = u.authRepo.InsertKeyPair(ctx, user, keyPair)
		if err != nil {
			return nil, constants.ErrInternal
		}
	}

	privKeyBytes, err := base64.StdEncoding.DecodeString(user.PrivKey)
	if err != nil {
		return nil, constants.ErrInternal
	}
	claims := jwtLib.MapClaims{
		"uid": user.ID.String(),
	}
	token, err := u.tokenManger.Generate(claims, ed25519.PrivateKey(privKeyBytes))
	if err != nil {
		return nil, constants.ErrInternal
	}

	return &dto.AuthResponseBody{
		Token: token,
		Type:  "Bearer",
	}, nil
}

func (u *authUsecase) ValidateToken(
	ctx context.Context, token string) (jwtLib.MapClaims, error) {
	rawClaims, err := u.tokenManger.GetClaims(token)
	if err != nil {
		return nil, err
	}

	uid := rawClaims["uid"].(string)
	if uid == "" {
		return nil, constants.ErrUnauthorized
	}

	user, err := u.authRepo.GetUserByID(ctx, uid)
	if err != nil {
		return nil, constants.ErrNotFound
	}

	pubKeyBytes, err := base64.StdEncoding.DecodeString(user.PubKey)
	if err != nil {
		return nil, constants.ErrInternal
	}

	return u.tokenManger.Validate(token, ed25519.PublicKey(pubKeyBytes))
}
