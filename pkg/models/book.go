package models

import (
	"com.shaily.tushar/go-bookStore/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB;

type Book struct{
	gorm.Model
	Name string `gorm:""json:"name"` 
	Publisher string `json:"publisher"`
};

func init(){
	config.Connect();
	db = config.GetDb();
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book{
	db.NewRecord(b);
	db.Create(&b);
	return b;
}

func GetAllBooks() []Book{
	var books []Book
	db.Find(&books)
	return books
}

func GetBookById(ID int64) (*Book,*gorm.DB){
	var getBook Book
	db := db.Where("ID=?",ID).Find(&getBook)
	return &getBook,db;
}

func DeleteBook(ID int64) Book{
	var book Book;
	db.Where("ID=?",ID).Find(&book);
	return book;
}