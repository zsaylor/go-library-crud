// Package utils contains several utility functions to avoid duplicate code.
package utils

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

// ParseBody takes the body from an HTTP request and unmarshals the JSON.
func ParseBody(r *http.Request, x interface{}) {
	if body, err := ioutil.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal([]byte(body), x); err != nil {
			return
		}
	}
}

// WriteJSON sets the http status and writes/responds with the given JSON object.
func WriteJSON(w http.ResponseWriter, res []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// ParseIdAsInt parses a given ID string as an integer
func ParseIdAsInt(bookId string) int64 {
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	return ID
}

// CheckResponseCode returns an error if the given expected response does not match the actual response.
func CheckResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("handler returned wrong status code: got %v want %v", expected, actual)
	}
}

// NewRequest returns an HTTP request with added path variables.
func NewRequest(method, path string, vars map[string]string) *http.Request {
	r := httptest.NewRequest(method, path, nil)
	return mux.SetURLVars(r, vars)
}
