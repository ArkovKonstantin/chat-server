package repository

import (
	"chat-server/models"
	"github.com/pkg/errors"
)

func (rep *repository) AddChat(chat models.Chat) (cid int, err error) {
	db, err := rep.p.GetConn()
	if err != nil {
		return cid, errors.Wrap(err, "get db connection err:")
	}

	err = db.QueryRow("INSERT INTO \"chat\" (name) VALUES ($1) RETURNING id", chat.Name).Scan(&cid)
	if err != nil {
		return cid, errors.Wrap(err, "get db operation err:")
	}

	for _, uid := range chat.Users {
		_, err = db.Exec("INSERT INTO \"user_to_chat\" (user_id, chat_id) VALUES ($1, $2)", uid, cid)
		if err != nil {
			return cid, errors.Wrap(err, "get db operation err:")
		}
	}

	return cid, nil
}

func (rep *repository) GetChatsByUserID(id int) (chats []*models.Chat, err error) {
	db, err := rep.p.GetConn()
	if err != nil {
		return chats, errors.Wrap(err, "get db connection err:")
	}
	rows, err := db.Query("SELECT * FROM chats($1)", id)
	if err != nil {
		return chats, errors.Wrap(err, "get db operation err:")
	}
	defer rows.Close()

	for rows.Next() {
		chat := new(models.Chat)
		err := rows.Scan(&chat.ID, &chat.Name, &chat.CreatedAt)
		if err != nil {
			return chats, errors.Wrap(err, "get db operation err:")
		}
		chats = append(chats, chat)
	}
	if err = rows.Err(); err != nil {
		return chats, errors.Wrap(err, "get db operation err:")
	}

	return
}
