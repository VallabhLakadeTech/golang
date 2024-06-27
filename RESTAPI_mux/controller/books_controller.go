package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type Books struct {
	ID              string `json:"id"`
	BookName        string `json:"bookname"`
	Author          string `json:"author"`
	PublicationYear int    `json:"publicationyear"`
}

var mapOfBooks map[string]Books

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Welcome to the library"))
}

func AddBooks(w http.ResponseWriter, r *http.Request) {
	book := Books{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&book)
	if err != nil {
		response := Response{
			StatusCode: http.StatusInternalServerError,
			Msg:        "error unmarshalling request",
			Err:        err,
			Body:       nil,
		}
		writeResponse(w, response, http.StatusInternalServerError)
		return
	}
	if len(mapOfBooks) == 0 {
		mapOfBooks = map[string]Books{}
	}
	book.ID = uuid.New().String()
	mapOfBooks[book.ID] = book
	writeResponse(w, book, http.StatusOK)
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	pathVariables := mux.Vars(r)
	id := pathVariables["id"]
	book, ok := mapOfBooks[id]
	if !ok {
		writeResponse(w, "Book not found", http.StatusNotFound)
		return
	}
	writeResponse(w, book, http.StatusOK)
}

func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	if len(mapOfBooks) > 0 {
		var listOfBooks []Books
		for _, book := range mapOfBooks {
			listOfBooks = append(listOfBooks, book)
		}
		writeResponse(w, listOfBooks, http.StatusOK)
		return
	}
	writeResponse(w, "No book found", http.StatusNotFound)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var book Books
	err := decoder.Decode(&book)
	if err != nil {
		writeResponse(w, "error decoding the provided input", http.StatusInternalServerError)
		return
	}
	_, ok := mapOfBooks[book.ID]
	if !ok {
		writeResponse(w, "book not found", http.StatusNotFound)
		return
	}
	mapOfBooks[book.ID] = book
	writeResponse(w, book, http.StatusOK)
}

func RemoveBook(w http.ResponseWriter, r *http.Request) {
	pathVariables := mux.Vars(r)
	bookToRemove := pathVariables["id"]
	if _, ok := mapOfBooks[bookToRemove]; !ok {
		writeResponse(w, "book not found", http.StatusNotFound)
		return
	}
	delete(mapOfBooks, bookToRemove)
	writeResponse(w, "book removed", http.StatusNotFound)
}

func GetByAuthor(w http.ResponseWriter, r *http.Request) {
	author := r.URL.Query().Get("author")
	fmt.Println("author: ", author)
	var authorsBook []Books
	for _, book := range mapOfBooks {
		if book.Author == author {
			authorsBook = append(authorsBook, book)
		}
	}
	if len(authorsBook) == 0 {
		writeResponse(w, "No book found of the provider author", http.StatusNotFound)
		return
	}
	writeResponse(w, authorsBook, http.StatusOK)
}

type Response struct {
	Msg        string
	StatusCode int
	Err        error
	Body       interface{}
}

func writeResponse(w http.ResponseWriter, response interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	enc := json.NewEncoder(w)
	err := enc.Encode(response)
	if err != nil {
		fmt.Println("error encoding the response")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(statusCode)
}
