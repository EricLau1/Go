package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/labstack/echo"
)

type Book struct {
	Title  string `json:"title"`
	Author Author `json:"author"`
}

var books []Book

func GetBooks(c echo.Context) error {
	// parametros da URL
	author := c.QueryParam("author")
	title := c.QueryParam("title")

	// parametro definido no Path
	dataType := c.Param("data")

	switch dataType {
	case "string":
		// retorna uma text/plain
		return c.String(http.StatusOK, fmt.Sprintf("%s, written by %s", strings.ToUpper(title), author))
	case "json":
		return c.JSON(http.StatusOK, map[string]string{
			"title":  title,
			"author": author,
		})
	default:
		// retorna um JSON
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Inform JSON or String in the URL",
		})
	}

}

// Fazendo parse do Body com Unmarshal
func AddBook(c echo.Context) error {
	book := Book{}
	defer c.Request().Body.Close()
	body, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		return c.String(http.StatusUnprocessableEntity, err.Error())
	}
	err = json.Unmarshal(body, &book)
	if err != nil {
		return c.String(http.StatusUnprocessableEntity, err.Error())
	}
	books = append(books, book)
	return c.JSON(http.StatusCreated, books)
}
