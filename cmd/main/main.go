package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	_ "github.com/go-sql-driver/mysql"

	"github.com/assimad8/go-bookstore/pkg/routes"

)


func main(){
	router := mux.NewRouter()
	routes.RegisterBookStoreRoutes(router)
	
	fmt.Println("Stating the Server on PORT::8080")
	http.Handle("/", router)
	if err := http.ListenAndServe(":8080",router);err !=nil{
		log.Fatal(err)
	}

}
