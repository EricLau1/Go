package main

import "fmt"

func main() {
	msg := make(chan string)
	done := make(chan bool)
	go send(msg)
	go receive(msg, done)

	fmt.Println(<-done)
}

func send(msg chan string) {
	msg <- "Hello World!"
}

func receive(msg <-chan string, done chan bool) {
	fmt.Println(<-msg)
	done <- true
}
