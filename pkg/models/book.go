// Package models defines Book model and the database functions called by controller functions in package controllers.
package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

type Book struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	Author    string `json:"author"`
	Publisher string `json:"publisher"`
	Genre     string `json:"genre"`
}

// InitDB initializes a mysql database connection with the given dataSource.
func InitDB(dataSource string) error {
	d, err := gorm.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}
	db = d
	db.AutoMigrate(&Book{})
	return err
}

// GetDB returns the database initialized by InitDB.
func GetDB() *gorm.DB {
	return db
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
