package test

import (
	"fmt"
	"time"
	"go-gorm/models"
)

func TestAutoMigrateAnimal() {
	models.DropTableAnimal()
	models.AutoMigrateAnimal()
}

func TestNewAnimal() {
	now := time.Now()
	animal := models.Animal{Name: "Rex", Birthday: &now}
	rows := models.NewAnimal(animal)
	fmt.Printf("%d linhas afetadas\n", rows)
}

func TestUpadateAnimalByModel() {
	animal := models.Animal{Name:"Rex"}
	rows := models.UpadateAnimalByModel(animal)
	fmt.Printf("%d linhas afetadas\n", rows)
}

func TestUpadateColumnAnimalByModel() {
	animal := models.Animal{Name:"Rex"}
	rows := models.UpadateColumnAnimalByModel(animal)
	fmt.Printf("%d linhas afetadas\n", rows)
}

func TestFindAnimalAndUpdate() {
	animal := models.Animal{Name:"Rex"}
	rows := models.FindAnimalAndUpdate(animal)
	fmt.Printf("%d linhas afetadas\n", rows)
}