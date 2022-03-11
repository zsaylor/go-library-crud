// Package controllers defines all the controller functions used by the routes in package routes.
package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"go-library/pkg/models"
	"go-library/pkg/utils"
	"net/http"
	"strconv"
)

// GetBooks gets all the books from the database and writes them back as a JSON.
func GetBooks(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetAllBooks()
	res, _ := json.Marshal(newBooks)
	WriteJSON(w, res)
}

// GetBookById receives an ID via request call and returns a JSON object of the book with the corresponding ID.
func GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID := ParseIdAsInt(bookId)
	bookDetails, _ := models.GetBookById(ID)
	res, _ := json.Marshal(bookDetails)
	WriteJSON(w, res)
}

// CreateBook creates a new book in the database and returns a JSON object of the newly added book.
func CreateBook(w http.ResponseWriter, r *http.Request) {
	newBook := &models.Book{}
	utils.ParseBody(r, newBook)
	b := newBook.CreateBook()
	res, _ := json.Marshal(b)
	WriteJSON(w, res)
}

// DeleteBook deletes a book with given ID and returns a JSON object of the deleted book.
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID := ParseIdAsInt(bookId)
	book := models.DeleteBook(ID)
	res, _ := json.Marshal(book)
	WriteJSON(w, res)
}

// UpdateBook finds book by ID and replaces any old data with the new data passed in the request.
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	var updateBook = &models.Book{}
	utils.ParseBody(r, updateBook)
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID := ParseIdAsInt(bookId)

	// Update the book fields only if new data in var updateBook is not null.
	bookDetails, db := models.GetBookById(ID)
	if updateBook.Name != "" {
		bookDetails.Name = updateBook.Name
	}
	if updateBook.Author != "" {
		bookDetails.Author = updateBook.Author
	}
	if updateBook.Publisher != "" {
		bookDetails.Publisher = updateBook.Publisher
	}
	if updateBook.Genre != "" {
		bookDetails.Genre = updateBook.Genre
	}
	db.Save(&bookDetails)

	res, _ := json.Marshal(bookDetails)
	WriteJSON(w, res)
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
