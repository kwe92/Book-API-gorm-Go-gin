package router

import (
	"book_api/controller"
	"book_api/database"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {

	router := gin.Default()

	router.POST("/book", controller.CreateBook(database.Database))

	router.POST("/books", controller.CreateBooks(database.Database))

	router.GET("/books", controller.GetBooks(database.Database))

	router.GET("/book/:id", controller.GetBook(database.Database))

	router.GET("/book/bytitle/:title", controller.GetBookByTitle(database.Database))

	router.POST("/books/:id", controller.UpdateBook(database.Database))

	router.DELETE("/books/:id", controller.DeleteBook(database.Database))

	return router
}
