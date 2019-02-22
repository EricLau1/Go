package test

import (
	"fmt"
	"go-gorm/models"
)

func TestCreateTableUser() {
	ok := models.TableUserExists()
	if !ok {
		models.CreateTableUser()
		fmt.Println("Tabela usuário criada com sucesso!")
		return
	}
	models.DropTableUser()
	fmt.Println("Tabela usuário foi dropada!")
}

func TestNewUser() {
	user := models.User{Nickname: "Jane Doe"}
	rows := models.NewUser(user)
	fmt.Printf("%d linhas afetadas\n", rows)
}

func TestFindAndUpdateUser() {
	user := models.User{Nickname: "Jane Doe"}
	rows := models.FindAndUpdateUser(user)
	fmt.Printf("%d linhas afetadas\n", rows)
}

func TestTableDeleteUser() {
	rows := models.TableDeleteUser(1)
	fmt.Printf("%d linhas afetadas\n", rows)
}

func TestTableUserDelete() {
	rows := models.DeleteUser(1)
	fmt.Printf("%d linhas afetadas\n", rows)
}

func TestDeleteAllUsers() {
	rows := models.DeleteAllUsers()
	fmt.Printf("%d linhas afetadas\n", rows)
}