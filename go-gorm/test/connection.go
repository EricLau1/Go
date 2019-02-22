package test

import (
	"fmt"
	"go-gorm/models"
)

func TestConnection() {
	db := models.Connect()
	defer db.Close()
	fmt.Println("Banco de dados conectado")
}

