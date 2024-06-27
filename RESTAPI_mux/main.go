package main

import (
	"fmt"
	"net/http"

	"github.com/VallabhLakadeTech/golang/RESTAPI_mux/controller"
	"github.com/gorilla/mux"
)

/*
Problem statement:
Develop a RESTful API for a Book Management System using Go. This system will allow users to perform CRUD (Create, Read, Update, Delete) operations on a collection of books. Each book will have a title, author, ISBN, publication year, and a status indicating whether it is available or checked out.

Requirements
API Endpoints:

Create a new book: POST /books
Retrieve a list of all books: GET /books
Retrieve details of a specific book by its ID: GET /books/{id}
Update details of a specific book by its ID: PUT /books/{id}
Delete a specific book by its ID: DELETE /books/{id}
Check out a book: PATCH /books/{id}/checkout
Return a book: PATCH /books/{id}/return
Data Model:
Each book should have the following fields:

ID (string, generated automatically)
Title (string, required)
Author (string, required)
ISBN (string, required, unique)
PublicationYear (int, required)
Status (string, can be "available" or "checked out")
Functional Requirements:

Users should be able to add new books to the collection.
Users should be able to view a list of all books in the collection.
Users should be able to view details of a specific book.
Users should be able to update information of an existing book.
Users should be able to delete a book from the collection.
Users should be able to check out a book (change status to "checked out").
Users should be able to return a book (change status to "available").
Non-Functional Requirements:

The API should return appropriate HTTP status codes for different operations (e.g., 200 OK, 201 Created, 400 Bad Request, 404 Not Found, 500 Internal Server Error).
Implement basic validation for the input data.
Use a simple in-memory store for data persistence (for learning purposes, a database integration can be added later).
Implementation Details
Router:

Use a router like gorilla/mux or gin-gonic/gin to handle routing of API endpoints.
Handlers:

Implement handler functions for each of the endpoints that will interact with the data model.
Data Storage:

Use a map or a slice as an in-memory data store for managing books.
Middleware:

(Optional) Implement middleware for logging requests and validating input data.
Error Handling:

Ensure proper error handling and return meaningful error messages to the client.
*/
func main() {

	router := mux.NewRouter()
	router.Use(middleware)

	router.HandleFunc("/", controller.HomeHandler).Methods("GET")

	// Subrouter for books
	booksSubrouter := router.PathPrefix("/library/books").Subrouter()
	booksSubrouter.HandleFunc("/", controller.AddBooks).Methods("POST")
	booksSubrouter.HandleFunc("/", controller.GetAllBooks).Methods("GET")
	booksSubrouter.HandleFunc("/{id}", controller.GetBook).Methods("GET")
	booksSubrouter.HandleFunc("/{id}", controller.UpdateBook).Methods("PUT")
	booksSubrouter.HandleFunc("/{id}", controller.RemoveBook).Methods("DELETE")
	booksSubrouter.HandleFunc("/author", controller.GetByAuthor).Methods("GET")

	http.ListenAndServe(":8085", router)
}

func middleware(httpHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Method: %v, URI: %v\n", r.Method, r.URL)
		httpHandler.ServeHTTP(w, r)
	})
}
