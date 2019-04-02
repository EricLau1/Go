package main

import (
	"log"
	"net/http"
	"os"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	port := os.Getenv("APP_PORT")
	secretKey := os.Getenv("SECRET_KEY")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(secretKey))
	})
	log.Fatal(http.ListenAndServe(port, nil))
}