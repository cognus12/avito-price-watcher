package avito

import (
	"apricescrapper/pkg/logger"
	"database/sql"
)

type Repository interface {
	CreateSubscibtion(url string, email string) error
	DeleteSubscibtion(url string, email string) error
}

type repository struct {
	db     *sql.DB
	logger logger.Logger
}

const CREATE_USER = `
	INSERT INTO users (id, email) 
	VALUES (?, ?)
`

func NewRepository(db *sql.DB) Repository {
	logger := logger.GetInstance()

	return &repository{db: db, logger: logger}
}

func (r *repository) CreateSubscibtion(url string, email string) error {
	query := `
		INSERT OR IGNORE INTO links (url) VALUES (?);
		INSERT OR IGNORE INTO users (email) VALUES (?);
		INSERT OR IGNORE INTO subscriptions VALUES(
			(SELECT id FROM links WHERE url = ?), 
			(SELECT id FROM users WHERE email = ?)
		);
	`
	_, err := r.db.Exec(query, url, email, url, email)

	return err
}

func (r *repository) DeleteSubscibtion(url string, email string) error {
	// TODO implement

	// query := `
	// 	DELETE FROM subscriptions VALUES(
	// 		(SELECT id FROM links WHERE url = ?),
	// 		(SELECT id FROM users WHERE email = ?)
	// 	);
	// `
	query := `
		DELETE FROM subscriptions 
		WHERE EXISTS (SELECT id FROM links WHERE url = ?) AND user_id = (SELECT id FROM users WHERE email = ?);
	`

	res, err := r.db.Exec(query, url, email, email, email, url, url)

	rows, _ := res.RowsAffected()

	r.logger.Info("Rows affected: %v", rows)

	return err
}
