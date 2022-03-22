package sqlite

import (
	"apricescrapper/pkg/logger"
	"database/sql"
	"sync"

	_ "github.com/mattn/go-sqlite3"
)

var instance *sql.DB

var lock = &sync.Mutex{}

const CREATE_TABLES = `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER UNIQUE NOT NULL PRIMARY KEY AUTOINCREMENT,
		email TEXT 
	);

	CREATE TABLE IF NOT EXISTS links (
		id INTEGER UNIQUE NOT NULL PRIMARY KEY AUTOINCREMENT, 
		url TEXT 

	);

	CREATE TABLE IF NOT EXISTS links_users (
		link_id INTEGER,
		user_id INTEGER,
		FOREIGN KEY (link_id) REFERENCES links(id),
		FOREIGN KEY (user_id) REFERENCES users(id)
	);
`

func initialize() {

	logger := logger.GetInstance()

	db, err := sql.Open("sqlite3", "internal/sqlite/store.db")

	logger.Info("Connected to database internal/sqlite/store.db")

	if err != nil {
		logger.Panic(err.Error())
	}

	instance = db

	createTables()
}

func createTables() {
	logger := logger.GetInstance()
	_, err := instance.Exec(CREATE_TABLES)

	if err != nil {
		logger.Fatal(err.Error())
	}

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
