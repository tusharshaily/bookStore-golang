package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"com.shaily.tushar/go-bookStore/pkg/models"
	"com.shaily.tushar/go-bookStore/pkg/utils"
	"github.com/gorilla/mux"
)

var NewBook models.Book;


func DeleteBook(w http.ResponseWriter, r* http.Request){
	vars := mux.Vars(r);
	bookId := vars["bookId"];
	ID , err := strconv.ParseInt(bookId,0,0);
	if err!=nil{
		fmt.Println("The book Id is not valid Int")
	}
	book := models.DeleteBook(ID)

	res , _ := json.Marshal(book)

	w.Header().Set("Content-Type","application/json");
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func CreateBook(w http.ResponseWriter,r* http.Request){
	createBook := &models.Book{};
	utils.ParseBody(r,createBook);
	b := createBook.CreateBook();
	res,_ := json.Marshal(b);
	w.Header().Set("Content-Type","application/json");
	w.WriteHeader(http.StatusOK);
	w.Write(res);
}

func GetAllBooks(w http.ResponseWriter, r* http.Request){
	allbooks := models.GetAllBooks();
	res ,_ := json.Marshal(allbooks);
	w.Header().Set("Content-Type","application/json");
	w.WriteHeader(http.StatusOK);
	w.Write(res);
}

func GetBookById(w http.ResponseWriter, r* http.Request){
	 vars := mux.Vars(r)
	bookId := vars["bookId"];
	ID , err := strconv.ParseInt(bookId,0,0);
	if err!=nil{
		fmt.Println("error while parsing the bookID")

	}
	bookDetails , _ := models.GetBookById(ID);
	res,_ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type","application/json");
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r* http.Request){
	updateBook := &models.Book{};
	utils.ParseBody(r,updateBook);
	
	vars := mux.Vars(r);
	bookId := vars["bookId"];
	Id , err := strconv.ParseInt(bookId,0,0);
	
	if err!=nil{
		fmt.Println("The bookId is not parsable to Int");
	}

	savedBook,db := models.GetBookById(Id);

	if updateBook.Name != ""{
		savedBook.Name = updateBook.Name;
	}

	if updateBook.Publisher != ""{
		savedBook.Publisher = updateBook.Publisher;
	}
	db.Save(&savedBook);

	res , _ := json.Marshal(savedBook);

	w.Header().Set("Content-type","application/json");
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

