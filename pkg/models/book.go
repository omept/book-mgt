package models

import (
	"github.com/jinzhu/gorm"
	"github.com/ong-gtp/book-mgt/pkg/config"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() (*Book, *gorm.DB) {
	db.NewRecord(b)
	db := db.Create(&b)
	return b, db
}

func GetAllBooks() ([]Book, *gorm.DB) {
	var Books []Book
	db := db.Find(&Books)
	return Books, db
}

func GetBookById(id int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("ID=?", id).Find(&getBook)
	return &getBook, db
}

func DeleteBook(id int64) (Book, *gorm.DB) {
	var book Book
	db := db.Where("ID=?", id).Delete(book)
	return book, db
}

func (b *Book) UpdateBook() (*Book, *gorm.DB) {
	db := db.Model(b).Updates(b)
	return b, db
}
