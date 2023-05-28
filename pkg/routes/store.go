package routes

import (
	"com.shaily.tushar/go-bookStore/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterRoutes = func(router * mux.Router){
	router.HandleFunc("/books",controller.CreateBook).Methods("POST");
	router.HandleFunc("/books",controller.GetAllBooks).Methods("GET");
	router.HandleFunc("/books/{bookId}",controller.GetBookById).Methods("GET");
	router.HandleFunc("/books/{bookId}",controller.UpdateBook).Methods("PUT");
	router.HandleFunc("/books/{bookId}",controller.DeleteBook).Methods("DELETE");
}