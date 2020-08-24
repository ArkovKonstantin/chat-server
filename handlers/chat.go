package handlers

import (
	"chat-server/models"
	"encoding/json"
	"github.com/pkg/errors"
	"io/ioutil"
	"log"
	"net/http"
)

func (h *handler) AddChat(w http.ResponseWriter, r *http.Request) {
	var chat models.Chat
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
	}
	defer r.Body.Close()

	err = json.Unmarshal(b, &chat)
	if err != nil {
		err = errors.Wrap(err, "invalid request format")
		Respond(w, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.rep.AddChat(chat)

	if err != nil {
		Respond(w, http.StatusBadRequest, struct {
			Error string `json:"error"`
		}{err.Error()})
		return
	}

	Respond(w, http.StatusCreated, struct {
		Id int `json:"id"`
	}{id})
}

func (h *handler) GetChatsByUser(w http.ResponseWriter, r *http.Request) {
	var data struct {
		User int `json:"user"`
	}
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
	}
	defer r.Body.Close()

	err = json.Unmarshal(b, &data)
	if err != nil {
		err = errors.Wrap(err, "invalid request format")
		Respond(w, http.StatusBadRequest, err.Error())
		return
	}
	chats, err := h.rep.GetChatsByUserID(data.User)

	if err != nil {
		Respond(w, http.StatusBadRequest, struct {
			Error string `json:"error"`
		}{err.Error()})
		return
	}

	Respond(w, http.StatusOK, struct {
		Chats []*models.Chat `json:"chats"`
	}{chats})
}
