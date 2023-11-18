package database

import (
	"book_api/model"
	"fmt"
	"log"

	"gorm.io/driver/sqlite" // Sqlite driver based on CGO

	// "github.com/glebarez/sqlite" // Pure go SQLite driver, checkout https://github.com/glebarez/sqlite for details
	"gorm.io/gorm"
)

var Database *gorm.DB

func ConnectAndLoad() {
	connect()

	err := Database.AutoMigrate(&model.Book{})

	if err != nil {
		log.Fatal(err.Error())
	}
}

func connect() {

	var err error

	fmt.Print(Database)

	Database, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	fmt.Print(Database)

	if err != nil {
		log.Fatal(err.Error())
	}

}
