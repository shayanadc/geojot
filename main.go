package main

import (
	"geo-jot/config"
	"geo-jot/container"
	"geo-jot/db"
	"geo-jot/handler"
	"geo-jot/service"
	"net/http"
)

func init() {
	app := &App{}

	app.Setup()
}

type App struct {
	dbClient *db.Client
}

func (app *App) Setup() {
	_ = config.LoadEnv(".env")
	conn := db.NewDatabaseConnection()
	app.dbClient = db.NewClient(conn)
	container.GetContainer().SetDBClient(app.dbClient)
}

func (app *App) CloseDB() {
	if app.dbClient != nil {
		app.dbClient.Close()
	}
}

func main() {
	service.StoreConcurrently()
	http.HandleFunc("/health/check", handler.HealthCheck)
	_ = http.ListenAndServe(":8080", nil)
}
