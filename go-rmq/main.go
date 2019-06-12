package main

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

func main() {
	conn, _ := amqp.Dial("amqp://guest:guest@localhost:5672")

	ch, _ := conn.Channel()

	q, _ := ch.QueueDeclare(
		"first.queue", // name string
		true,          // durable
		true,          // auto delete
		false,         // exclusive
		false,         // no wait
		nil)           // args amqp.Table

	ch.QueueBind(
		q.Name,       // name string
		"",           // key string
		"amq.fanout", // exchange string
		false,        // no wait
		nil)          // args amqp.Table

	msg := amqp.Publishing{
		Body: []byte("Hello World!"),
	}

	ch.Publish(
		"",     // exchange string
		q.Name, // key string
		false,  // mandatory
		false,  // immediate
		msg)    // msg amqp.Publishing

	msgs, err := ch.Consume(
		q.Name, //queue
		"",     // consumer
		true,   // auto Ack
		false,  // exclusive
		false,  // no local
		false,  // no Wait
		nil)    // args amqp.Table

	if err != nil {
		log.Fatal(err)
	}

	for m := range msgs {
		fmt.Println(string(m.Body))
	}
}
