package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Running...")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Go Docker Tutorial")
	})

	log.Fatal(http.ListenAndServe(":8081", nil))
}
