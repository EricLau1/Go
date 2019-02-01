package main

import(
	"fmt"
	"log"
	"go-postgres-transactions/models"
)

func main() {
	models.TestConnection()
	CreateUser()
	CreateFeedback()
}

func CreateUser() {
	var name string
	fmt.Print("Name: ")
	fmt.Scan(&name)
	_, err := models.NewUser(models.User{0, name})
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("Usuário criado com sucesso!")
}

func CreateFeedback() {
	var user models.User
	fmt.Print("User Id: ")
	fmt.Scan(&user.Id)
	user, err := models.GetUserById(user.Id)
	if err != nil {
		log.Fatal(err)
		return
	}
	feedback := models.Feedback{User:user}
	fmt.Print("Comentário: ")
	fmt.Scan(&feedback.Comment)
	_, err = models.NewFeedback(feedback)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Printf("%s escreveu: %s\n", user.Name, feedback.Comment)
}