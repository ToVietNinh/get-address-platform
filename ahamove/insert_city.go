package ahamove

import (
	"GetAddressGHN/helper"
	"fmt"
	"os"
	"time"

	"gorm.io/gorm"
)

type City struct {
	ID        int64     `json:"id" mapstructure:"id"`
	Name      string    `json:"name" mapstructure:"name"`
	Code      string    `json:"code" mapstructure:"code"`
	CreatedAt time.Time `json:"created_at" mapstructure:"created_at"`
	UpdatedAt time.Time `json:"updated_at" mapstructure:"updated_at"`
}

func InsertDistrictToMappingTable(db *gorm.DB, responseGetListcities []CityInfo) error {
	var resultMyCities []City
	err := db.Table("cities").
		Select("cities.*").
		Find(&resultMyCities).Error
	if err != nil {
		return err
	}
	// fmt.Println(resultMyCities)

	var queryCommand []string
	queryItem := `INSERT INTO shipping_provider_city_mappings (shipping_provider_id, city_id, provider_city_key) VALUES`
	queryCommand = append(queryCommand, queryItem)
	// Mapping city between 2 systems
	for _, myItem := range resultMyCities {
		for _, ahaItem := range responseGetListcities {
			if helper.StandardizeProvinceName(myItem.Name) == helper.StandardizeProvinceName(ahaItem.Name) ||
				helper.StandardizeProvinceName(myItem.Name) == helper.StandardizeProvinceName(ahaItem.NameVI) {
				queryItem := fmt.Sprintf(`(2, %d, '%s'),`, myItem.ID, ahaItem.ID)
				queryCommand = append(queryCommand, queryItem)
			}
		}
	}

	// Hue + Ba Ria - Vung Tau
	fmt.Println(queryCommand)
	fmt.Printf("Number record inserted in DB: %d \n", len(queryCommand)-1)
	fmt.Printf("Number record of Ahamove: %d \n", len(responseGetListcities))

	fileName := "ahamove/AHAMOVE_stg_mapping_city"

	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error opening the file:", err)
		return err
	}
	defer file.Close() // Ensure the file is closed when we're done

	for _, line := range queryCommand {
		_, err := file.WriteString(line + "\n")
		if err != nil {
			fmt.Println("Error writing to the file:", err)
			return err
		}
	}
	fmt.Println("Data has been written to the file.")

	return nil

}
