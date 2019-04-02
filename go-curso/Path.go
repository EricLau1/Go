package main

import "path"
import "fmt"

func main() {
	fmt.Println(path.Dir(""))
	fmt.Println(path.Join(".", "monitoramento"))
}