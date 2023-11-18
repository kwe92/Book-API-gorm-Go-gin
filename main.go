package main

import (
	"book_api/database"
	"book_api/router"
	"log"
)

func main() {

	database.ConnectAndLoad()

	router := router.SetupRouter()

	log.Fatal(router.Run(":8080"))

}
