package main

import (
	"chat-server/application"
	"chat-server/handlers"
	"chat-server/models"
	"chat-server/provider"
	"chat-server/repository"
	"log"
)

var (
	config models.Config
)

func init() {
	models.LoadConfig(&config)
}
func main() {
	p := provider.New(&config.SQLDataBase)
	err := p.Open()

	if err != nil {
		log.Fatal(err)
	}

	rep := repository.New(p)
	handler := handlers.New(rep)

	app := application.New(handler)
	app.Start()
}
