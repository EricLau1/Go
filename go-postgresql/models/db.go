package models

import (
  "fmt"
  "log"
  "database/sql"
  _ "github.com/lib/pq"
)

const(
  DRIVER = "postgres"
  PORT   = 5432 // porta default do postgres
  DBNAME = "test"
  USER   = "postgres"
  PASS   = "@root"
)

// user=%s password=%s dbname=%s sslmode=disable

func Connect() *sql.DB {

  URL := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", USER, PASS, DBNAME)

  con, err := sql.Open(DRIVER, URL)

  if err != nil {
    panic(err.Error())
  }

  return con
}

func TestConnection(){

  con := Connect()

  err := con.Ping()

  if err != nil {
    log.Fatal(err)
  }

  fmt.Println("Conectado com sucesso!")

  defer con.Close()
}
