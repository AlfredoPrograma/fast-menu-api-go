package routes

import "github.com/labstack/echo/v4"

func Load(g *echo.Group) {
	ExampleRoutes(g)
}
