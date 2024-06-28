package test

import (
	"encoding/json"
	"geo-jot/handler"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {

	req, err := http.NewRequest("GET", "/", nil)

	if err != nil {
		t.Fatal(err)
	}
	reqRecoder := httptest.NewRecorder()
	handler := http.HandlerFunc(handler.HealthCheck)

	handler.ServeHTTP(reqRecoder, req)

	if status := reqRecoder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := map[string]string{"message": "checked!"}
	var response map[string]string

	err = json.NewDecoder(reqRecoder.Body).Decode(&response)
	if err != nil {
		t.Fatal(err)
	}

	if response["message"] != expected["message"] {
		t.Errorf("handler returned unexpected body: got %v want %v",
			response["message"], expected["message"])
	}
	if contentType := reqRecoder.Header().Get("Content-Type"); contentType != "application/json" {
		t.Errorf("handler returned wrong content type: got %v want %v",
			contentType, "application/json")
	}

}
