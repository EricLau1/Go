package main

import (
	"errors"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

type Item struct {
	Title string
	Body  string
}

type API int

var database []Item

var (
	ErrItemNotFound = errors.New("Item not found")
)

func (a *API) GetDB(title string, reply *[]Item) error {
	*reply = database
	return nil
}

func (a *API) GetByName(title string, reply *Item) error {
	for _, val := range database {
		if title == val.Title {
			*reply = val
			return nil
		}
	}
	return ErrItemNotFound
}

func (a *API) AddItem(item Item, reply *Item) error {
	database = append(database, item)
	*reply = item
	return nil
}

func (a *API) EditItem(edit Item, reply *Item) error {
	for i, _ := range database {
		if edit.Title == database[i].Title {
			database[i] = edit
			*reply = database[i]
			return nil
		}
	}
	return ErrItemNotFound
}

func (a *API) DeleteItem(item Item, reply *Item) error {
	for i, val := range database {
		if item.Title == val.Title && item.Body == val.Body {
			database = append(database[:i], database[i+1:]...)
			*reply = val
			return nil
		}
	}
	return ErrItemNotFound
}

func main() {

	var api = new(API)
	err := rpc.Register(api)
	if err != nil {
		log.Fatal(err)
	}
	rpc.HandleHTTP()

	listener, err := net.Listen("tcp", ":4040")
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("listening rpc on port %d", 4040)
	err = http.Serve(listener, nil)
	if err != nil {
		log.Fatal(err)
	}
	/*
		fmt.Println("Initial database: ", database)
		i1 := Item{"first", "Item 1"}
		i2 := Item{"second", "Item 2"}
		i3 := Item{"third", "Item 3"}

		var api API
		var r Item
		api.AddItem(i1, &r)
		api.AddItem(i2, &r)
		api.AddItem(i3, &r)
		fmt.Println("Second database: ", database)

		api.DeleteItem(i2, &r)
		fmt.Println("Third database:", database)

		api.EditItem(Item{"third", "Item Up"}, &r)
		fmt.Println("Fourth database:", database)

		api.GetByName("fourth", &r)
		api.GetByName("first", &r)
		fmt.Println(r)
	*/
}
