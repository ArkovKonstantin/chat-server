package repository

import (
	"chat-server/models"
	"github.com/pkg/errors"
)

func (rep *repository) AddUser(user models.User) (id int, err error) {
	db, err := rep.p.GetConn()
	if err != nil {
		return id, errors.Wrap(err, "get db connection err:")
	}
	err = db.QueryRow("INSERT INTO \"user\" (username) VALUES($1) RETURNING id", user.Username).Scan(&id)
	if err != nil {
		return id, errors.Wrap(err, "user already exists:")
	}

	return id, nil
}
