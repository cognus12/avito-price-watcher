package sqlite

import (
	"database/sql"
	"sync"

	_ "github.com/mattn/go-sqlite3"
)

var instance *sql.DB

var lock = &sync.Mutex{}

func initialize(path string) error {
	db, err := sql.Open("sqlite3", path)

	if err != nil {
		return err
	}

	instance = db

	return nil
}

func createTables(schema string) error {
	_, err := instance.Exec(schema)

	return err

}

func New(schema string, path string) (*sql.DB, error) {
	if instance == nil {
		lock.Lock()

		defer lock.Unlock()

		if instance == nil {
			err := initialize(path)

			if err != nil {
				return instance, err
			}

			err = createTables(schema)

			return instance, err
		}
	}

	return instance, nil
}

func Close() error {
	var err error

	if instance != nil {
		err = instance.Close()
	}

	return err
}
