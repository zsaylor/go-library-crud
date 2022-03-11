// Package utils parses a JSON received in the request call via ParseBody function.
package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// ParseBody takes the body from the HTTP request and unmarshals the JSON.
func ParseBody(r *http.Request, x interface{}) {
	if body, err := ioutil.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal([]byte(body), x); err != nil {
			return
		}
	}
}
