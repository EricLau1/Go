package main

import (
	fmt "fmt"
	"log"

	proto "github.com/golang/protobuf/proto"
)

func main() {

	user := &User{Name: "Jhon Doe", Age: 10}

	product := &Product{Description: "Laranja", Price: 10.0, Quantity: 10, Status: true}

	data, err := proto.Marshal(user)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(data)
	fmt.Println(user.GetName())
	fmt.Println(user.GetAge())

	fmt.Println(product)
}
