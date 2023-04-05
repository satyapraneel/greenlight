package main

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestHealthCheck(t *testing.T) {
	//go test -timeout 30s -run ^TestHealthCheck$ github.com/satyapraneel/greenlight/cmd/api -count=1
	// go test -timeout 30s -run ./.. github.com/satyapraneel/greenlight/cmd/api -count=1``
	req, err := http.NewRequest(http.MethodGet, "/v1/healthcheck", nil)
	if err != nil {
		t.Fatal(err)
	}

	var cfg config
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	cfg.port = "4001"
	cfg.env = "testing"

	app := &application{
		config: cfg,
		logger: logger,
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(app.healthcheckHandler)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	healthCheckResponse := make(map[string]any)

	json.Unmarshal(rr.Body.Bytes(), &healthCheckResponse)

	if healthCheckResponse["env"] != cfg.env {
		t.Errorf("handler returned unexpected env: got %v want %v",
			healthCheckResponse["env"], cfg.env)
	}
}
