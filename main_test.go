package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// 測試gin功能是否正常
// 測試CI
// 測試CI2
func TestHomepageHandler(t *testing.T) {
	mockResponse := `{"message":"Welcome to the Tech Company listing API with Golang"}`
	r := SetUpRouter()
	r.GET("/", HomepageHandler)

	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a response recorder
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body
	// expected := `{"message":"pong"}`
	if rr.Body.String() != mockResponse {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), mockResponse)
	}
}
