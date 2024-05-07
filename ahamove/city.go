package ahamove

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type CityInfo struct {
	ID        string  `json:"_id"`
	Name      string  `json:"name"`
	NameVI    string  `json:"name_vi_vn"`
	CountryID string  `json:"country_id"`
	Level     float64 `json:"level"`
}

func GetDataFromAPIListCity(apiUrl string, apiKey string, requestBody map[string]interface{}) ([]CityInfo, error) {
	// Replace with your actual API endpoint

	// Convert the request body to JSON
	var requestBodyJSON []byte
	var err error
	if requestBody != nil {
		requestBodyJSON, err = json.Marshal(requestBody)
		if err != nil {
			return nil, err
		}
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
	var responseObject []CityInfo
	err = json.Unmarshal(responseBody, &responseObject)
	if err != nil {
		return nil, err
	}

	return responseObject, nil
}
