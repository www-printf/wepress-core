package usecases

import (
	"context"
	"slices"
	"time"

	"crypto/ed25519"
	"encoding/base64"

	jwtLib "github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

	"github.com/www-printf/wepress-core/config"
	"github.com/www-printf/wepress-core/modules/auth/domains"
	"github.com/www-printf/wepress-core/modules/auth/dto"
	"github.com/www-printf/wepress-core/modules/auth/repository"
	"github.com/www-printf/wepress-core/modules/auth/sessions"
	"github.com/www-printf/wepress-core/pkg/constants"
	"github.com/www-printf/wepress-core/pkg/errors"
	"github.com/www-printf/wepress-core/pkg/jwt"
	"github.com/www-printf/wepress-core/pkg/key"
	"github.com/www-printf/wepress-core/utils"
)

type AuthUsecase interface {
	UserLogin(ctx context.Context, req *dto.LoginRequestBody) (*dto.AuthResponseBody, *errors.HTTPError)
	ValidateToken(ctx context.Context, token string) (jwtLib.MapClaims, *errors.HTTPError)
	GetMe(ctx context.Context, uid string) (*dto.UserResponseBody, *errors.HTTPError)
	InitiateOAuth(ctx context.Context, provider string) (*dto.OauthResponseBody, *errors.HTTPError)
	HandleOAuthCallback(ctx context.Context, req *dto.OauthCallbackRequestBody) (*dto.AuthResponseBody, *errors.HTTPError)
}

type authUsecase struct {
	authRepo       repository.AuthRepository
	tokenManger    jwt.TokenManager
	sessionStorage sessions.SessionStorage
	oauthConfig    *config.OauthConfig
	oauthStrategy  OauthStrategy
}

func NewAuthUsecase(authRepo repository.AuthRepository, tokenMng jwt.TokenManager, sess sessions.SessionStorage, appConf *config.AppConfig) AuthUsecase {
	return &authUsecase{
		authRepo:       authRepo,
		tokenManger:    tokenMng,
		sessionStorage: sess,
		oauthConfig:    &appConf.OauthConfig,
	}
}

func (u *authUsecase) UserLogin(
	ctx context.Context, req *dto.LoginRequestBody) (*dto.AuthResponseBody, *errors.HTTPError) {
	user, err := u.authRepo.GetUserByEmail(ctx, req.Email)
	if err != nil || user == nil {
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
		"uid":  user.ID.String(),
		"role": user.Role,
	}
	token, err := u.tokenManger.Generate(claims, ed25519.PrivateKey(privKeyBytes))
	if err != nil {
		return nil, constants.HTTPInternal
	}

	sessUser := &domains.SessionUser{
		FullName:  user.FullName,
		Email:     user.Email,
		PubKey:    user.PubKey,
		CreatedAt: user.CreatedAt,
	}
	err = u.sessionStorage.SetUserSession(ctx, user.ID.String(), sessUser, u.tokenManger.GetExpireTime())
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

	var pubKeyStr string
	sess, err := u.sessionStorage.GetUserSession(ctx, uid)
	if err != nil {
		user, err := u.authRepo.GetUserByID(ctx, uid)
		if err != nil {
			return nil, constants.HTTPUnauthorized
		}
		pubKeyStr = user.PubKey
	} else {
		pubKeyStr = sess.PubKey
	}
	pubKeyBytes, err := base64.StdEncoding.DecodeString(pubKeyStr)
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
	user, err := u.sessionStorage.GetUserSession(ctx, uid)
	if err != nil {
		return nil, constants.HTTPUnauthorized
	}

	return &dto.UserResponseBody{
		ID:        uid,
		Email:     user.Email,
		Fullname:  user.FullName,
		CreatedAt: user.CreatedAt.String(),
		PubKey:    user.PubKey,
	}, nil
}

