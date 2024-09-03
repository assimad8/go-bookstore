package controllers

import (
	"encoding/json"
	"net/http"
	"time"

	// "fmt"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/assimad8/go-bookstore/pkg/models"
	"github.com/assimad8/go-bookstore/pkg/utils"
)

var NewBook models.Book

type Error struct {
	Data string `json:"error"`
}

func GetBooks(w http.ResponseWriter,r *http.Request) {
	newBooks := models.GetAllBooks()
	res, _ := json.Marshal(newBooks)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func GetBookById(w http.ResponseWriter,r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID,err := strconv.ParseInt(bookId,0,0)
	if err !=nil {
		err := Error{Data:"Invalid book ID"}
		res, _ := json.Marshal(err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(res)
		return 
	}
	bookDetails,_:= models.GetBookById(ID)
	if bookDetails.ID == 0 {
		err := Error{Data:"Book not found"}
		res, _ := json.Marshal(err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		w.Write(res)
		return 
	}
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func CreateBook(w http.ResponseWriter,r *http.Request) {
	book := new(models.Book) 
	utils.ParseBody(r,book)
	b := book.CreateBook()
	res, _ := json.Marshal(&b)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}
func DeleteBook(w http.ResponseWriter,r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID,err := strconv.ParseInt(bookId,0,0)
	if err !=nil {
		err := Error{Data:"Invalid book ID"}
		res, _ := json.Marshal(err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(res)
		return 
	}
	book:= models.DeleteBook(ID)
	if book.ID == 0 {
		err := Error{Data:"Book doesn't exist"}
		res, _ := json.Marshal(err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		w.Write(res)
		return 
	}
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write(res)
}
func UpdateBook(w http.ResponseWriter,r *http.Request) {
	bookUpdates := new(models.Book) 
	utils.ParseBody(r,bookUpdates)
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID,err := strconv.ParseInt(bookId,0,0)
	if err !=nil {
		err := Error{Data:"Invalid book ID"}
		res, _ := json.Marshal(err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(res)
		return 
	}
	bookDetails,db:= models.GetBookById(ID)
	if bookDetails.ID == 0 {
		err := Error{Data:"Book doesn't exist"}
		res, _ := json.Marshal(err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(res)
		return 
	}
	if bookUpdates.Author != "" {
		bookDetails.Author = bookUpdates.Author
	}
	if bookUpdates.Name != "" {
		bookDetails.Name = bookUpdates.Name
	}
	if bookUpdates.Publication != "" {
		bookDetails.Publication = bookUpdates.Publication
	}
	bookDetails.UpdatedAt = time.Now().UTC()
	db.Save(&bookDetails)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}