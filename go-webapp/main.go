package main

import (
	"fmt"
	"net/http"

	"./functions"
	"./models"
	"./routes"
	"./sessions"
	"./utils"
)

func main() {

	sessions.SessionOptions("localhost", "/", 3600, true)

	models.Init()

	// carregando os arquivos (views) html do diretorio templates
	utils.LoadTemplates("templates/*.html")

	fmt.Println("Servidor iniciado ", functions.Today(), "Porta: http://localhost:8080")
	r := routes.NewRouter()

	http.Handle("/", r)

	http.ListenAndServe(":8080", nil)

}
