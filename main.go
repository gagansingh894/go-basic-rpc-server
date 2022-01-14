package main

import (
	"flag"
	"fmt"
	"github.com/gagansingh894/go-basic-rpc-server/client"
	"github.com/gagansingh894/go-basic-rpc-server/server"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

func main() {
	isServerFlag := flag.Bool("server", false, "Run code for server")
	flag.Parse()

	if *isServerFlag {
		runServer()
	} else {
		runClient()
	}
}

func runServer() {
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

func runClient() {
	fmt.Println("Running client code")
	client.RunClient()
}
