package main

import (
	"geo-jot/handler"
	"net/http"
)

func main() {
	http.HandleFunc("/health/check", handler.HealthCheck)
	http.ListenAndServe(":8080", nil)
}
