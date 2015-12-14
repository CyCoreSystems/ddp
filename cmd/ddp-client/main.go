package main

import (
	"log"

	"github.com/cycoresystems/ddp/client"
)

func main() {
	c := client.NewClient("ws://localhost:1234/websocket")
	if err := c.Run(); err != nil {
		log.Fatal("Error running DDP client: " + err.Error())
	}
}
