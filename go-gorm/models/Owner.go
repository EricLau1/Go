package models

import (
	"github.com/jinzhu/gorm"
)

// Types http://gorm.io/docs/models.html

type Owner struct {
	gorm.Model
	Name 	 string `gorm:"type:varchar(35)"`
	Email	 string `gorm:"type:varchar(40);unique_index"`
	Password string	`gorm:"type:varchar(60)"`
	Gender   string `gorm:"type:enum('F', 'M')"`
	Status   uint8	`gorm:"type:char(1)"`
}

func AutoMigrateOwner() {
	db := Connect()
	defer db.Close()
	
	ok := db.HasTable(&Owner{})
	if !ok {
		db.AutoMigrate(&Owner{})
		return
	}
	db.DropTableIfExists(&Owner{})
}

func NewOwner(owner Owner) int64 {
	db := Connect()
	defer db.Close()
	rs := db.Create(&owner)
	return rs.RowsAffected
}

func UpdateTableOwner(id uint) int64 {
	db := Connect()
	defer db.Close()
	rs := db.Table("owners").Where("id = ?", id).Update("name", "Emma Updated")
	return rs.RowsAffected
}


// atualizando varias colunas de uma vez
func UpdatesOwner(owner Owner) int64 {
	db := Connect()
	defer db.Close()
	rs := db.Model(&owner).Updates(
		map[string]interface{}{
			"Name": "Batman",
			"Email": "brucewayne@email.com",
			"Password": "321",
			"Gender": "M",
			"Status": 1,
		},
	)
	return rs.RowsAffected
}

func UpdateColumnsOwner(owner Owner) int64 {
	db := Connect()
	defer db.Close()
	rs := db.Model(&owner).UpdateColumns(
		map[string]interface{}{
			"Name": "Susan",
			"Email": "susan@email.com",
			"Password": "123456",
			"Gender": "F",
			"Status": 3,
		},
	)
	return rs.RowsAffected
}

