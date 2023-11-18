package controller

import (
	"book_api/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateBook(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var bookInput model.CreateBookInput

		if err := ctx.ShouldBindJSON(&bookInput); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		book := model.Book{
			Title:  bookInput.Title,
			Author: bookInput.Author,
		}

		db.Create(&book)

		ctx.JSON(http.StatusOK, gin.H{"data": book})

	}
}

func GetBooks(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var books []model.Book

		if err := db.Find(&books).Error; err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		ctx.IndentedJSON(http.StatusOK, gin.H{"data": books})
	}
}

func GetBook(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var book model.Book

		if err := db.Where("id = ?", ctx.Param("id")).First(&book).Error; err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		ctx.IndentedJSON(http.StatusOK, gin.H{"data": book})
	}
}

func UpdateBook(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var updateBookInput model.UpdateBookInput

		var book model.Book

		if err := db.Where("id = ?", ctx.Param("id")).Find(&book).Error; err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return

		}
		if err := ctx.ShouldBindJSON(&updateBookInput); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		db.Model(&book).Updates(updateBookInput)

		ctx.JSON(http.StatusOK, gin.H{"data": book})

	}
}

func DeleteBook(db *gorm.DB) gin.HandlerFunc {

	var book model.Book

	return func(ctx *gin.Context) {

		var err error

		if err = db.Where("id = ?", ctx.Param("id")).First(&book).Error; err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err = db.Delete(&book).Error; err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"deleted_book": book})

	}

}
