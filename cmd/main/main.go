// Main creates a new mux router, registers the endpoint routes, and handles the requests on localhost.
package main

import (
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"go-library/pkg/routes"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterLibraryRoutes(r)
	http.Handle("/", r)
	fmt.Printf("Starting server at port 9010...\n")
	log.Fatal(http.ListenAndServe("localhost:9010", r))

}
