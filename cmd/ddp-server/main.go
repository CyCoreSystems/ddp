package main

import (
	"log"

	"github.com/CyCoreSystems/ddp/server"
)

func main() {
	s := server.NewServer("127.0.0.1:1234")
	if err := s.Run(); err != nil {
		log.Fatal("Error running DDP server: " + err.Error())
	}
}
