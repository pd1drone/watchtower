package database

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

func RegisterUser(db sqlx.Ext, username string, password string) error {

	counter := 0
	isadmin := false
	var user string

	rows, err := db.Queryx(`SELECT u.Username FROM Users as u
	WHERE u.Username=?`,
		username)

	if err != nil {
		return err
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&user)
		if err != nil {
			return err
		}
		counter++
	}

	if counter > 0 {
		return fmt.Errorf("Username exists, please use another username")
	}

	_, err = db.Exec(`INSERT INTO Users(
		Username,
		Password,
		IsAdmin
	)
	Values(?,?,?)`,
		username,
		password,
		isadmin,
	)

	if err != nil {
		return err
	}

	return nil
}
