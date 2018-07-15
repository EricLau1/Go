package main

/*
 * Conteudo baseado em: https://www.youtube.com/watch?v=joVuFbAzPYw&t
 *
 */

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {

	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))

}

func main() {

	fmt.Println("url disponivel: http://localhost:8080/")

	http.HandleFunc("/", index)
	http.HandleFunc("/sobre", sobre)
	http.HandleFunc("/contato", contato)
	http.HandleFunc("/aplicar", aplicar)

	// Função para habilitar os pacotes da pasta public
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))

	// ignora o icone do titulo
	http.Handle("/favicon.ico", http.NotFoundHandler())

	http.ListenAndServe(":8080", nil)
}

type pageData struct {
	Title string
	Nome  string
}

func index(w http.ResponseWriter, r *http.Request) {

	pd := pageData{
		Title: "Index",
	}

	err := tpl.ExecuteTemplate(w, "index.gohtml", pd)

	if err != nil {
		log.Println("LOGGED", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	fmt.Println(r.URL.Path)
	fmt.Println("let's go explore site!")
}

func sobre(w http.ResponseWriter, r *http.Request) {

	pd := pageData{
		Title: "Sobre",
	}

	err := tpl.ExecuteTemplate(w, "sobre.gohtml", pd)

	if err != nil {
		log.Println("LOGGED", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

}

func contato(w http.ResponseWriter, r *http.Request) {

	pd := pageData{
		Title: "Contato",
	}

	err := tpl.ExecuteTemplate(w, "contato.gohtml", pd)

	if err != nil {
		log.Println("LOGGED", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

}

func aplicar(w http.ResponseWriter, r *http.Request) {

	pd := pageData{
		Title: "Aplicar",
	}

	var nome string
	if r.Method == http.MethodPost {

		nome = r.FormValue("nome") // pega o valor do input com name="nome"

		pd.Nome = nome

	}

	err := tpl.ExecuteTemplate(w, "aplicar.gohtml", pd)

	if err != nil {
		log.Println("LOGGED", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

}
