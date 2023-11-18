package model

import (
	"time"

	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title  string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
}

func (book *Book) Save(db *gorm.DB) (*Book, error) {

	if err := db.Create(book).Error; err != nil {

		return &Book{}, err
	}

	return book, nil
}

func (book *Book) Update(db *gorm.DB, input UpdateBookInput) (*Book, error) {

	if err := db.Model(book).Updates(input).Error; err != nil {

		return &Book{}, err

	}

	return book, nil

}

func (book *Book) Delete(db *gorm.DB) (Book, error) {

	deletedBook := *book

	if err := db.Delete(&book).Error; err != nil {

		return Book{}, err
	}

	deletedBook.DeletedAt = gorm.DeletedAt{
		Time: time.Now(),
	}

	return deletedBook, nil
}

func FindBookById(db *gorm.DB, id string) (Book, error) {

	var book Book

	if err := db.Where("id = ?", id).First(&book).Error; err != nil {

		return Book{}, err

	}

	return book, nil
}
