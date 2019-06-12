package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type Item struct {
	Title string
	Body  string
}

func main() {
	var reply Item
	var db []Item

	client, err := rpc.DialHTTP("tcp", "localhost:4040")
	if err != nil {
		log.Fatal(err)
	}

	i1 := Item{"first", "Item 1"}
	i2 := Item{"second", "Item 2"}
	i3 := Item{"third", "Item 3"}

	client.Call("API.AddItem", i1, &reply)
	client.Call("API.AddItem", i2, &reply)
	client.Call("API.AddItem", i3, &reply)
	client.Call("API.GetDB", "", &db)

	fmt.Println("Database:", db)
}
