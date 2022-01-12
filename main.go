package main

import (
	"fmt"
	"github.com/gagansingh894/go-basic-rpc-server/client"
	"github.com/gagansingh894/go-basic-rpc-server/server"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"time"
)

func main() {
	go func() {
		time.Sleep(time.Second * 2)
		fmt.Println("Running client test")
		client.TestClient()
	}()

	api := new(server.API)
	err := rpc.Register(api)
	if err != nil {
		log.Fatal("error registering API", err)
	}
	rpc.HandleHTTP()
	listener, err := net.Listen("tcp", ":4040")
	if err != nil {
		log.Fatal("Listener error", err)
	}
	log.Printf("serving rpc on port %d", 4040)
	err = http.Serve(listener, nil)

	if err != nil {
		log.Fatal("error serving: ", err)
	}

}
