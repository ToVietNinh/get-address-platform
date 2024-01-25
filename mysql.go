package main

import (
	"os"

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
	// dsn := "admin:SuperSecr3t@tcp(127.0.0.1:3307)/fulfillment?parseTime=true"
	db, err := gorm.Open(mysql.Open(os.Getenv("LIVE_MYSQL_URL")), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}
	defer db.DB()

	return db
}
