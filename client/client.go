package client

import (
	"fmt"
	"log"
	"net/rpc"
)

type Item struct {
	Title string
	Body  string
}

func TestClient() {
	var reply Item
	var db []Item

	client, err := rpc.Dial("tcp", "localhost:4040")

	if err != nil {
		log.Fatal("Connection error: ", err)
	}

	a := Item{"One", "Item One"}
	b := Item{"Two", "Item Two"}
	c := Item{"Three", "Item Three"}

	_ = client.Call("API.AddItem", a, &reply)
	_ = client.Call("API.AddItem", b, &reply)
	_ = client.Call("API.AddItem", c, &reply)
	err = client.Call("API.GetDB", "", &db)

	fmt.Println("Database: ", db)
}
