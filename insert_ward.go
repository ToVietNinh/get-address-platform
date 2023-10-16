package main

import (
	"fmt"
	"os"
	"time"

	"gorm.io/gorm"
)

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

func insertWardToMappingTable(db *gorm.DB, listWardDataAll []Ward, listDistrictDataAll []District, listProvinceDataAll []Province) error {
	// Notice city and province is same

	// Mapper city id to name of it
	cityIDToNameMapper := make(map[int64]string)
	for _, province := range listProvinceDataAll {
		cityIDToNameMapper[int64(province.ProvinceID)] = province.ProvinceName
	}

	// Mapper district id to name of it
	districtIDToNameMapper := make(map[int64]string)
	districtIDToCityID := make(map[int64]int64)
	for _, district := range listDistrictDataAll {
		districtIDToNameMapper[int64(district.DistrictID)] = district.DistrictName
		districtIDToCityID[int64(district.DistrictID)] = int64(district.ProvinceID)
	}

	var resultMyAddress []Wards
	query := db.Table("wards").
		Select("wards.*, districts.name as district_name, cities.name as city_name").
		Joins("INNER JOIN districts ON districts.id = wards.district_id").
		Joins("INNER JOIN cities ON districts.city_id = cities.id").
		Find(&resultMyAddress)
	if query.Error != nil {
		panic("Error db")
	}

	var queryCommandToInsertAddress []string
	queryItem := `INSERT INTO shipping_provider_ward_mappings (shipping_provider_id, ward_id, provider_ward_key) VALUES`
	queryCommandToInsertAddress = append(queryCommandToInsertAddress, queryItem)

	for _, item := range resultMyAddress {
		for _, itemGHN := range listWardDataAll {
			fmt.Println(standardizeProvinceName(cityIDToNameMapper[districtIDToCityID[int64(itemGHN.DistrictID)]]))
			if fmt.Sprintf("%s-%s-%s", standardizeWardName(itemGHN.WardName), standardizeDistrictName(districtIDToNameMapper[int64(itemGHN.DistrictID)]), standardizeProvinceName(cityIDToNameMapper[districtIDToCityID[int64(itemGHN.DistrictID)]])) ==
				fmt.Sprintf("%s-%s-%s", standardizeWardName(item.Name), standardizeDistrictName(item.DistrictName), standardizeProvinceName(item.CityName)) {
				queryItem := fmt.Sprintf(`(7, %d, '%s'),`, item.ID, itemGHN.WardCode)
				queryCommandToInsertAddress = append(queryCommandToInsertAddress, queryItem)
			}
		}
	}

	fileName := "GHN_ward_mapping.txt"

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
