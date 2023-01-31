package logger

import (
	"fmt"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog"
)

const LOGS_FILENAME = "app.log"

var logger zerolog.Logger

func openFileWriter() *os.File {
	fileWriter, err := os.OpenFile(LOGS_FILENAME, os.O_APPEND|os.O_WRONLY, 0644)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return fileWriter
}

func Init() {
	fileWriter := openFileWriter()

	outputs := zerolog.MultiLevelWriter(fileWriter, os.Stdout)

	logger = zerolog.New(outputs).With().Timestamp().Logger()
}

func HttpMiddleware() echo.MiddlewareFunc {
	return middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogStatus: true,
		LogURI:    true,
		LogMethod: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			logger.Info().
				Str("method", v.Method).
				Str("uri", v.URI).
				Int("status", v.Status).
				Msg("request")

			return nil
		},
	})
}

func Use() *zerolog.Logger {
	return &logger
}
