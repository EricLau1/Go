package controllers

import (
  "net/http"
  "strconv"
  "encoding/json"
  "books/api/models"
  "books/api/utils"
  "github.com/gorilla/mux"
)

func GetBooks(w http.ResponseWriter, r *http.Request) {
  limit := 10
  total := models.CountBooks()
  page, begin := utils.Pagination(r, limit)
  books := models.PaginateBooks(begin, limit)
  pages := (total / int64(limit)) + (total % int64(limit))
  utils.ToJson(w, struct{
    Docs  []models.Book `json:"docs"`
    Limit int           `json:"limit"`
    Page  int           `json:"page"`
    Pages int64         `json:"pages"`
  }{
    Docs: books,
    Limit: limit,
    Page: page,
    Pages: pages,
  }, 
  http.StatusOK)
}

func GetBook(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  id, _ := strconv.ParseUint(vars["id"], 10, 64)
  book := models.GetBookById(id)
  utils.ToJson(w, book, http.StatusOK)
}

func PostBook(w http.ResponseWriter, r *http.Request) {
  body := utils.BodyParser(r)
  var book models.Book
  err := json.Unmarshal(body, &book)
  if err != nil {
    utils.ToJson(w, err.Error(), http.StatusUnprocessableEntity)
    return
  }
  err = models.NewBook(book)
  if err != nil {
    utils.ToJson(w, err.Error(), http.StatusUnprocessableEntity)
    return
  }
  utils.ToJson(w, "New Book Created", http.StatusCreated) 
}

func PutBook(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  id, _ := strconv.ParseUint(vars["id"], 10, 64)
  body := utils.BodyParser(r)
  var book models.Book
  err := json.Unmarshal(body, &book)
  if err != nil {
    utils.ToJson(w, err.Error(), http.StatusUnprocessableEntity)
    return
  }
  book.Id = id
  rows, err := models.UpdateBook(book)
  if err != nil {
    utils.ToJson(w, err.Error(), http.StatusUnprocessableEntity)
    return
  }
  utils.ToJson(w, rows, http.StatusOK)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  id, _ := strconv.ParseUint(vars["id"], 10, 64)
  _, err := models.DeleteBook(id)
  if err != nil {
    utils.ToJson(w, err.Error(), http.StatusUnprocessableEntity)
    return
  }
  w.WriteHeader(http.StatusNoContent)
}
