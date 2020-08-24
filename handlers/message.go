package handlers

import (
	"chat-server/models"
	"encoding/json"
	"github.com/pkg/errors"
	"io/ioutil"
	"log"
	"net/http"
)

func (h *handler) GetMessages(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Chat int `json:"chat"`
	}
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		Respond(w, http.StatusInternalServerError, err.Error())
	}
	defer r.Body.Close()

	err = json.Unmarshal(b, &data)
	if err != nil {
		err = errors.Wrap(err, "invalid request format")
		Respond(w, http.StatusBadRequest, err.Error())
		return
	}

	messages, err := h.rep.GetMessagesByChatID(data.Chat)

	if err != nil {
		Respond(w, http.StatusBadRequest, struct {
			Error string `json:"error"`
		}{err.Error()})
		return
	}
	Respond(w, http.StatusOK, struct {
		Messages []*models.Message `json:"messages"`
	}{messages})

}

func (h *handler) AddMessage(w http.ResponseWriter, r *http.Request) {
	var message models.AddMsgForm
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
	}
	defer r.Body.Close()

	err = json.Unmarshal(b, &message)
	if err != nil {
		err = errors.Wrap(err, "invalid request format")
		Respond(w, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.rep.AddMessage(message)

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
