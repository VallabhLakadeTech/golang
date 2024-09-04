package externalapi

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -coverprofile=coverage.out -timeout 99999s
// go tool cover -html=coverage.out
// go test -coverprofile=coverage.out -timeout 99999s -v -run TestPostData

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
		Client: mockServer.Client(),
	}
	data := []byte(`{"key":"value"}`)
	_, err := client.PostData(mockServer.URL, data)
	assert.Nil(t, err)

}

// Mock RoundTripper to simulate an error in http.Client.Do
type MockRoundTripper struct {
	RoundTripFunc func(req *http.Request) (*http.Response, error)
}

func (m *MockRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	return m.RoundTripFunc(req)
}

// func TestPostData_Errors(t *testing.T) {
// 	// Test case 1: Simulate error in http.NewRequest
// 	client := &APIClient{Client: &http.Client{}}
// 	_, err := client.PostData(":", nil) // Malformed URL to trigger http.NewRequest error
// 	if err == nil {
// 		t.Errorf("Expected error in http.NewRequest, got nil")
// 	}

// 	// Test case 2: Simulate error in http.Client.Do
// 	mockClient := &http.Client{
// 		Transport: &MockRoundTripper{
// 			RoundTripFunc: func(req *http.Request) (*http.Response, error) {
// 				return nil, errors.New("mocked error")
// 			},
// 		},
// 	}

// 	client = &APIClient{Client: mockClient}
// 	_, err = client.PostData("https://jsonplaceholder.typicode.com/posts", []byte(`{"key":"value"}`))
// 	if err == nil {
// 		t.Errorf("Expected error in http.Client.Do, got nil")
// 	}
// }
