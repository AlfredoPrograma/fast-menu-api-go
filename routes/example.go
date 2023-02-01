package routes

import (
	"fast-menu-api/controllers"

	"github.com/labstack/echo/v4"
)

func ExampleRoutes(g *echo.Group) {
	r := g.Group("/example")

	// swagger:route GET /example Example getAllExamples
	//
	// Lists all examples
	r.GET("", controllers.HandleGet)

	// swagger:route POST /example Example createExample
	//
	// Creates a new example
	r.POST("", controllers.HandlePost)
	r.PUT("", controllers.HandlePut)
	r.DELETE("", controllers.HandleDelete)
}
