package externalapi

import (
	"bytes"
	"errors"
	"io"
	"net/http"
)

type APIClient struct {
	Client *http.Client
}

func (apiClient *APIClient) PostData(url string, data []byte) ([]byte, error) {

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		// fmt.Println("Failed to create POST request: ", err)
		return nil, errors.New("Failed to create POST request: " + err.Error())

	}
	// Set the content type header
	req.Header.Set("Content-Type", "application/json")

	resp, err := apiClient.Client.Do(req)
	if err != nil {
		// fmt.Println("Failed to send POST request: ", err)
		return nil, errors.New("Failed to send POST request: " + err.Error())
	}
	defer resp.Body.Close()
	return io.ReadAll(resp.Body)

}
