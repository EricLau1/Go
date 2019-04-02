package main

import (
	"fmt"
	"reflect"
	"time"
)

func main() {

	name := "Jane Doe"
	date := time.Now()
	PI := 3.14159
	var fruits [3]string
	fruits[0] = "Laranja"
	fruits[1] = "Banana"
	fruits[2] = "Maçã"

	fib := []int{0, 1, 2, 3, 5, 8, 13}

	fmt.Printf("Nome: %s, Type => %s\n", name, reflect.TypeOf(name))
	fmt.Printf("PI: %.6f, Type => %s\n", PI, reflect.TypeOf(PI))
	fmt.Printf("Data: %+v, Type => %s\n", date, reflect.TypeOf(date))
	fmt.Printf("Frutas: %v, Type => %s\n", fruits, reflect.TypeOf(fruits))
	fmt.Printf("Fibonacci: %v, Type => %s\n", fib, reflect.TypeOf(fib))

}
