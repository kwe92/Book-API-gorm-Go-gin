package controller

import (
	"book_api/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// CreateBook: http handler that inserts new book into database from request
func CreateBook(db *gorm.DB) gin.HandlerFunc {

	return func(ctx *gin.Context) {
		// expected request input struct
		var bookInput model.CreateBookInput

		// unmarshal request body buffer into expected input
		// if err := json.NewDecoder(ctx.Request.Body).Decode(&bookInput); err != nil {
		// 	ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		// 	return
		// }

		// unmarshal request body buffer into expected input
		if err := ctx.ShouldBindJSON(&bookInput); err != nil {

			// if de-serialization or validation fails
			// write to response body with status code and error
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

			return
		}

		// instantiate new object from de-serialized request input
		book := model.Book{
			Title:  bookInput.Title,
			Author: bookInput.Author,
		}

		// call Save method on object to insert record in database
		if savedBook, err := book.Save(db); err != nil {

			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return

		} else {
			// serialize object writting the object as JSON to response body buffer
			ctx.JSON(http.StatusOK, gin.H{"data": savedBook})

		}

	}
}

// CreateBooks: http handler that inserts new books into database from request
func CreateBooks(db *gorm.DB) gin.HandlerFunc {

	return func(ctx *gin.Context) {

		// expected request body input array
		var booksInput []model.CreateBookInput

		if err := ctx.ShouldBindJSON(&booksInput); err != nil {

			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// create a []*Object the length of []Object's in expected input
		books := make(model.Books, len(booksInput))

		// for each object instantiate new object from expected input
		// assign the new object to created array at index i
		for i, bookInput := range booksInput {
			books[i] = &model.Book{
				Title:  bookInput.Title,
				Author: bookInput.Author,
			}
		}

		// call Save method on object to insert all records into the database
		if savedBooks, err := books.Save(db); err != nil {

			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return

		} else {

			ctx.JSON(http.StatusOK, gin.H{"data": savedBooks})

		}

	}
}

// GetBooks: http handler that retreives all books from database
func GetBooks(db *gorm.DB) gin.HandlerFunc {

	return func(ctx *gin.Context) {

		var books []model.Book

		// select all records and load entire table into slice of model object
		if err := db.Find(&books).Error; err != nil {

			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return

		}

		ctx.JSON(http.StatusOK, gin.H{"data": books})
	}
}

// GetBooks: http handler that retreives a single book from database
func GetBook(db *gorm.DB) gin.HandlerFunc {

	return func(ctx *gin.Context) {

		var book model.Book

		// find the first record that matches the Where clause
		if err := db.Where("id = ?", ctx.Param("id")).First(&book).Error; err != nil {

			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return

		}

		ctx.JSON(http.StatusOK, gin.H{"data": book})
	}
}

// GetBooks: http handler that updates a single book in database
func UpdateBook(db *gorm.DB) gin.HandlerFunc {

	return func(ctx *gin.Context) {

		// expected request input struct
		var updateBookInput model.UpdateBookInput

		// declare model struct variable
		var book model.Book

		// find the record you want to update and load record into model struct variable
		if err := db.Where("id = ?", ctx.Param("id")).Find(&book).Error; err != nil {

			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return

		}

		// read `deserialize` request body buffer into expected input struct
		if err := ctx.ShouldBindJSON(&updateBookInput); err != nil {

			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

			return
		}

		// call Update method on object to update record in database
		if updatedBook, err := book.Update(db, updateBookInput); err != nil {

			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return

		} else {

			ctx.JSON(http.StatusOK, gin.H{"updated_book": updatedBook})

		}

	}
}

// GetBooks: http handler that deletes a single book in database
func DeleteBook(db *gorm.DB) gin.HandlerFunc {

	return func(ctx *gin.Context) {

		// find record you want to delete and load record into model struct variable
		if book, err := model.FindBookById(db, ctx.Param("id")); err != nil {

			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return

		} else {

			// call Delete method on object to delete record from database
			if deletedBook, err := book.Delete(db); err != nil {

				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return

			} else {

				ctx.JSON(http.StatusOK, gin.H{"deleted_book": deletedBook})

			}

		}

	}

}
