package main

import (
	"fmt"
)

func Square(x int) int {
	return x * x
}

func main() {
	fmt.Println("Go Testing...")

	n := 2
	r := Square(n)

	fmt.Printf("square(%d): %d\n", n, r)
}
