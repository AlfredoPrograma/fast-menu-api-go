package server

import (
	"fast-menu-api/config/env"
	"fast-menu-api/routes"
	"strings"

	"github.com/labstack/echo/v4"
)

func Run() {
	e := echo.New()

	v1 := e.Group("/api/v1")

	routes.Load(v1)

	port := strings.Join([]string{":", env.Get("PORT")}, "")
	e.Logger.Fatal(e.Start(port))
}