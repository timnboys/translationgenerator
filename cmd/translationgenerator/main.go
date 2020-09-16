package main

import (
	"encoding/json"
	database "github.com/TicketsBot/database/translations"
	"github.com/TicketsBot/translationgenerator/server"
	"github.com/TicketsBot/translationgenerator/server/endpoints"
	"io/ioutil"
)

func main() {
	loadEnglish()

	server := server.NewServer()
	server.ConfigureEngine()
	server.Start()
}

func loadEnglish() {
	bytes, err := ioutil.ReadFile("en.json")
	if err != nil {
		panic(err)
	}

	var array []endpoints.Translation
	if err := json.Unmarshal(bytes, &array); err != nil {
		panic(err)
	}

	endpoints.English = make(map[database.MessageId]string)
	for _, message := range array {
		endpoints.English[message.Id] = message.Value
	}
}
