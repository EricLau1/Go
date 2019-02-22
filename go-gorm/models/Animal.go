package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Animal struct {
	gorm.Model // id será gerado automaticamente
	Name string `gorm:"type:varchar(20)"`
	Birthday *time.Time
}

func DropTableAnimal() {
	db := Connect()
	defer db.Close()
	db.Debug().DropTableIfExists(&Animal{})
}

// cria a tabela automaticamente baseada no modelo
func AutoMigrateAnimal() {
	db := Connect()
	defer db.Close()
	db.Debug().AutoMigrate(&Animal{})
}

func NewAnimal(animal Animal) int64 {
	db := Connect()
	defer db.Close()
	rs := db.Create(&animal)
	return rs.RowsAffected
}

// atualiza todas as linhas
func UpadateAnimalByModel(animal Animal) int64 {
	db := Connect()
	defer db.Close()
	rs := db.Model(&animal).Update("Name", "Super Dog")
	return rs.RowsAffected
}

// atualiza todas as linhas
func UpadateColumnAnimalByModel(animal Animal) int64 {
	db := Connect()
	defer db.Close()
	rs := db.Model(&animal).UpdateColumn("Name", "Pikachu")
	return rs.RowsAffected
}

// atualiza a linha que for encontrada
func FindAnimalAndUpdate(animal Animal) int64 {
	db := Connect()
	defer db.Close()
	rs := db.Find(&animal).Update("Name", "Toggepi")
	return rs.RowsAffected
}