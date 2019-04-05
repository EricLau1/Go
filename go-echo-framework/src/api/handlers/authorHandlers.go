package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo"
)

type Author struct {
	Name string `json:"name"`
}

// Fazendo parse do Body com NewDecoder
func AddAuthor(c echo.Context) error {
	author := Author{}
	defer c.Request().Body.Close()

	err := json.NewDecoder(c.Request().Body).Decode(&author)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity)
	}
	return c.JSON(http.StatusCreated, author)
}
