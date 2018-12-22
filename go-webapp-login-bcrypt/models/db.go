package models

import(
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

const DRIVER   = "mysql"
const USER     = "root"
const PASSWORD = "@root"
const DBNAME   = "test"

func Connect() (*sql.DB, error) {

	URL := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s", USER, PASSWORD, DBNAME) 

	con, err := sql.Open( DRIVER, URL )

	if err != nil {

		return nil, err

	}

	err = con.Ping() 

	if err != nil {

		return nil, err

	}

	return con, nil

}