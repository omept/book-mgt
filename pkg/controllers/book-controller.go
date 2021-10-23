package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/ong-gtp/book-mgt/pkg/errors"
	"github.com/ong-gtp/book-mgt/pkg/models"
)

func CreateBook(w http.ResponseWriter, r *http.Request) {

}

func GetBook(w http.ResponseWriter, r *http.Request) {
	newBook, err := models.GetAllBooks()
	errors.ErrorCheck(err)

	res, err := json.Marshal(newBook)
	errors.ErrorCheck(err)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]

	id, err := strconv.ParseInt(bookId, 0, 0)
	errors.ErrorCheck(err)

	book, dbInst := models.GetBookById(id)
	errors.DBErrorCheck(dbInst)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {

}

func DeleteBook(w http.ResponseWriter, r *http.Request) {

}
