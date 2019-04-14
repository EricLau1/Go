package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func hardTask(counter int) {
	defer wg.Done()
	for i := 1; i <= 10; i++ {
		time.Sleep(time.Second)
		fmt.Printf("Hard task %d...\n", counter)
	}
	fmt.Printf("END %d\n", counter)
}

func main() {

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go hardTask(i + 1)
		go hardTask(i + 1)
	}

	wg.Wait()
}
