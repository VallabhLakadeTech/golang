package externalapi

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExternalAPI(t *testing.T) {

	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"key": "value"})
	}))
	defer mockServer.Close()

	client := NewAPIClient()

	_, err := client.FetchData(mockServer.URL)
	assert.Nil(t, err)
}

func TestPostData(t *testing.T) {

	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Method != http.MethodPost {
			t.Errorf("Expected method POST, got %v", r.Method)
		}

		if r.Header.Get("Content-Type") != "application/json" {
			t.Errorf("Expected content type application/json, got %v", r.Header.Get("Content-Type"))
		}

		// Read the body data and check it
		body, err := io.ReadAll(r.Body)
		if err != nil {
			t.Fatalf("Failed to read request body: %v", err)
		}
		//Unmarshal it input struct for different URLs and return expectedBody and write response accordingly

		expectedBody := `{"key":"value"}`
		if string(body) != expectedBody {
			t.Errorf("Expected body %s, got %s", expectedBody, string(body))
		}

		// Respond with a mock response
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"response":"success"}`))

	}))
	defer mockServer.Close()

	client := &APIClient{
		client: mockServer.Client(),
	}
	data := []byte(`{"key":"value"}`)
	_, err := client.PostData(mockServer.URL, data)
	assert.Nil(t, err)

}
