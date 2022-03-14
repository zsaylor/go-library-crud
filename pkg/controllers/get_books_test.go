package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestGetBooks tests the GetBooks controller function
func TestGetBooks(t *testing.T) {
	// Define new GET request
	req, err := http.NewRequest("GET", "/api/v1/book/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Set up handler
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetBooks)

	// Check that we get correct response code
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check that we get the expected response body
	expected := `[{"id":1,"name":"Gravity's Rainbow","author":"Thomas Pynchon","publisher":"Vintage","genre":"World Historical Literature"}]`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}
