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

		if savedBook, err := book.Save(db); err != nil {

			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return

		} else {

			ctx.JSON(http.StatusOK, gin.H{"data": savedBook})

		}

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

		if updatedBook, err := book.Update(db, updateBookInput); err != nil {

			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return

		} else {

			ctx.JSON(http.StatusOK, gin.H{"updated_book": updatedBook})

		}

	}
}

func DeleteBook(db *gorm.DB) gin.HandlerFunc {

	return func(ctx *gin.Context) {

		if book, err := model.FindBookById(db, ctx.Param("id")); err != nil {

			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return

		} else {

			if deletedBook, err := book.Delete(db); err != nil {

				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return

			} else {

				ctx.JSON(http.StatusOK, gin.H{"deleted_book": deletedBook})

			}

		}

	}

}
