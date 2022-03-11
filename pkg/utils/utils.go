// Package utils contains several utility functions to avoid duplicate code.
package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

// ParseBody takes the body from an HTTP request and unmarshals the JSON.
func ParseBody(r *http.Request, x interface{}) {
	if body, err := ioutil.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal([]byte(body), x); err != nil {
			return
		}
	}
}

// WriteJSON sets the http status and writes/responds with the given JSON object
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
