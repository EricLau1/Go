package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan bool)
	go task(done)
	fmt.Println(<-done)
}

func task(ch chan bool) {
	time.Sleep(time.Second * 3)
	ch <- true
}
