// Main creates a new mux router, registers the endpoint routes, and handles the requests on localhost.
package main

import (
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"go-library/pkg/models"
	"go-library/pkg/routes"
	"log"
	"net/http"
)

func main() {
	// Initialize database connection
	err := models.InitDB("root:password@tcp(127.0.0.1:3306)/go-library?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err)
	}

	// Create new router and register the routes
	r := mux.NewRouter()
	routes.RegisterLibraryRoutes(r)
	http.Handle("/", r)

	// Start server
	fmt.Printf("Starting server at port 9010...\n")
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}


