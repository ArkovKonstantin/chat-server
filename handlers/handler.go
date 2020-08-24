package handlers

import (
	"chat-server/repository"
	"encoding/json"
	"log"
	"net/http"
)

type Handler interface {
	AddUser(w http.ResponseWriter, r *http.Request)

	AddChat(w http.ResponseWriter, r *http.Request)
	GetChatsByUser(w http.ResponseWriter, r *http.Request)

	AddMessage(w http.ResponseWriter, r *http.Request)
	GetMessages(w http.ResponseWriter, r *http.Request)
}

type handler struct {
	rep repository.Repository
}

func New(rep repository.Repository) Handler {
	return &handler{rep}
}

func Respond(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	b, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}
	w.Write(b)
}