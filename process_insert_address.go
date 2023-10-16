package main

import (
	"fmt"
	"log"
	"os"
	"time"

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
	// API get list cities of GHN
	responseGetProvince, err := GetDataFromAPIProvince(fmt.Sprintf("%s/%s", os.Getenv("API_URL"), getProvinceEndPoint), os.Getenv("API_KEY"), requestBodyForGetProvince)
	if err != nil {
		log.Fatalf("Error fetching data: %v", err)
	}
	listProvinceData := responseGetProvince.Data

	cityIDToNameMapper := make(map[int64]string)

	for _, province := range listProvinceData {
		cityIDToNameMapper[int64(province.ProvinceID)] = province.ProvinceName
	}

	// API get list district of GHN
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

	// API get list ward of GHN
	var listWardDataAll []Ward
	for _, item := range listDistrictDataAll {
		requestBodyForGetWard := map[string]interface{}{
			"district_id": item.DistrictID,
		}
		responseGetWard, err := GetDataFromAPIWard(fmt.Sprintf("%s/%s", os.Getenv("API_URL"), getWardEndPoint), os.Getenv("API_KEY"), requestBodyForGetWard)
		if err != nil {
			log.Fatalf("Error fetching data: %v", err)
		}
		listWardDataForDistrictItem := responseGetWard.Data
		listWardDataAll = append(listWardDataAll, listWardDataForDistrictItem...)
	}
	// Operate with database of fulfillment system
	db := ConnectToMySQL()

	type ShippingProviderDistrictMapping struct {
		ID                  int64     `json:"id" mapstructure:"id"`
		ShippingProviderID  int64     `json:"shipping_provider_id" mapstructure:"shipping_provider_id"`
		DistrictID          int64     `json:"district_id" mapstructure:"district_id"`
		ProviderDistrictKey string    `json:"provider_district_key" mapstructure:"provider_district_key"`
		CreatedAt           time.Time `json:"created_at" mapstructure:"created_at"`
		UpdatedAt           time.Time `json:"updated_at" mapstructure:"updated_at"`
	}

	// GEN INSERT DISTRICT COMMAND
	err = insertDistrictToMappingTable(db, listDistrictDataAll, listProvinceData)
	if err != nil {
		panic(err)
	}

	// GEN INSERT WARD COMMAND
	// err = insertWardToMappingTable(db, listWardDataAll, listDistrictDataAll, listProvinceData)
	// if err != nil {
	// 	panic(err)
	// }

}
