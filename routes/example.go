package routes

import (
	"fast-menu-api/controllers"

	"github.com/labstack/echo/v4"
)

func ExampleRoutes(g *echo.Group) {
	r := g.Group("/example")

	r.GET("", controllers.HandleGet)
	r.POST("", controllers.HandlePost)
	r.PUT("", controllers.HandlePut)
	r.DELETE("", controllers.HandleDelete)
}
