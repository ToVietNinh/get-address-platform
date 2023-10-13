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

	var queryCommandToInsertAddress []string

	queryItem := `INSERT INTO shipping_provider_district_mappings (shipping_provider_id, district_id, provider_district_key) VALUES`
	fmt.Println(queryItem)

	queryCommandToInsertAddress = append(queryCommandToInsertAddress, queryItem)

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

	type Districts struct {
		ID        int64     `json:"id" mapstructure:"id"`
		Code      string    `json:"code" mapstructure:"code"`
		Name      string    `json:"name" mapstructure:"name"`
		CityID    int64     `json:"city_id" mapstructure:"city_id"`
		CityName  string    `json:"city_name" mapstructure:"city_name"`
		CreatedAt time.Time `json:"created_at" mapstructure:"created_at"`
		UpdatedAt time.Time `json:"updated_at" mapstructure:"updated_at"`
	}

	type Wards struct {
		ID           int64     `json:"id" mapstructure:"id"`
		Name         string    `json:"name" mapstructure:"name"`
		Code         string    `json:"code" mapstructure:"code"`
		DistrictID   int64     `json:"district_id" mapstructure:"district_id"`
		DistrictName string    `json"district_name" mapstructure:"district_name"`
		CityName     string    `json:"city_name" mapstructure:"city_name"`
		CreatedAt    time.Time `json:"created_at" mapstructure:"created_at"`
		UpdatedAt    time.Time `json:"updated_at" mapstructure:"updated_at"`
	}
	var resultMyAddress []Districts
	query := db.Table("districts").
		Select("districts.*, cities.name as city_name").
		Joins("INNER JOIN cities ON districts.city_id = cities.id").
		Find(&resultMyAddress)
	if query.Error != nil {
		panic("Error db")
	}

	// Insert to data for create data for wardMapping
	// for _, item := range resultMyAddress {
	// 	for _, item1 := range listWardDataAll {
	// 		if item1.WardName == item.Name {
	// 			queryItem := fmt.Sprintf(`(7, %d, '%s'),`, item.ID, item1.WardCode)
	// 			queryCommandToInsertAddress = append(queryCommandToInsertAddress, queryItem)
	// 		}
	// 	}
	// }

	// Insert to data for create data for districtMapping
	for _, item := range resultMyAddress {
		for _, item1 := range listDistrictDataAll {
			if fmt.Sprintf("%s-%s", standardizeDistrictName(item1.DistrictName), standardizeProvinceName(cityIDToNameMapper[int64(item1.ProvinceID)])) == fmt.Sprintf("%s-%s", standardizeDistrictName(item.Name), standardizeProvinceName(item.CityName)) {
				queryItem := fmt.Sprintf(`(7, %d, '%s'),`, item.ID, item1.Code)
				queryCommandToInsertAddress = append(queryCommandToInsertAddress, queryItem)
			}
			fmt.Println(fmt.Sprintf("%s-%s", item1.DistrictName, cityIDToNameMapper[int64(item1.ProvinceID)]))
			fmt.Println(fmt.Sprintf("%s-%s", item.Name, item.CityName))
		}
	}

	fileName := "GHN_district_mapping.txt"

	// Open the file for writing. Create the file if it doesn't exist or truncate it if it does.
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error opening the file:", err)
		return
	}
	defer file.Close() // Ensure the file is closed when we're done

	for _, line := range queryCommandToInsertAddress {
		_, err := file.WriteString(line + "\n")
		if err != nil {
			fmt.Println("Error writing to the file:", err)
			return
		}
	}
	fmt.Println("Data has been written to the file.")
	fmt.Println(len(listDistrictDataAll))
}
