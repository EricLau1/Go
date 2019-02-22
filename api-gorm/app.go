package main

import (
	"net/http"
	"log"
	"fmt"
	"api-gorm/models"
	"api-gorm/routes"
)

func main() {
	models.AutoMigrations()
	listen(9000)
}

func listen(p int) {
	port := fmt.Sprintf(":%d", p)
	fmt.Printf("\n\nListening on port %s...\n\n", port)
	r := routes.NewRouter()
	log.Fatal(http.ListenAndServe(port, r))
}