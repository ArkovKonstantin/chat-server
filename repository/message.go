package repository

import (
	"chat-server/models"
	"github.com/pkg/errors"
)

func (rep *repository) AddMessage(message models.AddMsgForm) (id int, err error) {
	db, err := rep.p.GetConn()
	if err != nil {
		return id, errors.Wrap(err, "get db connection err:")
	}

	err = db.QueryRow("INSERT INTO \"message\" (chat, author, text) VALUES ($1, $2, $3) RETURNING id", message.Chat, message.Author, message.Text).Scan(&id)
	if err != nil {
		return id, errors.Wrap(err, "get db operation err:")
	}

	return id, err
}

func (rep *repository) GetMessagesByChatID(id int) (messages []*models.Message, err error) {
	db, err := rep.p.GetConn()
	if err != nil {
		return messages, errors.Wrap(err, "get db connection err:")
	}
	rows, err := db.Query("SELECT * FROM messages($1)", id)
	if err != nil {
		return messages, errors.Wrap(err, "get db operation err:")
	}
	defer rows.Close()

	for rows.Next() {
		m := new(models.Message)
		err := rows.Scan(&m.ID, &m.Chat, &m.Author, &m.Text, &m.CreatedAt)
		if err != nil {
			return messages, errors.Wrap(err, "get db operation err:")
		}
		messages = append(messages, m)
	}
	if err = rows.Err(); err != nil {
		return messages, errors.Wrap(err, "get db operation err:")
	}
	return
}
