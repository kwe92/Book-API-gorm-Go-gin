package model

import (
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Books: Slice of pointers to Book type for batch insert into database
type Books []Book

// Save: insert data from model into database
func (books *Books) Save(db *gorm.DB) (Books, error) {

	var result *gorm.DB

	// batch insert records and invoke hooks
	// skip hook invocation with &gorm.Session
	if result = db.Session(&gorm.Session{SkipHooks: true}).Create(books); result.Error != nil {
		return []Book{}, result.Error
	}

	return *books, nil
}

type Book struct {
	gorm.Model

	Title  string `gorm:"uniqueIndex:compositeindex;not null" json:"title" binding:"required"`
	Author string `gorm:"uniqueIndex:compositeindex;not null" json:"author" binding:"required"`
}

// Save: insert data from model into database
func (book *Book) Save(db *gorm.DB) (*Book, error) {

	var result *gorm.DB

	// note: *gorm.DB.Create must be passed a pointer
	if result = db.Create(book); result.Error != nil {

		return &Book{}, result.Error
	}

	return book, nil
}

// BeforeSave: gorm hook called before a model is inserted into the database
func (book *Book) BeforeSave(*gorm.DB) error {

	fmt.Printf("\n\nBEFORE SAVING BOOK: %+v\n\n", book)

	return nil

}

// Update: updates data for the associated book record in database
func (book *Book) Update(db *gorm.DB, input UpdateBookInput) (*Book, error) {

	originalBook := *book

	var result *gorm.DB

	// specify model you want to perfom operations on and update record
	if result = db.Model(book).Updates(input); result.Error != nil {

		return &Book{}, result.Error

	}

	log.Printf(
		"updated book from: %+v to: %+v\n\n",
		gin.H{
			"title":  originalBook.Title,
			"author": originalBook.Author,
		},
		gin.H{
			"title":  book.Title,
			"author": book.Author,
		},
	)

	log.Println("rows affected:", result.RowsAffected)

	return book, nil

}

// Delete: remove book record from database
func (book *Book) Delete(db *gorm.DB) (*Book, error) {

	if err := db.Delete(&book).Error; err != nil {

		return &Book{}, err
	}

	return book, nil
}

// FindBookById: query database for book record with given id
func FindBookById(db *gorm.DB, id string) (Book, error) {

	// destination struct pointer
	var book Book

	// query with format string in where clause
	if err := db.Limit(1).Where("id = ?", id).Find(&book).Error; err != nil {

		return Book{}, err

	}
	// if destinaion struct has zero-value for id return not found to client
	if book.ID == 0 {
		return Book{}, errors.New(fmt.Sprintf("could not find a book with the id: %s", id))
	}

	return book, nil
}

// FindBookByTitle: query database for book record with given title
func FindBookByTitle(db *gorm.DB, title string) (Book, error) {

	filteredTitle := strings.Replace(title, "-", " ", -1)
	filteredTitle = strings.Replace(filteredTitle, "_", " ", -1)

	// destination struct pointer
	var book Book

	// query with struct in where clause
	if err := db.Limit(1).Where(&Book{Title: filteredTitle}).Find(&book).Error; err != nil {

		return Book{}, err

	}

	// if destinaion struct has zero-value for title return not found to client
	if book.Title == "" {
		return Book{}, errors.New(fmt.Sprintf("could not find a book with the title: %s", filteredTitle))
	}

	return book, nil
}

// Batch Insert

//   - passing a Slice to Create method inserts multiple records of a model
//   - skip hook invocation with Session(&gorm.Session{SkipHooks: true})

// Updating Records

//   - Update method: updates a single field / column
//   - Updates method: updates multiple fields / columns

// Deleting Records

//   - Delete method: deletes record and assigns value to DeletedAt in destination struct

// Selecting Specific Fields (Columns)

//   - *gorm.Select(coulmn_1, coulmn_2, ..., coulmn_n)
//   -  if you omit the Select method when querying
//      gorm selects all fields / columns by default e.g. SELECT * FROM table_name;

// Destination Types

//   - struct: single record
//   - slice: multiple records, invokes hooks unless explicitly turned off
//   - map[string] interface{}: single record without invoking hooks

// Query Types

//   - raw query string
//   - formatted query string
//   - struct query //! note: zero values are ommited from the struct query string unless you specify struct search fields
//   - map[string] interface{} query

// Query Single Records

//   - use db.First or db.Limit(1).Find
//   - do not use db.Find without method chaining Limit(1)
//     for a single object as this will query the entire table for one record

// uniqueIndex

//   - struct tag that creates index contraints
//   - if fields have the same uniqueIndex value
//   - the composition of those fields must be unique
