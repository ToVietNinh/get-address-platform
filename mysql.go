package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	ID   uint
	Name string
	Age  int
}

func ConnectToMySQL() *gorm.DB {
	// Open a connection to the MySQL database
	dsn := "root:SuperSecr3t@tcp(127.0.0.1:3306)/fulfillment?parseTime=true"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}
	defer db.DB()

	return db
}
