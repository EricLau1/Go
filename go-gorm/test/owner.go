package test

import (
	"fmt"
	"go-gorm/models"
)

func TestAutoMigrateOwner() {
	models.AutoMigrateOwner()
}

func TestNewOwner() {
	owners := []models.Owner{
		models.Owner{Name:"Emma", Email: "emma@email.com", Password: "123", Gender: "F", Status: 1},
		models.Owner{Name:"Ana", Email: "ana@email.com", Password: "123", Gender: "F", Status: 1},
		models.Owner{Name:"Matt", Email: "matt@email.com", Password: "123", Gender: "M", Status: 1},
	}
	for _, o := range owners {
		rows := models.NewOwner(o)
		fmt.Printf("%d linhas afetadas\n", rows)
	} 
}

func TestUpdateTableOwner() {
	rows := models.UpdateTableOwner(1)
	fmt.Printf("%d linhas afetadas\n", rows)
}

func TestUpdatesOwner() {
	owner := models.Owner{Name:"Emma", Email: "emma@email.com", Password: "123", Gender: "F", Status: 0}
	rows := models.UpdatesOwner(owner)
	fmt.Printf("%d linhas afetadas\n", rows)
}

func TestUpdateColumnsOwner() {
	owner := models.Owner{Name:"Emma", Email: "emma@email.com", Password: "123", Gender: "F", Status: 0}
	rows := models.UpdateColumnsOwner(owner)
	fmt.Printf("%d linhas afetadas\n", rows)
}