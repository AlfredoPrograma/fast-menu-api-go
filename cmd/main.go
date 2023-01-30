package main

import (
	"fast-menu-api/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	v1 := e.Group("/api/v1")

	routes.Load(v1)

	e.Logger.Fatal(e.Start(":5000"))
}
