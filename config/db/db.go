package db

import (
	"fast-menu-api/config/env"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	user := env.Get("POSTGRES_USER")
	password := env.Get("POSTGRES_PASSWORD")
	database := env.Get("POSTGRES_DB")
	port := env.Get("POSTGRES_PORT")
	host := env.Get("HOST")

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, database,
	)

	// TODO: Add error handling
	db, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	return db
}
