package main

// download package go get -u golang.org/x/crypto/bcrypt

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {

	password := "secret"

	cost := bcrypt.DefaultCost

	hash, err := bcrypt.GenerateFromPassword([]byte(password), cost)

	if err != nil {
		panic(err.Error())
	}

	fmt.Println(hash)

	var passwd string
	fmt.Print("Password: ")
	fmt.Scan(&passwd)

	fmt.Println("Voce digitou: ", passwd)

	error := bcrypt.CompareHashAndPassword(hash, []byte(passwd))

	if error != nil {
		fmt.Println(error.Error())
		return
	}

	fmt.Println("Senhas iguais")
}
