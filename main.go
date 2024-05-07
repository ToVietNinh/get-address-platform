package main

import (
	// "GetAddressGHN/ghn"
	"GetAddressGHN/ahamove"
	"log"
)

func main() {
	// ghn.ProcessInsertAddressInDatabase()
	err := ahamove.ProcessInsertCityInDatabase()
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
}
