package main

import (
	"github.com/TicketsBot/translationgenerator/config"
	"github.com/TicketsBot/translationgenerator/server"
)

func main() {
	conf := config.LoadConfig()

	server := server.NewServer(conf)
	server.ConfigureEngine()
	server.Start()
}
