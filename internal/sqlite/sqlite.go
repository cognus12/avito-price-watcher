package sqlite

import (
	"apricescrapper/pkg/logger"
	"database/sql"
	"sync"

	_ "github.com/mattn/go-sqlite3"
)

var instance *sql.DB

var lock = &sync.Mutex{}

func initialize() {
	logger := logger.GetInstance()

	db, err := sql.Open("sqlite3", "internal/sqlite/store.db")

	if err != nil {
		logger.Panic(err.Error())
	}

	instance = db
}

func createTables(schema string) {
	logger := logger.GetInstance()
	_, err := instance.Exec(schema)

	if err != nil {
		logger.Fatal(err.Error())
	}

}

func New(schema string) *sql.DB {
	if instance == nil {
		lock.Lock()

		defer lock.Unlock()

		if instance == nil {
			initialize()
			createTables(schema)

			return instance
		}
	}

	return instance
}

func Close() {
	if instance != nil {
		instance.Close()
	}

}
