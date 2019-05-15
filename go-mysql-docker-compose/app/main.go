package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("Server running...")
	http.HandleFunc("/db", handler)
	log.Fatal(http.ListenAndServe(":9000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(os.Getenv("DATABASE_HOST"))
	host := os.Getenv("DATABASE_HOST")
	if host == "" {
		host = "127.0.0.1"
	}
	db, err := sql.Open("mysql", fmt.Sprintf("root:@root@tcp(%s:3306)/stormlight", host))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	w.Write([]byte("mysql on..."))
}
