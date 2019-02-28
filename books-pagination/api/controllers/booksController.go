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
  total := models.CountBooks()
  limit := 5
  page, begin := utils.Pagination(r, limit)
  pages := (total / limit)
  if (total % limit) != 0 {
    pages++
  }
  books := models.PaginateBooks(begin, limit)
  if len(books) == 0 {
    utils.ToJson(w, books, http.StatusOK)
    return
  }
  utils.ToJson(w, struct{
    Docs []models.Book `json:"docs"`
    Total int `json:"total"`
    Limit int `json:"limit"`
    Page int `json:"page"`
    Pages int `json:"pages"`
  }{
    Docs: books,
    Total: len(books),
    Limit: limit,
    Page: page,
    Pages: pages,
  }, http.StatusOK)
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
