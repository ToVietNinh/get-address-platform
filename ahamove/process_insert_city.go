package ahamove

import (
	"GetAddressGHN/helper"
	"fmt"
	"os"

	"github.com/joho/godotenv"

	"log"
)

const (
	getListCityEndpoint = "order/cities?country_id=VN"
)

func ProcessInsertCityInDatabase() error {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	responseGetListcities, err := GetDataFromAPIListCity(fmt.Sprintf("%s/%s", os.Getenv("AHAMOVE_STG_URL"), getListCityEndpoint), "", map[string]interface{}{})
	if err != nil {
		log.Fatalf("Error fetching data: %v", err)
		return err
	}
	db := helper.ConnectToMySQL()
	err = InsertDistrictToMappingTable(db, responseGetListcities)
	if err != nil {
		log.Fatalf("Error fetching data: %v", err)
		return err
	}
	return nil
}
