package config

import (
	"os"
	"strings"

	"github.com/caarlos0/env/v11"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Config struct {
	AppConfig AppConfig
}

type AppConfig struct {
	Environment string `env:"ENV"`
	Host        string `env:"APP_HOST"`
	Port        string `env:"APP_PORT"`
	BaseURL     string `env:"APP_BASE_URL"`
	DatabaseDSN string `env:"DATABASE_DSN"`
	// RedisURL    string `env:"REDIS_URI"`

	AutoMigrate bool   `env:"AUTO_MIGRATE"`
	LogLevel    string `env:"LOG_LEVEL"`

	Validator  echo.Validator
	CORSConfig middleware.CORSConfig

	Issuer string `env:"JWT_ISS" envDefault:"wepress"`

	S3Config S3Config
}

type S3Config struct {
	Region          string `env:"S3_REGION"`
	BucketName      string `env:"S3_BUCKET_NAME"`
	EndPoint        string `env:"S3_ENDPOINT"`
	AccessKey       string `env:"S3_ACCESS_KEY"`
	SecretKey       string `env:"S3_SECRET_KEY"`
	PresignedExpire int    `env:"S3_PRESIGNED_EXPIRE" envDefault:"600"`
	MaxSize         int64  `env:"S3_MAX_UPLOAD_SIZE" envDefault:"104857600"`
}

type AppValidator struct {
	validator *validator.Validate
}

func (cv *AppValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func InitConfig() (*Config, error) {
	var cfg Config
	if err := env.Parse(&cfg); err != nil {
		return nil, err
	}
	allowOrigins := os.Getenv("CORS_ALLOW_ORIGINS")
	if allowOrigins == "" {
		allowOrigins = "*"
	}
	var AllowOrigins []string
	if strings.Contains(allowOrigins, ",") {
		AllowOrigins = strings.Split(allowOrigins, ",")
	} else {
		AllowOrigins = []string{allowOrigins}
	}

	cfg.AppConfig.CORSConfig = middleware.CORSConfig{
		AllowOrigins: AllowOrigins,
	}

	return &cfg, nil
}
