package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog/log"
	echoSwagger "github.com/swaggo/echo-swagger"
	"github.com/www-printf/wepress-core/cmd/api/di"
	"github.com/www-printf/wepress-core/config"
	_ "github.com/www-printf/wepress-core/docs"
	"github.com/www-printf/wepress-core/pkg/clusters"
)

// @title WePress API
// @version 1.0
// @description This is the API Document for WePress
// @termsOfService http://swagger.io/terms/

// @license.name MIT
// @license.url http://opensource.org/licenses/MIT

// @BasePath /api/v1
// @securityDefinitions.apiKey	AccessToken
// @in header
// @name Authorization
// @description Enter the token with the `Bearer ` prefix, e.g. `Bearer jwt_token_string`.

func main() {
	cfg, err := config.InitConfig()
	if err != nil {
		panic(err)
	}

	if log.Log() != nil {
		defer log.Log().Send()
	}

	e := echo.New()

	e.Use(middleware.LoggerWithConfig(config.GetEchoLogConfig(&cfg.AppConfig)))
	e.Use(middleware.Recover())
	e.Use(middleware.RequestID())
	e.HideBanner = true
	e.Validator = cfg.AppConfig.Validator

	container := di.BuildDIContainer(&cfg.AppConfig)

	api := e.Group("/api/v1")
	api.GET("/ping", func(c echo.Context) error {
		return c.String(200, "pong")
	})

	err = di.RegisterModules(api, container)
	if err != nil {
		log.Fatal().Msgf("Error when registering modules: %v", err)
	}

	e.GET("/docs", func(c echo.Context) error {
		return c.Redirect(http.StatusPermanentRedirect, "/docs/index.html")
	})
	e.GET("/docs/*", echoSwagger.WrapHandler)

	go func() {
		if err := e.Start(
			fmt.Sprintf("%s:%s", cfg.AppConfig.Host, cfg.AppConfig.Port),
		); err != nil {
			log.Fatal().Msgf("Error when starting server: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	container.Invoke(func(cm clusters.ClusterManager) { cm.Close() })
	if err := e.Shutdown(ctx); err != nil {
		log.Fatal().Msgf("Error when shutting down server: %v", err)
	}
}
