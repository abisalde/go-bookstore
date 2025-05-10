package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/abisalde/go-bookstore/pkg/models"
	"github.com/abisalde/go-bookstore/pkg/utils"
	"github.com/gorilla/mux"
)

var NewBook models.Book

func CreateBook(w http.ResponseWriter, r *http.Request) {

	createBook := &models.Book{}
	utils.ParseBody(r, createBook)
	book := createBook.CreateBook()

	w.Header().Set("Content-Type", "application/json")
	res, _ := json.Marshal(book)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetAllBooks()
	res, _ := json.Marshal(newBooks)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 10, 64)

	if err != nil {
		fmt.Println("Failed to parse Book ID: %w", err)
		return
	}

	if ID <= 0 {
		fmt.Println("Book ID must be positive: %w", ID)
		return
	}

	bookDetails, _ := models.GetBookById(ID)
	res, _ := json.Marshal(bookDetails)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 10, 64)

	if err != nil {
		fmt.Println("Failed to parse Book ID: %w", err)
		return
	}

	if ID <= 0 {
		fmt.Println("Book ID must be positive: %w", ID)
		return
	}

	book, db_err := models.DeleteBook(ID)

	if db_err != nil {
		fmt.Println("Book ID must be positive: %w", db_err)
	}
	res, _ := json.Marshal(book)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	var updatedBook = &models.Book{}
	utils.ParseBody(r, updatedBook)

	vars := mux.Vars(r)
	bookId := vars["bookId"]

	ID, err := strconv.ParseInt(bookId, 10, 64)

	if err != nil {
		fmt.Println("Failed to parse Book ID: %w", err)
		return
	}

	if ID <= 0 {
		fmt.Println("Book ID must be positive: %w", ID)
		return
	}

	updatedBookResult, db_err := models.UpdateBook(ID, updatedBook)

	if db_err != nil {
		fmt.Println("Failed to update book: %w", ID)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	res, _ := json.Marshal(updatedBookResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
