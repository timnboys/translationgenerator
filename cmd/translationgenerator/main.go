package main

import (
	"github.com/TicketsBot/translationgenerator/server"
)

func main() {
	server := server.NewServer()
	server.ConfigureEngine()
	server.Start()
}
