package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func ProcessInsertAddressInDatabase() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	requestBodyForGetProvince := map[string]interface{}{
		"shop_id":       885,
		"from_district": 1447,
		"to_district":   1442,
	}

	// Create a request body (if needed)
	responseGetProvince, err := GetDataFromAPIProvince(fmt.Sprintf("%s/%s", os.Getenv("API_URL"), getProvinceEndPoint), os.Getenv("API_KEY"), requestBodyForGetProvince)
	if err != nil {
		log.Fatalf("Error fetching data: %v", err)
	}
	listProvinceData := responseGetProvince.Data

	var listDistrictDataAll []District
	for _, item := range listProvinceData {
		requestBodyForGetDistrict := map[string]interface{}{
			"province_id": item.ProvinceID,
		}
		responseGetDistrict, err := GetDataFromAPIDistrict(fmt.Sprintf("%s/%s", os.Getenv("API_URL"), getDistrictEndPoint), os.Getenv("API_KEY"), requestBodyForGetDistrict)
		if err != nil {
			log.Fatalf("Error fetching data: %v", err)
		}
		listDistrictDataOfProvinceItem := responseGetDistrict.Data
		listDistrictDataAll = append(listDistrictDataAll, listDistrictDataOfProvinceItem...)
	}
	// var listWardDataAll []Ward
	// for _, item := range listDistrictDataAll {
	// 	requestBodyForGetWard := map[string]interface{}{
	// 		"district_id": item.DistrictID,
	// 	}
	// 	responseGetWard, err := GetDataFromAPIWard(fmt.Sprintf("%s/%s", os.Getenv("API_URL"), getWardEndPoint), os.Getenv("API_KEY"), requestBodyForGetWard)
	// 	if err != nil {
	// 		log.Fatalf("Error fetching data: %v", err)
	// 	}
	// 	listWardDataForDistrictItem := responseGetWard.Data
	// 	listWardDataAll = append(listWardDataAll, listWardDataForDistrictItem...)
	// }

	// Write into file

	// Data formatter to insert
	var data []string
	for _, item := range listDistrictDataAll {
		data = append(data, fmt.Sprintf("%s-%s", item.Code, item.DistrictName))
	}

	fileName := "district.txt"

	// Open the file for writing. Create the file if it doesn't exist or truncate it if it does.
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error opening the file:", err)
		return
	}
	defer file.Close() // Ensure the file is closed when we're done

	for _, line := range data {
		_, err := file.WriteString(line + "\n")
		if err != nil {
			fmt.Println("Error writing to the file:", err)
			return
		}
	}
	fmt.Println("Data has been written to the file.")

	var queryCommandToInsertAddress []string

	queryItem := `(INSERT INTO shipping_provider_district_mappings (shipping_provider_id, district_id, provider_district_key))`
	fmt.Println(queryItem)

	queryCommandToInsertAddress = append(queryCommandToInsertAddress, queryItem)

}
