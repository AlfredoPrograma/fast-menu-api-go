package main

import (
	"fast-menu-api/config/db"
	"fast-menu-api/config/env"
	"fast-menu-api/server"
)

func main() {
	env.Load()
	db.Connect()
	server.Run()
}
