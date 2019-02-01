package models

import(
	"fmt"
  "log"
	"database/sql"
	_ "github.com/lib/pq"
)

const(
	DRIVER = "postgres"
	USER   = "postgres"
	PASS   = "@root"
	DBNAME = "dbtransactions"
)

func Connect() *sql.DB {
	URL := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", USER, PASS, DBNAME)
	db, err := sql.Open(DRIVER, URL)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return db
}

func TestConnection() {
	con := Connect()
	defer con.Close()
	err := con.Ping()
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("Banco de dados conectado!")
}
