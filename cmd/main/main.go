package main

import (
	"log"
	"net/http"
	"com.shaily.tushar/go-bookStore/pkg/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)


func main(){
	r := mux.NewRouter();
	routes.RegisterRoutes(r);
	http.Handle("/",r);
	log.Fatal(http.ListenAndServe("localhost:9093",r));
}