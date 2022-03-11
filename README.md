# Go Library
## A library management API written in GO.
The goal of this project was to get familiar with writing APIs in Go and 
connecting them to a database. As I am a lover of literature, I themed the
API as a simple library inventory management tool.

## Requirements
Go 1.17

Mysql v1.5.0

Gorilla/mux v1.8.0

Gorm v1.9.16

Inflection v1.0.0

## Instructions
In order to run this API, you must have a database set up on your machine, and edit the
app.go file with the right connection configurations (your DB user, password, name, and port).

Then, from the main directory, perform the following command:

`go run cmd/main/main.go`

The server should now be up and running, awaiting any API calls. 

## Project Structure

Folder | Description
--- | ---
/cmd | Contains the main package which initializes the server.
/pkg | Contains all other packages utilized by main:
/config | Sets the database connection configurations.
/controllers | Defines the controllers functions used by the endpoints.
/models | Creates the Book model & database call functions.
/routes | Defines the API routes & assigns their functions.
/utils | Contains utility functions to avoid duplicate code.

## Resources
[Gorilla Mux](https://pkg.go.dev/github.com/gorilla/mux) - HTTP Routing/URL Matching

[GORM](https://gorm.io/index.html) - ORM (Object-relational mapping) library for Go

[MySql](https://www.mysql.com/) - For a free, open-source database

[MySQLWorkbench](https://www.mysql.com/products/workbench/) - For a free, open-source database visualizer