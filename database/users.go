package database

import (
	"github.com/jmoiron/sqlx"
)

type Users struct {
	ID       int64  `json:"ID"`
	Username string `json:"Username"`
	Password string `json:""`
	IsAdmin  bool   `json:"IsAdmin"`
}

type UsersArray struct {
	UsersArray []*Users `json:"UsersArray"`
}

func CreateUser(db sqlx.Ext, username string, password string, isadmin bool) error {

	_, err := db.Exec(`INSERT INTO Users (
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

func DeleteUser(db sqlx.Ext, ID int64) error {

	_, err := db.Exec(`DELETE FROM Users WHERE ID = ? `, ID)

	if err != nil {
		return err
	}

	return nil
}

func UpdateUserWithPassword(db sqlx.Ext, ID int64, username string, password string, isadmin bool) error {

	_, err := db.Exec(`UPDATE Users SET 
		Username =?,
		Password =?,
		IsAdmin = ? WHERE ID= ?`,
		username,
		password,
		isadmin,
		ID,
	)

	if err != nil {
		return err
	}

	return nil
}

func UpdateUserWithOutPassword(db sqlx.Ext, ID int64, username string, isadmin bool) error {

	_, err := db.Exec(`UPDATE Users SET 
		Username =?,
		IsAdmin = ? WHERE ID= ?`,
		username,
		isadmin,
		ID,
	)

	if err != nil {
		return err
	}

	return nil
}

func ReadUsers(db sqlx.Ext) (*UsersArray, error) {

	usersArray := make([]*Users, 0)
	var ID int64
	var Username string
	var IsAdmin bool

	rows, err := db.Queryx(`SELECT ID,
				Username,
				IsAdmin FROM Users`)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&ID, &Username, &IsAdmin)
		if err != nil {
			return nil, err
		}

		usersArray = append(usersArray, &Users{
			ID:       ID,
			Username: Username,
			IsAdmin:  IsAdmin,
		})

	}
	return &UsersArray{
		UsersArray: usersArray,
	}, nil
}
