package main

import (
	"bufio"
	"fmt"
	"go-rabbitmq/rabbitmq"
	"os"
)

func main() {
	var option int
	fmt.Println("[1] - Send")
	fmt.Println("[2] - Receive")
	fmt.Print("Escolha: ")
	fmt.Scan(&option)

	switch option {
	case 1:
		consoleReader := bufio.NewReader(os.Stdin)
		fmt.Print("Message: ")
		message, _ := consoleReader.ReadString('\n')
		rabbitmq.SendMessageRMQ(message)
	case 2:
		rabbitmq.ReceiveMessageRMQ()
	default:
		fmt.Println("Invalid option")
	}

}
