// Controllers_test contains all tests for the API endpoints within the controllers package.
package controllers

import (
	"bytes"
	"encoding/json"
	"go-library/pkg/models"
	"go-library/pkg/utils"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

// TestMain opens a new connection to the test database and resets the table 'books'.
func TestMain(m *testing.M) {
	err := models.InitDB("root:password@tcp(127.0.0.1:3306)/test-lib?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatalf("failed to open database: %v", err)
	}

	db := models.GetDB()
	db.Exec("DELETE FROM books;")
	db.Exec("ALTER TABLE books AUTO_INCREMENT = 1;")

	exitVal := m.Run()
	os.Exit(exitVal)
}

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

	// Check that returned book has the correct fields
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

// TestGetBookById tests the GetBookById controller function.
func TestGetBookById(t *testing.T) {
	// Define new GET request
	var m = make(map[string]string)
	m["bookId"] = "1"
	req := utils.NewRequest("GET", "/api/v1/book/", "", m)

	// Set up handler
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetBookById)

	// Check that we get correct response code
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check that we get expected response body
	expected := `{"id":1,"name":"Gravity's Rainbow","author":"Thomas Pynchon","publisher":"Vintage","genre":"World Historical Literature"}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

// TestGetBookByIdNotFound tests that we get the correct response code (400) when searching a nonexistent book ID.
func TestGetBookByIdNotFound(t *testing.T) {
	// Define new GET request with non-existent ID
	var m = make(map[string]string)
	m["bookId"] = "99"
	req := utils.NewRequest("GET", "/api/v1/book/", "", m)

	// Set up handler
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetBookById)
	handler.ServeHTTP(rr, req)

	// Check that we return correct status code
	if status := rr.Code; status == http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}
}

// TestUpdateBook tests that a POST call successfully edits a book entry's information.
func TestUpdateBook(t *testing.T) {
	// Define new POST request
	var updateBook = `{"Name": "V", "Author": "Thomas Pynchon", "Publisher": "Vintage", "Genre": "World Historical Literature"}`
	var m = make(map[string]string)
	m["bookId"] = "1"
	req := utils.NewRequest("GET", "/api/v1/book/", updateBook, m)

	// Set up handler
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(UpdateBook)
	handler.ServeHTTP(rr, req)

	// Check that we get correct response code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check that we get expected response body
	expected := `{"id":1,"name":"V","author":"Thomas Pynchon","publisher":"Vintage","genre":"World Historical Literature"}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

// TestDeleteBook tests that a DELETE call successfully returns an empty entry.
func TestDeleteBook(t *testing.T) {
	// Define new DELETE request
	var m = make(map[string]string)
	m["bookId"] = "1"
	req := utils.NewRequest("DELETE", "/api/v1/book/", "", m)

	// Set up handler
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(DeleteBook)
	handler.ServeHTTP(rr, req)

	// Check that we get correct response code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check that we get expected response body
	expected := `{"id":0,"name":"","author":"","publisher":"","genre":""}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}
