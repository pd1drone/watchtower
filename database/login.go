package database

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type LoginResponse struct {
	ID       int64  `json:"ID"`
	Username string `json:"Username"`
	Password string `json:"Password"`
	IsAdmin  bool   `json:"IsAdmin"`
}

func Login(db sqlx.Ext, username string, password string) (*LoginResponse, error) {

	counter := 0
	var id int64
	var user string
	var pass string
	var isadmin bool

	rows, err := db.Queryx(`SELECT u.ID, u.Username, u.Password, u.IsAdmin FROM Users as u
	WHERE u.Username=? AND u.Password=?`,
		username, password)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&id, &user, &pass, &isadmin)
		if err != nil {
			return nil, err
		}
		counter++
	}

	if counter == 0 {
		return nil, fmt.Errorf("User does not exists")
	}

	return &LoginResponse{
		ID:       id,
		Username: user,
		Password: pass,
		IsAdmin:  isadmin,
	}, nil
}
