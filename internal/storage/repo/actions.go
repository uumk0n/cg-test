package repo

import (
	"cg-test/internal/entities"
	"database/sql"
)

type ActionsRepo struct {
	db *sql.DB
}

func NewActionsRepo(db *sql.DB) *ActionsRepo {
	return &ActionsRepo{db: db}
}

func (r *ActionsRepo) SaveAction(action entities.Actions) error {
	query := "INSERT INTO actions (user_id, action_time, request_result, temp_c) VALUES (?, ?, ?, ?)"
	_, err := r.db.Exec(query, action.UserID, action.ActionTime, action.RequestResult, action.TempC)
	return err
}
