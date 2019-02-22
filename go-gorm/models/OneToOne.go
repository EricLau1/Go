package models

/* 

Place.go
package model
import ()
type Place struct {
 ID int `gorm:primary_key`
 Name string
 Town Town
 TownId int `gorm:"ForeignKey:id"` //this foreignKey tag didn't works
}

Town.go
package model
import ()
type Town struct {
 ID int `gorm:"primary_key"`
 Name string
}

main.go
package main

import (
 _ "database/sql"
 _ "github.com/go-sql-driver/mysql"
 "github.com/jinzhu/gorm"
 "26_GO_GORM/One2One_Relationship/model"
 "fmt"
)

//var Db *gorm.Db

func main() {
 //Init Db connection

 Db, _ := gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/testmapping?charset=utf8&parseTime=True")
 defer Db.Close()

 Db.DropTableIfExists(&model.Place{}, &model.Town{})

 Db.AutoMigrate(&model.Place{}, &model.Town{})
 //We need to add foreign keys manually.
 Db.Model(&model.Place{}).AddForeignKey("town_id", "towns(id)", "CASCADE", "CASCADE")

 t1 := model.Town{
 Name: "Pune",
 }
 t2 := model.Town{
 Name: "Mumbai",
 }
 t3 := model.Town{
 Name: "Hyderabad",
 }

 p1 := model.Place{
 Name: "Katraj",
 Town: t1,
 }
 p2 := model.Place{
 Name: "Thane",
 Town: t2,
 }
 p3 := model.Place{
 Name: "Secundarabad",
 Town: t3,
 }

 Db.Save(&p1) //Saving one to one relationship
 Db.Save(&p2)
 Db.Save(&p3)

 fmt.Println("t1==>", t1, "p1==>", p1)
 fmt.Println("t2==>", t2, "p2s==>", p2)
 fmt.Println("t2==>", t3, "p2s==>", p3)

 //Delete
 Db.Where("name=?", "Hyderabad").Delete(&model.Town{})

 //Update
 Db.Model(&model.Place{}).Where("id=?", 1).Update("name", "Shivaji Nagar")

 //Select
 places := model.Place{}
 towns := model.Town{}
 fmt.Println("Before Association", places)
 Db.Where("name=?", "Shivaji Nagar").Find(&places)
 fmt.Println("After Association", places)
 err := Db.Model(&places).Association("town").Find(&places.Town).Error
 fmt.Println("After Association", towns, places)
 fmt.Println("After Association", towns, places, err)

 defer Db.Close()
}

*/