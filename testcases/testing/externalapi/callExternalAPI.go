package externalapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type APIClient struct {
	client *http.Client
}

func NewAPIClient() *APIClient {

	return &APIClient{
		client: &http.Client{},
	}

}

func (apiClient *APIClient) FetchData(url string) (map[string]interface{}, error) {

	resp, err := apiClient.client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var data map[string]interface{}
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, err
	}
	return data, nil
}

func (apiClient *APIClient) PostData(url string, data []byte) ([]byte, error) {

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		log.Fatalf("Failed to create POST request: %v", err)
	}
	// Set the content type header
	req.Header.Set("Content-Type", "application/json")

	resp, err := apiClient.client.Do(req)
	if err != nil {
		log.Fatalf("Failed to send POST request: %v", err)
	}
	defer resp.Body.Close()
	fmt.Println("Printing this")
	return io.ReadAll(resp.Body)

}
