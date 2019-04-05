package handlers

import (
	"net/http"

	"github.com/labstack/echo"
)

type Contact struct {
	Email string `json:"email"`
}

type Contacts []Contact

var contacts Contacts

// Fazendo parse do Body com Bind
func AddContact(c echo.Context) error {
	contact := Contact{}

	err := c.Bind(&contact)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity)
	}
	contacts = append(contacts, contact)
	return c.JSON(http.StatusCreated, contacts)
}
