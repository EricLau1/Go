package main

/*
  Video Original: https://www.youtube.com/watch?v=DWNozbk_fuk

  importar o pacote do driver mysql pelo cmd:

    > go get github.com/go-sql-driver/mysql

*/

import (
  "fmt"
  "database/sql"
  _ "github.com/go-sql-driver/mysql"
)

type User struct {
  Id   int `json:"id"`
  Nome string `json:"nome"`
}

func main() {
  fmt.Println("Go MySQL Tutorial")

  db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/testGO")

  if err != nil {
    panic(err.Error())
  }

  defer db.Close()

  fmt.Println("Banco conectado com sucesso!")

  nomes := []string{
    "Merry",
    "Shane",
    "Marri",
    "Nathan",
    "Patrick",
  }

  insert, err := db.Query("INSERT INTO users (nome) VALUES ('"+ nomes[2] +"')")

  if err != nil {
    panic(err.Error())
  }

  defer insert.Close()

  fmt.Println("Insert realizado com sucesso!")

  rs, err := db.Query("SELECT * FROM users")

  if err != nil {
    panic(err.Error())
  }

  for rs.Next() {
    var user User

    err = rs.Scan(&user.Id, &user.Nome)

    if err != nil {
      panic(err.Error())
    }

    fmt.Println("Id: ", user.Id, ", Nome: ", user.Nome)

  }

  fmt.Println("Select realizado com sucesso!")

}
