package app

import (
	"fmt"
	"goapp/app/routes"
	"log"
	"net/http"
)

const PORT = 9000

func Run() {
	fmt.Println("Listening", PORT)
	listen(PORT)
}

func listen(port int) {
	r := routes.NewRouter()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), r))
}
