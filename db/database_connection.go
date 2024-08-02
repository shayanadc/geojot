package db

import (
	"os"
)

type DatabaseConnection struct {
	Url  string
	Name string
}

func NewDatabaseConnection() *DatabaseConnection {
	DB_URL := os.Getenv("DB_URL")
	DB_NAME := os.Getenv("DB_NAME")

	return &DatabaseConnection{
		Url:  DB_URL,
		Name: DB_NAME,
	}
}
