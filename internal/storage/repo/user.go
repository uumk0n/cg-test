package repo

import (
	"cg-test/internal/entities"
	"database/sql"
	"errors"
)

type UserRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (r *UserRepo) FindUserByLogin(login string) (*entities.User, error) {
	query := "SELECT id, login, hashed_password, fio, api_token FROM users WHERE login = ?"
	row := r.db.QueryRow(query, login)

	var user entities.User
	err := row.Scan(&user.ID, &user.Login, &user.HashedPassword, &user.FIO, &user.APIToken)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // User not found
		}
		return nil, err
	}

	return &user, nil
}

func (r *UserRepo) FindUserByApiToken(apiToken string) (*entities.User, error) {
	query := "SELECT id, login, hashed_password, fio, api_token FROM users WHERE api_token = ?"
	row := r.db.QueryRow(query, apiToken)

	var user entities.User
	err := row.Scan(&user.ID, &user.Login, &user.HashedPassword, &user.FIO, &user.APIToken)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *UserRepo) SaveUser(user entities.User) error {
	query := "INSERT INTO users (login, hashed_password, fio, api_token) VALUES (?, ?, ?, ?)"
	_, err := r.db.Exec(query, user.Login, user.HashedPassword, user.FIO, user.APIToken)
	return err
}
