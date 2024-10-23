package config

import (
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gopkg.in/yaml.v3"
)

type Config struct {
	AppConfig AppConfig `yaml:"app"`
}

type AppConfig struct {
	Environment string `yaml:"environment"`
	Port        string `yaml:"port"`
	BaseURL     string `yaml:"base_url"`
	DatabaseDSN string `yaml:"database_dsn"`
	// RedisURL    string `yaml:"redis_url"`

	Validator  echo.Validator        `yaml:"-"`
	CORSConfig middleware.CORSConfig `yaml:"-"`

	AutoMigrate bool   `yaml:"auto_migrate"`
	LogLevel    string `yaml:"log_level"`

	AuthProvider string `yaml:"auth_provider"`
}

type AppValidator struct {
	validator *validator.Validate
}

func (cv *AppValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func InitConfig() (*Config, error) {
	yamlData, err := os.ReadFile("config/config.yaml")
	if err != nil {
		return nil, err
	}

	var config Config
	yaml.Unmarshal(yamlData, &config)
	return &config, nil
}
