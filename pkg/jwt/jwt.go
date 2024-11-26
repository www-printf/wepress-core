package jwt

import (
	"crypto/ed25519"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/www-printf/wepress-core/config"
)

type TokenManager interface {
	Generate(claims jwt.MapClaims, privateKey ed25519.PrivateKey) (string, error)
	Validate(token string, publicKey ed25519.PublicKey) (jwt.MapClaims, error)
	GetClaims(token string) (jwt.MapClaims, error)
	GetExpireTime() time.Duration
}

type tokenManager struct {
	expiredTime time.Duration
	issuer      string
}

func NewTokenManager(appConf *config.AppConfig) TokenManager {
	return &tokenManager{
		expiredTime: time.Duration(appConf.TokenExpire) * time.Second,
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
		return "", err
	}

	return signedToken, nil
}

func (j *tokenManager) Validate(token string, publicKey ed25519.PublicKey) (jwt.MapClaims, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return publicKey, nil
	})
	if err != nil || !parsedToken.Valid {
		return nil, err
	}
	exp, err := parsedToken.Claims.GetExpirationTime()
	if err != nil {
		return nil, err
	}
	if time.Now().After(exp.Time) {
		return nil, err
	}

	uid := parsedToken.Claims.(jwt.MapClaims)["uid"]
	if uid == nil {
		return nil, errors.New("cannot find uid in token")
	}

	return parsedToken.Claims.(jwt.MapClaims), nil
}

func (j *tokenManager) GetClaims(token string) (jwt.MapClaims, error) {
	parsedToken, _, _ := new(jwt.Parser).ParseUnverified(token, jwt.MapClaims{})
	mapClaims := parsedToken.Claims.(jwt.MapClaims)
	return mapClaims, nil
}

func (j *tokenManager) GetExpireTime() time.Duration {
	return j.expiredTime
}
