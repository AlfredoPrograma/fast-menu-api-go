package middlewares

import (
	"fast-menu-api/logger"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/pkg/errors"
)

func Recover() echo.MiddlewareFunc {
	return middleware.RecoverWithConfig(middleware.RecoverConfig{
		LogErrorFunc: func(c echo.Context, err error, stack []byte) error {
			logger.Use().
				Error().
				Stack().
				Err(errors.New((err.Error()))).
				Msg("")

			return nil
		},
	})
}
