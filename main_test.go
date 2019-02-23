package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetRequest(t *testing.T) {
	req, err := http.NewRequest("GET", "localhost:3005", nil)
	if err != nil {
		t.Fatalf("Error making GET request: %+v", err)
	}

	recorder := httptest.NewRecorder()

	simpleHandler(recorder, req)

	if recorder.Code != 200 {
		t.Errorf("Expected response code to be %d, but got %d", 200, recorder.Code)
	}
}