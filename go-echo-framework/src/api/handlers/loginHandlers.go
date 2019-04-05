package handlers

import (
	"net/http"
	"time"

	"github.com/labstack/echo"
)

func Login(c echo.Context) error {
	username := c.QueryParam("username")
	password := c.QueryParam("password")

	if username == "admin" && password == "123456" {
		cookie := &http.Cookie{}

		cookie.Name = "sessionID"
		cookie.Value = "mindawakebodyasleep"
		cookie.Expires = time.Now().Add(30 * time.Minute)

		c.SetCookie(cookie)

		return c.String(http.StatusOK, "Logged successfully!")
	}
	return c.String(http.StatusUnauthorized, "Login error. Check your username or password")
}
