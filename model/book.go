package model

import (
	"fmt"
	"log"
	"time"

	"gorm.io/gorm"
)

// Books: Slice of pointers to Book type for batch insert into database
type Books []*Book

// Save: insert data from model into database
func (books *Books) Save(db *gorm.DB) (Books, error) {

	var result *gorm.DB

	// batch insert will invoke all implemented hooks for the model inserted
	// can skip hooks with &gorm.Session
	if result = db.Session(&gorm.Session{SkipHooks: true}).Create(books); result.Error != nil {
		return []*Book{}, result.Error
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

func (book *Book) Update(db *gorm.DB, input UpdateBookInput) (*Book, error) {

	originalBook := *book

	var result *gorm.DB

	if result = db.Model(book).Updates(input); result.Error != nil {

		return &Book{}, result.Error

	}

	log.Printf("updated book from: %+v to: %+v\n\n", originalBook, *book)

	log.Println("rows affected:", result.RowsAffected)

	return book, nil

}

func (book *Book) Delete(db *gorm.DB) (*Book, error) {

	deletedBook := book

	if err := db.Delete(&book).Error; err != nil {

		return &Book{}, err
	}

	deletedBook.DeletedAt.Time = time.Now()

	return deletedBook, nil
}

func FindBookById(db *gorm.DB, id string) (Book, error) {

	var book Book

	if err := db.Limit(1).Where("id = ?", id).Find(&book).Error; err != nil {

		return Book{}, err

	}

	return book, nil
}

// Locating Single Records

//   - use db.First or db.Limit(1).Find
//   - do not use db.Find without method chaining Limit(1)
//     for a single object as this will query the entire table for one record

// uniqueIndex

//   - a struct tag that creates and index contraints
//   - if fields have the same uniqueIndex value
//   - that section of the table must be unique
