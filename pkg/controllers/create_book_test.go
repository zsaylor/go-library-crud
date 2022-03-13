package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestCreateBook tests the CreateBook controller function
func TestCreateBook(t *testing.T) {
	// Create a new book entry
	var newBook = []byte(`{"Name": "Gravity's Rainbow", "Author": "Thomas Pynchon", "Publisher": "Vintage", "Genre": "World Historical Literature"}`)

	// Define new POST request
	req, err := http.NewRequest("POST", "/api/v1/book/", bytes.NewBuffer(newBook))
	if err != nil {
		t.Fatal(err)
	}

	// Set up handler
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateBook)

	// Check that we get correct response code
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}


	var expected map[string]interface{}
	json.Unmarshal(rr.Body.Bytes(), &expected)

	if expected["name"] != "Gravity's Rainbow" {
		t.Errorf("expected book name 'Gravity's Rainbow' got '%v'", expected["name"])
	}

	if expected["author"] != "Thomas Pynchon" {
		t.Errorf("expected author 'Thomas Pynchon' got '%v'", expected["author"])
	}

	if expected["publisher"] != "Vintage" {
		t.Errorf("expected publisher 'Vintage' got '%v'", expected["publisher"])
	}

	if expected["genre"] != "World Historical Literature" {
		t.Errorf("expected genre 'World Historical Literature' got '%v'", expected["genre"])
	}
}
