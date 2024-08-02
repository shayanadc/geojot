package main

import (
	"fmt"
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
	app.dbClient = db.NewClient()
	container.GetContainer().SetDBClient(app.dbClient)
}

func (app *App) CloseDB() {
	if app.dbClient != nil {
		app.dbClient.Close()
	}
}

func main() {

	GetVehiclesWithNearestParcel := service.Find()

	fmt.Println(GetVehiclesWithNearestParcel)

	http.HandleFunc("/health/check", handler.HealthCheck)
	http.ListenAndServe(":8080", nil)
}
