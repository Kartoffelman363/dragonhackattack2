package service

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

/*
Service for sending custom api requests
*/
func ApiRequests(url string, body string, requestMethod string, contentType string) (*string, error) {
	// Creates new request
	request, err := http.NewRequest(requestMethod, url, bytes.NewBuffer([]byte(body)))
	if err != nil {
		return nil, fmt.Errorf("error creating request %v", err)
	}

	request.Header.Set("Content-Type", "application/json")

	// Send the request
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, fmt.Errorf("error sending request %v", err)
	}
	defer response.Body.Close()

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	strBody := string(responseBody)
	// Return the response body as a byte slice
	return &strBody, nil

}
