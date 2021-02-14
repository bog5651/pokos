package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func GetInstance() (*sql.DB, error) {
	if db == nil {
		localDb, err := sql.Open("sqlite3", "allan.db")
		if err != nil {
			panic(err)
		}
		db = localDb
		return localDb, err
	}
	return db, nil
}

func CloseDb() {
	if db != nil {
		db.Close()
	}
}
