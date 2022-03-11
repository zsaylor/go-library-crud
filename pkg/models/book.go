// Package models defines Book model and the database functions called by controller functions in package controllers.
package models

import (
	"github.com/jinzhu/gorm"
	"go-library/pkg/config"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name      string `gorm: json:"name"`
	Author    string `json:"author"`
	Publisher string `json:"publisher"`
	Genre     string `json:"genre"`
}

// init initializes database connection.
func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

// CreateBook creates a new record book b and inserts it into the database.
func (b *Book) CreateBook() *Book {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

// GetAllBooks returns all books from the database as a slice.
func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

// GetBookById finds and returns a book with the given ID.
func GetBookById(Id int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("ID=?", Id).Find(&getBook)
	return &getBook, db
}

// DeleteBook finds and deletes a book with the given ID.
func DeleteBook(ID int64) Book {
	var book Book
	db.Where("ID=?", ID).Delete(book)
	return book
}
