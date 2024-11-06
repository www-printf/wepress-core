package jwt

import (
	"crypto/ed25519"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/rs/zerolog/log"
	"github.com/www-printf/wepress-core/config"
	"github.com/www-printf/wepress-core/pkg/constants"
)

type TokenManager interface {
	Generate(claims jwt.MapClaims, privateKey ed25519.PrivateKey) (string, error)
	Validate(token string, publicKey ed25519.PublicKey) (jwt.MapClaims, error)
	GetClaims(token string) (jwt.MapClaims, error)
}

type tokenManager struct {
	expiredTime time.Duration
	issuer      string
}

func NewTokenManager(appConf *config.AppConfig) TokenManager {
	return &tokenManager{
		expiredTime: time.Hour * 24,
		issuer:      appConf.Issuer,
	}
}

func (j *tokenManager) Generate(claims jwt.MapClaims, privateKey ed25519.PrivateKey) (string, error) {
	claims["exp"] = time.Now().Add(j.expiredTime).Unix()
	claims["iat"] = time.Now().Unix()
	claims["iss"] = j.issuer

	token := jwt.NewWithClaims(jwt.SigningMethodEdDSA, claims)
	signedToken, err := token.SignedString(privateKey)
	if err != nil {
		log.Error().Err(err).Msg("failed to sign token")
		return "", err
	}

	return signedToken, nil
}

func (j *tokenManager) Validate(token string, publicKey ed25519.PublicKey) (jwt.MapClaims, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return publicKey, nil
	})
	if err != nil || !parsedToken.Valid {
		log.Error().Err(err).Msg("failed to validate token")
		return nil, constants.ErrUnauthorized
	}
	exp, err := parsedToken.Claims.GetExpirationTime()
	if err != nil {
		log.Error().Err(err).Msg("failed to get expiration time")
		return nil, constants.ErrUnauthorized
	}
	if time.Now().After(exp.Time) {
		log.Error().Msg("token is expired")
		return nil, constants.ErrExpired
	}

	uid := parsedToken.Claims.(jwt.MapClaims)["uid"]
	if uid == nil {
		log.Error().Msg("uid is missing in token")
		return nil, constants.ErrInvalid
	}

	return parsedToken.Claims.(jwt.MapClaims), nil
}

func (j *tokenManager) GetClaims(token string) (jwt.MapClaims, error) {
	parsedToken, _, _ := new(jwt.Parser).ParseUnverified(token, jwt.MapClaims{})
	mapClaims := parsedToken.Claims.(jwt.MapClaims)
	return mapClaims, nil
}