func (u *authUsecase) InitiateOAuth(ctx context.Context, provider string) (*dto.OauthResponseBody, *errors.HTTPError) {
	if !slices.Contains(u.oauthConfig.Providers, provider) {
		return nil, constants.HTTPBadRequest
	}

	if u.oauthStrategy == nil {
		switch provider {
		case "github":
			u.oauthStrategy = NewGithubOauthStrategy(&u.oauthConfig.Github)
		case "google":
			u.oauthStrategy = NewGoogleOauthStrategy(&u.oauthConfig.Google)
		case "facebook":
			u.oauthStrategy = NewFacebookOauthStrategy(&u.oauthConfig.Facebook)
		}
	}

	sess, err := u.oauthStrategy.GenerateOauthSession()
	if err != nil {
		return nil, constants.HTTPInternal
	}

	err = u.sessionStorage.SetOauthSession(ctx, sess, 10*time.Minute)
	if err != nil {
		return nil, constants.HTTPInternal
	}

	return &dto.OauthResponseBody{
		URL: sess.URL,
	}, nil
}

func (u *authUsecase) HandleOAuthCallback(
	ctx context.Context, req *dto.OauthCallbackRequestBody) (*dto.AuthResponseBody, *errors.HTTPError) {
	sess, err := u.sessionStorage.GetOauthSession(ctx, req.State)
	if err != nil {
		return nil, constants.HTTPBadRequest
	}

	if req.Error != "" {
		return nil, constants.HTTPBadRequest
	}

	if sess.Provider != req.Provider {
		return nil, constants.HTTPBadRequest
	}

	if u.oauthStrategy == nil {
		switch sess.Provider {
		case "github":
			u.oauthStrategy = NewGithubOauthStrategy(&u.oauthConfig.Github)
		case "google":
			u.oauthStrategy = NewGoogleOauthStrategy(&u.oauthConfig.Google)
		case "facebook":
			u.oauthStrategy = NewFacebookOauthStrategy(&u.oauthConfig.Facebook)
		}
	}

	oauthDTO := &dto.OauthCallBackTransfer{
		Code:     req.Code,
		Verifier: sess.Verifier,
	}
	tok, err := u.oauthStrategy.ExchangeToken(oauthDTO)
	if err != nil {
		return nil, constants.HTTPInternal
	}

	userInfo, err := u.oauthStrategy.GetUserInfo(tok)
	if err != nil {
		return nil, constants.HTTPInternal
	}

	user, err := u.authRepo.GetUserByEmail(ctx, userInfo.Email)
	if err != nil {
		return nil, constants.HTTPInternal
	}

	if user == nil {
		keyPair, err := key.GenerateKeyPair()
		if err != nil {
			return nil, constants.HTTPInternal
		}
		password, err := utils.GenerateSecret()
		if err != nil {
			return nil, constants.HTTPInternal
		}
		user = &domains.User{
			ID:       uuid.New(),
			Email:    userInfo.Email,
			FullName: "N/A",
			Password: password,
			PubKey:   keyPair["pubkey"],
			PrivKey:  keyPair["privkey"],
			Role:     "user",
		}
		err = u.authRepo.InsertUser(ctx, user)
		if err != nil {
			return nil, constants.HTTPInternal
		}
	}

	privKeyBytes, err := base64.StdEncoding.DecodeString(user.PrivKey)
	if err != nil {
		return nil, constants.HTTPInternal
	}
	claims := jwtLib.MapClaims{
		"uid":  user.ID.String(),
		"role": user.Role,
	}
	token, err := u.tokenManger.Generate(claims, ed25519.PrivateKey(privKeyBytes))
	if err != nil {
		return nil, constants.HTTPInternal
	}

	sessUser := &domains.SessionUser{
		FullName:  user.FullName,
		Email:     user.Email,
		PubKey:    user.PubKey,
		CreatedAt: user.CreatedAt,
	}
	err = u.sessionStorage.SetUserSession(ctx, user.ID.String(), sessUser, u.tokenManger.GetExpireTime())
	if err != nil {
		return nil, constants.HTTPInternal
	}

	return &dto.AuthResponseBody{
		Token: token,
		Type:  "Bearer",
	}, nil
}
