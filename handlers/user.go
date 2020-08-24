package handlers

import (
	"chat-server/models"
	"encoding/json"
	"github.com/pkg/errors"
	"io/ioutil"
	"log"
	"net/http"
)

func (h *handler) AddUser(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
	}
	defer r.Body.Close()

	var user models.User
	err = json.Unmarshal(b, &user)
	if err != nil {
		err = errors.Wrap(err, "invalid request format")
		Respond(w, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.rep.AddUser(user)

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
