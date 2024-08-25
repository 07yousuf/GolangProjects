package models

import (
	"github.com/07yousuf/BasicBookManagementSystem/PKG/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:"" json:"Name"`
	Author      string `json:"Author"`
	Publication string `json:"Publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}
func (b *Book) CreateBook() *Book {
	db.NewRecord(b)
	db.Create(&b)
	return b
}
func GetAllBooks() []Book {
	var Books []Book
	db.Find(Books)
	return Books
}
func GetByID(Id int64) (*Book, *gorm.DB) {
	var GetBook Book
	db := db.Where("Id=?", Id).Find(&GetBook)
	return &GetBook, db
}
func DeleteBook(Id int64) Book {
	var book Book
	db.Where("Id=?", Id).Delete(book)
	return book
}
