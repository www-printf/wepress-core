package config

import (
	"time"

	"github.com/labstack/echo/v4/middleware"
)

func GetEchoLogConfig(appConfig *AppConfig) middleware.LoggerConfig {
	echoLogCnf := middleware.DefaultLoggerConfig
	echoLogCnf.CustomTimeFormat = time.RFC3339

	if appConfig.Environment == "production" {
		echoLogCnf.Format = `{"timestamp":"${time_custom}","level":"${level}","message":"${message}","service":"echo","id":"${id}","method":"${method}","uri":"${uri}","status":${status},"error":"${error}","latency":"${latency_human}"}`
	}

	return echoLogCnf
}
