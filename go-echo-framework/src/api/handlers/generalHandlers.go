package handlers

import (
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

func Hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello world!")
}

func MainJwt(c echo.Context) error {

	user := c.Get("user")
	token := user.(*jwt.Token)

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return echo.NewHTTPError(http.StatusUnauthorized)
	}

	log.Println("User Name: ", claims["name"].(string), "User ID: ", claims["jti"])

	return c.String(http.StatusOK, "Secret Page By Json Web Token")
}

func MainAdmin(c echo.Context) error {
	return c.String(http.StatusOK, "Secret Page By Admin")
}

func MainCookie(c echo.Context) error {
	return c.String(http.StatusOK, "Secret Page By Cookie")
}
