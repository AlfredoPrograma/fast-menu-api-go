// FastMenu API
//
// FastMenu is a web application dedicated to majority of restaurants and food
// companies, which want to generate _amazing_ digital menus
//
// Schemes: http, https
// Version: 1.0.0
// Host: localhost:5000
// Base path: /api/v1
//
// Consumes:
// - application/json
//
// Produces:
// - application/json
//
// swagger:meta
package main

import (
	"fast-menu-api/config/db"
	"fast-menu-api/config/env"
	"fast-menu-api/logger"
	"fast-menu-api/server"
)

func main() {
	logger.Init()

	env.Load()
	db.Connect()

	server.Run()
}
