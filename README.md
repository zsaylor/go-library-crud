# Go Library
## 📚 A library management API written in GO. 📚
The goal of this project was to get familiar with writing APIs in Go and 
connecting them to a database. As I am a lover of literature, I themed the
API as a simple library inventory management tool.

This project utlizies the CRUD framework:
- **C**reate -> a new book entry
- **R**ead -> all of the books in the library (or just one)
- **U**pdate -> the data of a given book
- **D**elete -> a book from the library

## Requirements
#### See the `go.mod` file for the full list of required packages.
Go 1.17

Mysql v1.5.0

Gorilla/mux v1.8.0

Gorm v1.9.16

Inflection v1.0.0

## Instructions
In order to run this API, you must have a database connection already set up on your machine. Create a new database called 'go-library'. Then, edit the
main.go file with the right connection configurations (your DB user, password, name, and port). Check the [gorm docs](https://gorm.io/docs/connecting_to_the_database.html) for more details about connecting to a database.

Then, from the main directory, perform the following command:

`go run cmd/main/main.go`

The server should now be up and running, awaiting any new API calls. 

## Project Structure

Folder | Description
--- | ---
/cmd | Contains the main package which initializes the server.
/pkg | Contains all other packages utilized by main:
/controllers | Defines the controllers functions used by the endpoints.
/models | Creates the Book model & database call functions.
/routes | Defines the API routes & assigns their functions.
/utils | Contains utility functions to avoid duplicate code.

## Testing
I tested this project using Go's standard httptest packge. The tests are contained in the directory 'pkg/controllers/controllers_test.go'.

To run the tests, you will first need to create a new database named 'test-lib', and make sure the connection details in the TestMain function of the `controllers_test.go` file are correct.

Then run the command `go test ./...`.

## Resources
[Gorilla Mux](https://pkg.go.dev/github.com/gorilla/mux) - HTTP Routing/URL Matching

[GORM](https://gorm.io/index.html) - ORM (Object-relational mapping) library for Go

[httptest](https://pkg.go.dev/net/http/httptest) - Package httptest provides utilities for HTTP testing.

[MySql](https://www.mysql.com/) - For a free, open-source database

[MySQLWorkbench](https://www.mysql.com/products/workbench/) - For a free, open-source database visualizer

## Further Improvements
There are several improvements that could be made to this project:
- Abstract the connection details to make it easier to integrate with any database.
- Expand the Book model to include more book-related data.
- Create additional endpoints to allow for more queries by the user (query by Title, Genre, Author, etc).
