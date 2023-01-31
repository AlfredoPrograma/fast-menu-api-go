package db

import (
	"fast-menu-api/config/env"
	"fast-menu-api/logger"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Handler *gorm.DB

func Connect() {
	user := env.Get("POSTGRES_USER")
	password := env.Get("POSTGRES_PASSWORD")
	database := env.Get("POSTGRES_DB")
	port := env.Get("POSTGRES_PORT")
	host := env.Get("HOST")

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, database,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		logger.Use().Error().
			Msg(err.Error())
	}

	Handler = db
}
