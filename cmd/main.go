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
