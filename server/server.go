package server

import (
	"fast-menu-api/config/env"
	"fast-menu-api/logger"
	"fast-menu-api/middlewares"
	"fast-menu-api/routes"
	"strings"

	"github.com/labstack/echo/v4"
)

func Run() {
	e := echo.New()

	e.Use(middlewares.HttpLogger())
	e.Use(middlewares.Recover())

	v1 := e.Group("/api/v1")

	routes.Load(v1)

	port := strings.Join([]string{":", env.Get("PORT")}, "")

	if err := e.Start(port); err != nil {
		logger.Use().
			Error().
			Err(err).
			Msg("")
	}
}
