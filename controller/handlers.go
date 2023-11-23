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

			// if de-serialization or validation fail
			// return error to client
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

			return
		}

		// instantiate new object from request input
		book := model.Book{
			Title:  bookInput.Title,
			Author: bookInput.Author,
		}

		// insert record in database
		if savedBook, err := book.Save(db); err != nil {

			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return

		} else {
			// serialize object writting object as JSON to response body buffer
			ctx.JSON(http.StatusOK, gin.H{"data": savedBook})

		}

	}
}

// CreateBooks: http handler that inserts new books into database from request
func CreateBooks(db *gorm.DB) gin.HandlerFunc {

	return func(ctx *gin.Context) {

		// expected request input
		var booksInput []model.CreateBookInput

		if err := ctx.ShouldBindJSON(&booksInput); err != nil {

			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// create Slice the length of expected request input
		books := make(model.Books, len(booksInput))

		// for each object instantiate new object from expected input
		// assign the new object to created Slice at index i
		for i, bookInput := range booksInput {
			books[i] = model.Book{
				Title:  bookInput.Title,
				Author: bookInput.Author,
			}
		}

		// insert all records into database
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

		// declare destination slice
		var books model.Books

		// select all records and load into destination slice
		if err := db.Order("author asc, title asc").Find(&books).Error; err != nil {

			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return

		}

		ctx.IndentedJSON(http.StatusOK, gin.H{"data": books})
	}
}

// GetBooks: http handler that retreives a single book from database by id
func GetBook(db *gorm.DB) gin.HandlerFunc {

	return func(ctx *gin.Context) {

		// declare destination struct
		var book model.Book

		var err error

		// find first record matching Where clause and load into destination struct
		if book, err = model.FindBookById(db, ctx.Param("id")); err != nil {

			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return

		}

		ctx.JSON(http.StatusOK, gin.H{"data": book})
	}
}

// / GetBooks: http handler that retreives a single book from database by title
func GetBookByTitle(db *gorm.DB) gin.HandlerFunc {

	return func(ctx *gin.Context) {

		// declare destination struct
		var book model.Book

		var err error

		// find first record matching Where clause and load into destination struct
		if book, err = model.FindBookByTitle(db, ctx.Param("title")); err != nil {

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

		// declare destination struct
		var book model.Book

		var err error

		// find record to update and load record into destination struct
		if book, err = model.FindBookById(db, ctx.Param("id")); err != nil {

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

		// find record to delete and load record into destination struct
		if book, err := model.FindBookById(db, ctx.Param("id")); err != nil {

			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return

		} else {

			// call Delete method on object to delete record from database
			if deletedBook, err := book.Delete(db); err != nil {

				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return

			} else {
				// return deleted object with DeletedAt time
				ctx.JSON(http.StatusOK, gin.H{"deleted_book": deletedBook})

			}

		}

	}

}
