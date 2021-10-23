package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/ong-gtp/book-mgt/pkg/errors"
	"github.com/ong-gtp/book-mgt/pkg/models"
	"github.com/ong-gtp/book-mgt/pkg/utils"
)

func NotFound(w http.ResponseWriter, r *http.Request) {
	utils.OkEmpty("Invalid URL", w)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	book := models.Book{}
	utils.ParseBody(r, &book)
	fmt.Println(book)
	b, err := book.CreateBook()
	errors.DBErrorCheck(err)

	res, err2 := json.Marshal(b)
	errors.ErrorCheck(err2)

	utils.Ok(res, w)
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	newBook, err := models.GetAllBooks()
	errors.DBErrorCheck(err)

	res, er2 := json.Marshal(newBook)
	errors.ErrorCheck(er2)

	utils.Ok(res, w)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]

	id, err := strconv.ParseInt(bookId, 0, 0)
	errors.ErrorCheck(err)

	book, dbInst := models.GetBookById(id)
	errors.DBErrorCheck(dbInst)

	if book.ID == 0 {
		utils.OkEmpty("Invalid book", w)
		return
	}

	res, er2 := json.Marshal(book)
	errors.ErrorCheck(er2)

	utils.Ok(res, w)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]

	id, err := strconv.ParseInt(bookId, 0, 0)
	errors.ErrorCheck(err)
	book, dbInst := models.GetBookById(id)
	errors.DBErrorCheck(dbInst)
	if book.ID == 0 {
		utils.OkEmpty("Invalid book", w)
		return
	}

	_, dbInst2 := models.DeleteBook(id)
	errors.DBErrorCheck(dbInst2)

	utils.OkEmpty("Deleted successfully", w)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]

	id, err := strconv.ParseInt(bookId, 0, 0)
	errors.ErrorCheck(err)

	// get book from database
	bk, dbInst := models.GetBookById(id)
	errors.DBErrorCheck(dbInst)
	if bk.ID == 0 {
		utils.OkEmpty("Invalid book", w)
		return
	}

	// update book with json from route body
	book := models.Book{} // the response body should conform with book type
	// decode response body into book
	decoder := json.NewDecoder(r.Body)
	err3 := decoder.Decode(&book)
	errors.ErrorCheck(err3)
	// update bk with request
	bk.Name = book.Name
	bk.Author = book.Author
	bk.Publication = book.Publication
	newBk, err2 := bk.UpdateBook()
	errors.DBErrorCheck(err2)
	// prepare updated data for response
	res, err4 := json.Marshal(newBk)
	errors.ErrorCheck(err4)

	utils.Ok(res, w)
}
