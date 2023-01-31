package main

import (
	"fast-menu-api/config/env"
	"fast-menu-api/server"
)

func main() {
	env.Load()
	server.Run()
}
