package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetDataFromAPIProvince(apiUrl string, apiKey string, requestBody map[string]interface{}) (*ResponseForProvince, error) {
	// Replace with your actual API endpoint

	// Convert the request body to JSON
	requestBodyJSON, err := json.Marshal(requestBody)
	if err != nil {
		return nil, err
	}

	// Create a new HTTP client
	client := &http.Client{}

	// Create a request
	req, err := http.NewRequest("GET", apiUrl, bytes.NewBuffer(requestBodyJSON))
	if err != nil {
		return nil, err
	}

	// Add authorization header
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", apiKey))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Token", apiKey)

	// Send the request
	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	// Read the response body
	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	// Parse the response
	var responseObject ResponseForProvince
	err = json.Unmarshal(responseBody, &responseObject)
	if err != nil {
		return nil, err
	}

	return &responseObject, nil
}

func GetDataFromAPIDistrict(apiUrl string, apiKey string, requestBody map[string]interface{}) (*ResponseForDistrict, error) {
	// Replace with your actual API endpoint

	// Convert the request body to JSON
	requestBodyJSON, err := json.Marshal(requestBody)
	if err != nil {
		return nil, err
	}

	// Create a new HTTP client
	client := &http.Client{}

	// Create a request
	req, err := http.NewRequest("GET", apiUrl, bytes.NewBuffer(requestBodyJSON))
	if err != nil {
		return nil, err
	}

	// Add authorization header
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", apiKey))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Token", apiKey)

	// Send the request
	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	// Read the response body
	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	// Parse the response
	var responseObject ResponseForDistrict
	err = json.Unmarshal(responseBody, &responseObject)
	if err != nil {
		return nil, err
	}

	return &responseObject, nil
}

func GetDataFromAPIWard(apiUrl string, apiKey string, requestBody map[string]interface{}) (*ResponseForWard, error) {
	// Replace with your actual API endpoint

	// Convert the request body to JSON
	requestBodyJSON, err := json.Marshal(requestBody)
	if err != nil {
		return nil, err
	}

	// Create a new HTTP client
	client := &http.Client{}

	// Create a request
	req, err := http.NewRequest("GET", apiUrl, bytes.NewBuffer(requestBodyJSON))
	if err != nil {
		return nil, err
	}

	// Add authorization header
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", apiKey))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Token", apiKey)

	// Send the request
	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	// Read the response body
	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	// Parse the response
	var responseObject ResponseForWard
	err = json.Unmarshal(responseBody, &responseObject)
	if err != nil {
		return nil, err
	}

	return &responseObject, nil
}
