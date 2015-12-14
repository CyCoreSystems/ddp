package main

import (
	"log"

	"github.com/cycoresystems/ddpd/server"
)

func main() {
	s := server.NewServer(":1234")
	if err := s.Run(); err != nil {
		log.Fatal("Error running DDP server: " + err.Error())
	}
}
