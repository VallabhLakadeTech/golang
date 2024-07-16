package test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/VallabhLakadeTech/golang/RESTAPI_mux/controller"
	"github.com/stretchr/testify/assert"
)

func mockServer() *httptest.Server {

	f := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		fmt.Fprintln(w, "")
	}
	return httptest.NewServer(http.HandlerFunc(f))
}

func TestGetAllBooks(t *testing.T) {
	url := "/library/books/"
	r := httptest.NewRequest("GET", url, nil)
	w := httptest.NewRecorder()
	http.DefaultServeMux.ServeHTTP(w, r)

	assert.Equals(t, w.Code, 200)

	var listOfBooks []controller.Books
	err := json.NewDecoder(w.Body).Decode(&listOfBooks)
	assert.Nil(t, err)

	assert.Equals(t, len(listOfBooks), 2)

}
