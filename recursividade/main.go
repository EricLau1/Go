package main

import "fmt"

func fat(n int) int {
	if n < 1 {
		return 1
	}
	return n * fat(n-1)
}

func main() {
	fmt.Println(fat(5000))
}
