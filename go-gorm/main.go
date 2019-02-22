package main 

import (
	"fmt"
	"go-gorm/models"
)

func main() {

	FindFirst()
	FindLast()
	FindAll() 
}

// retorna a primeira linha da tabela
func FindFirst() {
	fmt.Println("Primeira linha:")
	db := models.Connect()
	defer db.Close()
	var owner models.Owner
	db.First(&owner)
	fmt.Println(owner)
}

func FindLast() {
	fmt.Println("Ultima linha:")
	db := models.Connect()
	defer db.Close()
	var owner models.Owner
	db.Last(&owner)
	fmt.Println(owner)	
}

func FindAll() {
	fmt.Println("Lista:")
	db := models.Connect()
	defer db.Close()
	var owner []models.Owner
	db.Find(&owner)
	fmt.Println(owner)
}	

/* 

	db.Where("address = ?", "Los Angeles").First(&user)
	//SELECT * FROM user_models WHERE address=’Los Angeles’ limit 1;
	
	db.Where("address = ?", "Los Angeles").Find(&user)
	//SELECT * FROM user_models WHERE address=’Los Angeles’;
	
	db.Where("address <> ?", "New York").Find(&user)
	//SELECT * FROM user_models WHERE address<>’Los Angeles’;
	
	// IN
	db.Where("name in (?)", []string{"John", "Martin"}).Find(&user)
	
	// LIKE
	db.Where("name LIKE ?", "%ti%").Find(&user)
	
	// AND
	db.Where("name = ? AND address >= ?", "Martin", "Los Angeles").Find(&user)
*/


