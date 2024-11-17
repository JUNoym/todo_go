package models

import (
	"database/sql"
	"log"
	"todo_app/config"

	_ "github.com/mattn/go-sqlite3"
)

var Db *sql.DB

var err error

func init() {
	Db, err = sql.Open(config.Config.SQLDriver, config.Config.DBName)
	if err != nil {
		log.Fatalln(err)
	}

	cmdU := `CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		uuid STRING NOT NULL UNIQUE,
		name STRING,
		email STRING,
		password STRING,
		created_at DATETIME
	)`

	Db.Exec(cmdU)

}
