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
