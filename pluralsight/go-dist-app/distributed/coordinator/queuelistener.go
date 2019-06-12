package coordinator

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"pluralsight/go-dist-app/distributed/dto"
	"pluralsight/go-dist-app/distributed/qutils"

	"github.com/streadway/amqp"
)

const rmqURL = "amqp://guest:guest@localhost:5672"

type QueueListener struct {
	conn    *amqp.Connection
	ch      *amqp.Channel
	sources map[string]<-chan amqp.Delivery
	ea      *EventAggregator
}

func NewQueueListener(ea *EventAggregator) *QueueListener {
	ql := QueueListener{
		sources: make(map[string]<-chan amqp.Delivery),
		ea:      ea,
	}

	ql.conn, ql.ch = qutils.GetChannel(rmqURL)

	return &ql
}

func (ql *QueueListener) DiscoverySensors() {
	ql.ch.ExchangeDeclare(
		qutils.SensorDiscoveryExchange, //name string,
		"fanout",                       //kind string,
		false,                          //durable bool,
		false,                          //autoDelete bool,
		false,                          //internal bool,
		false,                          //noWait bool,
		nil)                            //args amqp.Table,
	ql.ch.Publish(
		qutils.SensorDiscoveryExchange, //exchange,
		"",                             //key,
		false,                          //mandatory,
		false,                          //immediate,
		amqp.Publishing{})              //msg amqp.Publishing
}

func (ql *QueueListener) ListenForNewSource() {
	q := qutils.GetQueue("", ql.ch, true)
	ql.ch.QueueBind(
		q.Name,       // name string
		"",           // key string
		"amq.fanout", // exchange string
		false,        // noWait
		nil)          // args amqp.Table

	msgs, _ := ql.ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil)

	ql.DiscoverySensors()

	fmt.Println("Ouvindo novas fontes")
	for msg := range msgs {
		fmt.Println("Nova fonte descoberta")
		ql.ea.PublishEvent("DataSourceDiscovered", string(msg.Body))
		sourceChan, _ := ql.ch.Consume(
			string(msg.Body),
			"",
			true,
			false,
			false,
			false,
			nil)
		if ql.sources[string(msg.Body)] == nil {
			ql.sources[string(msg.Body)] = sourceChan

			go ql.AddListener(sourceChan)
		}
	}
}

func (ql *QueueListener) AddListener(msgs <-chan amqp.Delivery) {
	for msg := range msgs {
		r := bytes.NewReader(msg.Body)
		d := gob.NewDecoder(r)
		sd := new(dto.SensorMessage)
		d.Decode(&sd)

		fmt.Printf("Mensagem recebida: %v\n", sd)

		ed := EventData{
			Name:      sd.Name,
			Timestamp: sd.Timestamp,
			Value:     sd.Value,
		}

		ql.ea.PublishEvent("MensagemRecebida_"+msg.RoutingKey, ed)
	}
}
