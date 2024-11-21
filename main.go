package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Server is starting...!!")
	go StartServer()
	// add delay of 1 second to ensure server is up before client starts
	time.Sleep(time.Second * 15)
	fmt.Println("Client is starting...")
	StartClient()
}
