package main

import (
	"fmt"
	"time"
	"log"
	"net/http"
	"io/ioutil"
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
	client := &http.Client{}

	// prepara uma requisição
	req, _ := http.NewRequest("GET", "http://localhost:9000/", nil)

	// grava o token no header
	req.Header.Set("Token", validToken)

	// envia a requisição
	res, err := client.Do(req)
	if err != nil {
		fmt.Fprintf(w, "Error: %s", err.Error())
	}

	// resposta da requisição
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	// mostra o conteúdo da requisição
	fmt.Fprintf(w, string(body))
}

func handleRequests() {
	http.HandleFunc("/", homeGetHandler)
	log.Fatal(http.ListenAndServe(":3000", nil))
}