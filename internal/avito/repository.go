package avito

import (
	"apricescrapper/pkg/logger"
	"database/sql"
)

type Repository interface {
	CreateUser(email string) error
	CreateLink(url string) error
	GetUser(email string) (UserDTO, error)
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

func (r *repository) CreateUser(email string) error {

	statement, err := r.db.Prepare(CREATE_USER)

	defer statement.Close()

	r.logger.Info("Call CreateUser method")

	if err != nil {
		r.logger.Error(err.Error())

		return err
	}

	statement.Exec(nil, email)

	return nil
}

func (r *repository) GetUser(email string) (UserDTO, error) {

	row := r.db.QueryRow("SELECT * FROM users WHERE email = ?", email)

	// defer row.Close()

	var u UserDTO

	err := row.Scan(&u.Id, &u.Email)

	if err != nil {
		return u, err
	}

	r.logger.Info("Get user %+v", u)

	return u, nil
}

func (r *repository) CreateLink(url string) error {
	//  TODO implement CreateLink
	return nil
}
