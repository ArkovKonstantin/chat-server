package repository

import (
	"chat-server/models"
	"chat-server/provider"
)

type Repository interface {
	AddUser(user models.User) (int, error)

	AddChat(chat models.Chat) (int, error)
	GetChatsByUserID(id int) (chats []*models.Chat, err error)

	AddMessage(message models.AddMsgForm) (int, error)
	GetMessagesByChatID(id int) ([]*models.Message, error)
}

type repository struct {
	p provider.Provider
}

func New(p provider.Provider) Repository {
	return &repository{p}
}

