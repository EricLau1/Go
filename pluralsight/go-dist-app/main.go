package main

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

func main() {
	go client()
	go server()

	var a string
	fmt.Scanln(&a)
}

func client() {
	conn, ch, q := getQueue()
	defer conn.Close()
	defer ch.Close()
	msgs, err := ch.Consume(q.Name, "", true, false, false, false, nil)
	failOnError(err, "Registro do cliente falhou")
	for msg := range msgs {
		fmt.Printf("Mensagem recebida: %s\n", msg.Body)
	}
}

func server() {
	conn, ch, q := getQueue()
	defer conn.Close()
	defer ch.Close()

	msg := amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte("Hello RabbitMQ"),
	}

	for {
		ch.Publish("", q.Name, false, false, msg)
	}
}

func getQueue() (*amqp.Connection, *amqp.Channel, *amqp.Queue) {
	conn, err := amqp.Dial("amqp://localhost:5672")
	failOnError(err, "Conexão com o RabbitMQ falhou")
	ch, err := conn.Channel()
	failOnError(err, "Falhou ao abrir o canal")
	q, err := ch.QueueDeclare("Hello", false, false, false, false, nil)
	failOnError(err, "Declaração da fila falhou")
	return conn, ch, &q
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}
