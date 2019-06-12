package main

import (
	"bytes"
	"encoding/gob"
	"log"
	"pluralsight/go-dist-app/distributed/datamanager"
	"pluralsight/go-dist-app/distributed/dto"
	"pluralsight/go-dist-app/distributed/qutils"
)

const rmqURL = "amqp://guest:guest@localhost:5672"

func main() {
	conn, ch := qutils.GetChannel(rmqURL)
	defer conn.Close()
	defer ch.Close()

	msgs, err := ch.Consume(
		qutils.PersistReadingsQueue,
		"",
		false,
		true,
		false,
		false,
		nil)

	if err != nil {
		log.Fatalln("Falha ao acessar as mensagens")
	}
	for msg := range msgs {
		buf := bytes.NewReader(msg.Body)
		dec := gob.NewDecoder(buf)
		sd := &dto.SensorMessage{}
		dec.Decode(sd)

		err := datamanager.SaveReader(sd)

		if err != nil {
			log.Printf("Falha ao salvar. Sensor %v. Error: %s", sd.Name, err.Error())
		} else {
			msg.Ack(false)
		}
	}
}
