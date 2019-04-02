package main

import "fmt"

type Person struct {
	Name string
	Gender string
}

func main() {
	people := []Person{Person{"Batgirl", "F"}, Person{"Batman", "M"}}
	fmt.Println("Capacidade =>", cap(people))
	fmt.Println("Tamanho =>", len(people))
}