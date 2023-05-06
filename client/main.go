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

	client, err := rpc.DialHTTP("tcp", "localhost:4040")  // creates a client by calling rpc.DialHTTP with the network protocol and address of the server

	if err != nil {
		log.Fatal("Connection error: ", err)
	}

	a := Item{"First", "A first item"}
	b := Item{"Second", "A second item"}
	c := Item{"Third", "A third item"}

	client.Call("API.AddItem", a, &reply)  //calls the server's AddItem method three times with different items as arguments and stores the response in reply
	client.Call("API.AddItem", b, &reply)
	client.Call("API.AddItem", c, &reply)
	client.Call("API.GetDB", "", &db)//then calls the server's GetDB method with an empty string as argument to retrieve the current state of the database, and stores the result in the db variable

	fmt.Println("Database: ", db)

	client.Call("API.EditItem", Item{"Second", "A new second item"}, &reply) // updating  item with new body

	client.Call("API.DeleteItem", c, &reply)
	client.Call("API.GetDB", "", &db)
	fmt.Println("Database: ", db)

	client.Call("API.GetByName", "First", &reply)
	fmt.Println("first item: ", reply)

}
