package main

import "fmt"

func diagonalSecundaria(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if j == n-1-i {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println("")
	}
}

func main() {
	diagonalSecundaria(5)
}
