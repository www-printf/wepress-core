package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/rs/zerolog/log"
	"github.com/www-printf/wepress-core/config"
	"github.com/www-printf/wepress-core/modules/auth/dto"
)

type JWTManager interface {
	GenerateToken(claims jwt.MapClaims) (dto.Token, error)
	ValidateToken(token string) (jwt.Claims, error)
}

type JWTManagerImpl struct {
	secretKey   string
	expiredTime time.Duration
	issuer      string
}

func NewJWTManager(appConf *config.AppConfig) JWTManager {
	secretKey := "secret"
	expiredTime := time.Hour * 24
	issuer := "wepress"
	return &JWTManagerImpl{
		secretKey:   secretKey,
		expiredTime: expiredTime,
		issuer:      issuer,
	}
}

func (j *JWTManagerImpl) GenerateToken(claims jwt.MapClaims) (dto.Token, error) {
	claims["exp"] = time.Now().Add(j.expiredTime).Unix()
	claims["iat"] = time.Now().Unix()
	claims["iss"] = j.issuer

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(j.secretKey))
	if err != nil {
		log.Error().Err(err).Msg("failed to sign token")
		return "", err
	}

	return dto.Token(signedToken), nil
}

func (j *JWTManagerImpl) ValidateToken(token string) (jwt.Claims, error) {
	return nil, nil
}
