package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type Database struct {
	Db *sql.DB
}

type DbExecutor interface {
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
	Exec(query string, args ...interface{}) (sql.Result, error)
	Close() error
}

func (d *Database) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return d.Db.Query(query, args)
}

func (d *Database) QueryRow(query string, args ...interface{}) *sql.Row {
	return d.Db.QueryRow(query, args)
}

func (d *Database) Exec(query string, args ...interface{}) (sql.Result, error) {
	return d.Db.Exec(query, args)
}

func (d *Database) Close() error {
	return d.Db.Close()
}

var Db DbExecutor

func GetInstance() (DbExecutor, error) {
	if Db == nil {
		localDb, err := sql.Open("sqlite3", "allan.db")
		if err != nil {
			panic(err)
		}
		Db = &Database{Db: localDb}
		return Db, err
	}
	return Db, nil
}

func CloseDb() {
	if Db != nil {
		Db.Close()
	}
}
