package main

import (
	"fmt"
	"time"
	"log"
	"net/http"
	jwt "github.com/dgrijalva/jwt-go"
)

var mySigingKey = []byte("s3cr3tk3y")

func GenerateJWT() (string, error) {
	
	token := jwt.New(jwt.SigningMethodHS256)
	
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["user"] = "Jane Doe"
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(mySigingKey)

	if err != nil {
		fmt.Errorf("Something went wrong: %s", err.Error()) // Algo deu errado
		return "", err
	}
	return tokenString, nil
}

func main() {
	fmt.Println("My Simple Client")	
	handleRequests()
}

func homeGetHandler(w http.ResponseWriter, r *http.Request) {
	validToken, err := GenerateJWT()
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	fmt.Fprintf(w, validToken)
}

func handleRequests() {
	http.HandleFunc("/", homeGetHandler)
	log.Fatal(http.ListenAndServe(":3000", nil))
}