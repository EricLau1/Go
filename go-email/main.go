package main

import (
	"fmt"
	"log"
	"net/smtp"
)

func main() {

	hostURL := "smtp.gmail.com"
	hostPort := "587"
	emailSender := "eric.devtt@gmail.com"
	password := "SENHA DO EMAIL_SENDER"
	emailReceiver := "ericlau.oliveira@gmail.com"
	
	emailAuth := smtp.PlainAuth(
		"",
		emailSender,
		password,
		hostURL,
	)

	msg := []byte("To: " + emailReceiver + "\r\n" +
				  "Subject: Enviando Email com Go \r\n" +
				  "Olá! Tudo bem com você?")

	err := smtp.SendMail(
		hostURL + ":" + hostPort,
		emailAuth,
		emailSender,
		[]string{emailReceiver}, 
		msg)
	
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("Email Enviado!")
}