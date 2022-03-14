package controllers

import (
	"go-library/pkg/utils"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetBookById(t *testing.T) {
	// Define new GET request
	var m = make(map[string]string)
	m["bookId"] = "1"
	req := utils.NewRequest("GET", "/api/v1/book/", m)

	// Set up handler
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetBookById)

	// Check that we get correct response code
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `{"id":1,"name":"Gravity's Rainbow","author":"Thomas Pynchon","publisher":"Vintage","genre":"World Historical Literature"}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

