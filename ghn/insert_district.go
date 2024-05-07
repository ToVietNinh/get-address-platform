package ghn

import (
	"fmt"
	"os"
	"time"

	"GetAddressGHN/helper"

	"gorm.io/gorm"
)

type Districts struct {
	ID        int64     `json:"id" mapstructure:"id"`
	Code      string    `json:"code" mapstructure:"code"`
	Name      string    `json:"name" mapstructure:"name"`
	CityID    int64     `json:"city_id" mapstructure:"city_id"`
	CityName  string    `json:"city_name" mapstructure:"city_name"`
	CreatedAt time.Time `json:"created_at" mapstructure:"created_at"`
	UpdatedAt time.Time `json:"updated_at" mapstructure:"updated_at"`
}

func insertDistrictToMappingTable(db *gorm.DB, listDistrictDataAll []District, listProvinceDataAll []Province) error {

	cityIDToNameMapper := make(map[int64]string)

	for _, province := range listProvinceDataAll {
		cityIDToNameMapper[int64(province.ProvinceID)] = province.ProvinceName
	}
	var resultMyAddress []Districts
	query := db.Table("districts").
		Select("districts.*, cities.name as city_name").
		Joins("INNER JOIN cities ON districts.city_id = cities.id").
		Find(&resultMyAddress)
	if query.Error != nil {
		panic("Error db")
	}

	var queryCommandToInsertAddress []string
	queryItem := `INSERT INTO shipping_provider_district_mappings (shipping_provider_id, shipping_provider_code, district_id, provider_district_key) VALUES`
	queryCommandToInsertAddress = append(queryCommandToInsertAddress, queryItem)

	for _, item := range resultMyAddress {
		for _, item1 := range listDistrictDataAll {
			if fmt.Sprintf("%s-%s", helper.StandardizeDistrictName(item1.DistrictName), helper.StandardizeProvinceName(cityIDToNameMapper[int64(item1.ProvinceID)])) == fmt.Sprintf("%s-%s", helper.StandardizeDistrictName(item.Name), helper.StandardizeProvinceName(item.CityName)) {
				queryItem := fmt.Sprintf(`(7, 'GHN', %d, '%d'),`, item.ID, item1.DistrictID)
				queryCommandToInsertAddress = append(queryCommandToInsertAddress, queryItem)
				fmt.Println(queryCommandToInsertAddress)
			}
		}
	}
	fileName := "GHN_live_mapping_district"

	// Open the file for writing. Create the file if it doesn't exist or truncate it if it does.
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error opening the file:", err)
		return err
	}
	defer file.Close() // Ensure the file is closed when we're done

	for _, line := range queryCommandToInsertAddress {
		_, err := file.WriteString(line + "\n")
		if err != nil {
			fmt.Println("Error writing to the file:", err)
			return err
		}
	}
	fmt.Println("Data has been written to the file.")

	return nil
}
