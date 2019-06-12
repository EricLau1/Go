package qutils

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

const SensorDiscoveryExchange = "SensorDiscovery"
const PersistReadingsQueue = "PersistReading"

func GetChannel(url string) (*amqp.Connection, *amqp.Channel) {
	conn, err := amqp.Dial(url)
	failOnError(err, "Falha ao estabelecer conexão com o Message Broker")
	ch, err := conn.Channel()
	failOnError(err, "Falha ao receber o canal para a conexão")
	return conn, ch
}

// Retorna a fila
func GetQueue(name string, ch *amqp.Channel, autoDelete bool) *amqp.Queue {
	q, err := ch.QueueDeclare(
		name,       // name string
		false,      // durable bool
		autoDelete, // autoDelete bool
		false,      // exclusive bool
		false,      // noWait bool
		nil)        // args amqp.Table
	failOnError(err, "Falha ao declarar a fila")

	return &q
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}
