package models

import (
	"github.com/abisalde/go-bookstore/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `json:"name" gorm:"type:varchar(100);not null"`
	Author      string `json:"author" gorm:"type:varchar(100);not null"`
	Publication string `json:"publication" gorm:"type:varchar(100)"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	db.Create(&b)
	return b
}

func GetAllBooks() []Book {
	var books []Book
	db.Find(&books)

	return books
}

func GetBookById(Id int64) (*Book, *gorm.DB) {
	var book Book

	result := db.Where("ID = ?", Id).First(&book)
	return &book, result
}

func DeleteBook(Id int64) (Book, error) {
	var book Book
	result := db.Where("ID = ?", Id).Delete(&book)
	return book, result.Error
}

func UpdateBook(Id int64, updatedBook *Book) (*Book, error) {
	var book Book

	result := db.Where("ID = ?", Id).First(&book)

	if result.Error != nil {
		return nil, result.Error
	}

	result = db.Model(&book).Updates(updatedBook)
	if result.Error != nil {
		return nil, result.Error
	}

	return &book, nil
}
