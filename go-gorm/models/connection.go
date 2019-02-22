package models 

import (
	"fmt"
	"log"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

const (
	USER = "root"
	PASS = "@root"
	DBNAME = "gorm" 
)

func Connect() *gorm.DB {
	URL := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?charset=utf8&parseTime=True", USER, PASS, DBNAME)
	db, err := gorm.Open("mysql", URL)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return db
}
