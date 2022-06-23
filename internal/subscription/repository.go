package subscription

import (
	"apricescrapper/internal/apperror"
	"apricescrapper/pkg/logger"
	"database/sql"
	"fmt"
)

type Repository interface {
	CreateSubscibtion(url string, email string) error
	DeleteSubscibtion(url string, email string) error
	FindAll() []SubscribtionDTO
}

type repository struct {
	db     *sql.DB
	logger logger.Logger
}

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
	res, err := r.db.Exec(query, url, email, url, email)

	if err != nil {
		return err
	}

	created, err := res.RowsAffected()

	if err != nil {
		return err
	}

	if created == 0 {
		alreadyExists := apperror.ErrAlreadyExists

		alreadyExists.Message = fmt.Sprintf("Email %v is already subscribed for url %v", email, url)

		return alreadyExists
	}

	r.logger.Info("Created subscribtion, url: %v, email: %v", url, email)

	return err
}

func (r *repository) DeleteSubscibtion(url string, email string) error {
	query := `
		DELETE FROM subscriptions 
		WHERE EXISTS (SELECT id FROM links WHERE url = ?) AND user_id = (SELECT id FROM users WHERE email = ?);
	`

	res, err := r.db.Exec(query, url, email, email, email, url, url)

	if err != nil {
		return err
	}

	deleted, err := res.RowsAffected()

	if err != nil {
		return err
	}

	if deleted == 0 {
		notFound := apperror.ErrNotFound

		notFound.Message = fmt.Sprintf("email %v is not subscribed for url %v", email, url)

		return apperror.ErrNotFound
	}

	r.logger.Info("Deleted subscribtion, url: %v, email: %v", url, email)

	return err
}

func (r *repository) FindAll() []SubscribtionDTO {
	query := `
		SELECT l.url, u.email
		FROM subscriptions as s
		INNER JOIN links as l
			ON s.link_id = l.id
		INNER JOIN users as u
			ON s.user_id = u.id
	`

	rows, err := r.db.Query(query)

	if err != nil {
		r.logger.Panic(err.Error())
	}

	defer rows.Close()

	var entries []SubscribtionDTO

	for rows.Next() {
		item := SubscribtionDTO{}
		err := rows.Scan(&item.Url, &item.Email)
		if err != nil {
			r.logger.Errorf(err)
			continue
		}
		entries = append(entries, item)
	}

	return entries
}

/*

SELECT m.name, cp.id_category
FROM manufacturer as m
INNER JOIN product as p
    ON m.id_manufacturer = p.id_manufacturer
INNER JOIN category_product as cp
    ON p.id_product = cp.id_product
WHERE cp.id_category = 'some value'

*/
