package rest

import (
	"fmt"
	"watchtower/database"

	"github.com/jmoiron/sqlx"
	"gopkg.in/ini.v1"
)

type WatchTowerConfiguration struct {
	WatchtowerDB *sqlx.DB
}

func New() (*WatchTowerConfiguration, error) {

	// read config file
	cfg, err := ini.Load("/root/watchtower/config.ini")
	if err != nil {
		return nil, fmt.Errorf("Fail to read file: %v", err)
	}

	dbSection := cfg.Section("db")
	user := dbSection.Key("user").String()
	password := dbSection.Key("password").String()
	dbhost := dbSection.Key("dbhost").String()
	dbport := dbSection.Key("dbport").String()
	dbname := dbSection.Key("dbname").String()

	wtdb, err := database.InitializeWatchTowerDatabase(dbname, user, password, dbhost, dbport)
	if err != nil {
		return nil, err
	}

	return &WatchTowerConfiguration{
		wtdb,
	}, nil
}
