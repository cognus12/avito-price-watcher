package sqlite

import (
	"apricescrapper/pkg/logger"
	"database/sql"
	"sync"

	_ "github.com/mattn/go-sqlite3"
)

var instance *sql.DB

var lock = &sync.Mutex{}

const CREATE_USERS_TABLE = `
	CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY AUTOINCREMENT, email TEXT)
`

const CREATE_LINKS_TABLE = `
	CREATE TABLE IF NOT EXISTS links (id INTEGER PRIMARY KEY AUTOINCREMENT, url TEXT, subscribers INTEGER FOREIGN KEY (userid) REFERENCES users(id))
`

func initialize() {

	logger := logger.GetInstance()

	db, err := sql.Open("sqlite3", "internal/sqlite/store.db")

	if err != nil {
		logger.Panic(err.Error())
	}

	instance = db

	createTables()
}

func createTables() {
	logger := logger.GetInstance()
	statement, err := instance.Prepare(CREATE_USERS_TABLE)

	if err != nil {
		logger.Fatal(err.Error())
	}

	statement.Exec()
}

func Close() {
	if instance != nil {
		instance.Close()
	}

}

func GetInstance() *sql.DB {
	if instance == nil {
		lock.Lock()
		defer lock.Unlock()

		if instance == nil {
			initialize()
			return instance
		}
	}

	return instance
}
