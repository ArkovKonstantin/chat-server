package handlers

import (
	"chat-server/models"
	"chat-server/repository/mocks"
	"github.com/golang/mock/gomock"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestAddUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	rep := mocks.NewMockRepository(ctrl)

	u := models.User{Username: "user_1"}
	rep.EXPECT().AddUser(u).Return(1, nil)

	handlerInterface := New(rep)

	handler := http.HandlerFunc(handlerInterface.AddUser)

	r, err := http.NewRequest("POST", "users/add", strings.NewReader("{\"username\": \"user_1\"}"))
	if err != nil {
		t.Errorf("error: %v", err)
	}

	w := httptest.NewRecorder()

	handler.ServeHTTP(w, r)

	if w.Code != http.StatusCreated {
		t.Errorf("Wrong code. Expected %d, got %d", http.StatusCreated, w.Code)
	}

	expected := `{"id":1}`

	if w.Body.String() != expected {
		t.Errorf(`expected %s, got %s`, expected, w.Body.String())
	}

}

func TestAddChat(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	rep := mocks.NewMockRepository(ctrl)

	chat := models.Chat{Name: "chat_1", Users: []int{1, 2}}
	rep.EXPECT().AddChat(chat).Return(1, nil)

	handlerInterface := New(rep)

	handler := http.HandlerFunc(handlerInterface.AddChat)

	r, err := http.NewRequest("POST", "chats/add", strings.NewReader("{\"name\": \"chat_1\", \"users\": [1, 2]}"))
	if err != nil {
		t.Errorf("error: %v", err)
	}

	w := httptest.NewRecorder()

	handler.ServeHTTP(w, r)

	if w.Code != http.StatusCreated {
		t.Errorf("Wrong code. Expected %d, got %d", http.StatusCreated, w.Code)
	}

	expected := `{"id":1}`

	if w.Body.String() != expected {
		t.Errorf(`expected %s, got %s`, expected, w.Body.String())
	}
}

func TestGetChatsByUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	rep := mocks.NewMockRepository(ctrl)
	chats := make([]*models.Chat, 0)
	chats = append(chats, &models.Chat{ID: 1, Name: "chat_1", CreatedAt: "2020-08-24 12:15:20.526098"})
	rep.EXPECT().GetChatsByUserID(1).Return(chats, nil)

	handlerInterface := New(rep)

	handler := http.HandlerFunc(handlerInterface.GetChatsByUser)

	r, err := http.NewRequest("POST", "chats/get", strings.NewReader("{\"user\": 1}"))
	if err != nil {
		t.Errorf("error: %v", err)
	}

	w := httptest.NewRecorder()

	handler.ServeHTTP(w, r)

	if w.Code != http.StatusOK {
		t.Errorf("Wrong code. Expected %d, got %d", http.StatusOK, w.Code)
	}

	expected := `{"chats":[{"id":1,"name":"chat_1","created_at":"2020-08-24 12:15:20.526098"}]}`

	if w.Body.String() != expected {
		t.Errorf(`expected %s, got %s`, expected, w.Body.String())
	}
}

func TestGetMessages(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	rep := mocks.NewMockRepository(ctrl)

	messages := make([]*models.Message, 0)
	messages = append(messages, &models.Message{ID: 1, Chat: "chat_1", Author: "user_1", Text: "hi"})
	messages = append(messages, &models.Message{ID: 2, Chat: "chat_1", Author: "user_2", Text: "hi"})

	rep.EXPECT().GetMessagesByChatID(1).Return(messages, nil)

	handlerInterface := New(rep)

	handler := http.HandlerFunc(handlerInterface.GetMessages)

	r, err := http.NewRequest("POST", "messages/get", strings.NewReader("{\"chat\": 1}"))
	if err != nil {
		t.Errorf("error: %v", err)
	}

	w := httptest.NewRecorder()

	handler.ServeHTTP(w, r)

	if w.Code != http.StatusOK {
		t.Errorf("Wrong code. Expected %d, got %d", http.StatusOK, w.Code)
	}

	expected := `{"messages":[{"id":1,"chat":"chat_1","author":"user_1","text":"hi","created_at":""},{"id":2,"chat":"chat_1","author":"user_2","text":"hi","created_at":""}]}`

	if w.Body.String() != expected {
		t.Errorf(`expected %s, got %s`, expected, w.Body.String())
	}
}

func TestAddMessage(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	rep := mocks.NewMockRepository(ctrl)

	m := models.AddMsgForm{Chat: 1, Author: 1, Text: "hi"}
	rep.EXPECT().AddMessage(m).Return(1, nil)

	handlerInterface := New(rep)

	handler := http.HandlerFunc(handlerInterface.AddMessage)

	r, err := http.NewRequest("POST", "messages/add", strings.NewReader("{\"chat\": 1, \"author\": 1, \"text\": \"hi\"}"))
	if err != nil {
		t.Errorf("error: %v", err)
	}

	w := httptest.NewRecorder()

	handler.ServeHTTP(w, r)

	if w.Code != http.StatusCreated {
		t.Errorf("Wrong code. Expected %d, got %d", http.StatusCreated, w.Code)
	}

	expected := `{"id":1}`

	if w.Body.String() != expected {
		t.Errorf(`expected %s, got %s`, expected, w.Body.String())
	}
}
