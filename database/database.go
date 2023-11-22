package database

import (
	"book_api/model"
	"log"

	"gorm.io/driver/sqlite" // Sqlite driver based on CGO

	"gorm.io/gorm"
)

// global database variable used to communicate with currently connected database
var Database *gorm.DB

// ConnectAndMigrate: connects to database and migrates schema
// creating and altering tables for models if the table does not exist within the database
// or if model fields / struct tags are added
func ConnectAndMigrate() {

	// connect to database
	connect()

	// drop existing table on server start
	if err := Database.Exec("drop table books;").Error; err != nil {
		// TODO: gracefully shutdown the database
		log.Fatal(err.Error())
	}

	// create and alter table in database from model
	err := Database.AutoMigrate(&model.Book{})

	if err != nil {
		// TODO: gracefully shutdown the database
		log.Fatal(err.Error())
	}
}

// connect: connect to database
func connect() {

	var err error

	// connect to database and initialize global database variable
	Database, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	// TODO: gracefully shutdown the database
	if err != nil {
		log.Fatal(err.Error())
	}

}
