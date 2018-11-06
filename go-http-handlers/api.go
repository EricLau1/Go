package main

/* Video tutorial : https://www.youtube.com/watch?v=HiCph4_fN6M */

/*
	pacote mux:

	\> go get github.com/gorilla/mux

	\> go get github.com/gorilla/handlers

*/

import (
	"fmt"
	"net/http"

	"github.com/gorilla/handlers"

	"github.com/gorilla/mux"
)

func RootEndPoint(response http.ResponseWriter, request *http.Request) {

	response.Write([]byte("Hello world!"))

}

func main() {

	fmt.Println("Server running...")

	router := mux.NewRouter()

	headers := handlers.AllowedHeaders([]string{"X-Request", "Content-type", "application/json", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	origins := handlers.AllowedOrigins([]string{"*"})

	router.HandleFunc("/", RootEndPoint).Methods("GET")

	http.ListenAndServe(":3000", handlers.CORS(headers, methods, origins)(router))

}
