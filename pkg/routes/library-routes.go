// Package routes registers the endpoints and their respective functions.
package routes

import (
	"github.com/gorilla/mux"
	"go-library/pkg/controllers"
)

var RegisterLibraryRoutes = func(router *mux.Router) {
	router.HandleFunc("/api/v1/book/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/api/v1/book/", controllers.GetBooks).Methods("GET")
	router.HandleFunc("/api/v1/book/{bookId}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/api/v1/book/{bookId}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/api/v1/book/{bookId}", controllers.DeleteBook).Methods("DELETE")
}
