package handlers

import (
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

type JwtClaims struct {
	Name string `json:"name"`
	jwt.StandardClaims
}

func createJwt() (string, error) {
	claims := JwtClaims{
		"admin",
		jwt.StandardClaims{
			Id:        "main_user_id",
			ExpiresAt: time.Now().Add(10 * time.Minute).Unix(),
		},
	}
	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	token, err := rawToken.SignedString([]byte("mySecretKey"))
	if err != nil {
		log.Println(err)
		return "", err
	}
	return token, nil
}

func JwtLogin(c echo.Context) error {
	username := c.QueryParam("username")
	password := c.QueryParam("password")

	if username == "admin" && password == "123456" {

		token, err := createJwt()
		if err != nil {
			c.String(http.StatusUnauthorized, err.Error())
		}

		return c.String(http.StatusOK, token)
	}
	return c.String(http.StatusUnauthorized, "Login error. Check your username or password")
}
